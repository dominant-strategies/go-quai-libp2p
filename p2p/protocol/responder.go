package protocol

import (
	"crypto/rand"
	"crypto/sha256"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/pkg/errors"

	"github.com/dominant-strategies/go-quai/log"
)

// processJoinResponse is a callback function used by the quai protocol handler
// whenever a peer sends a request to join the quai network.
// This function implements the responder side of the protocol handshake.
func processJoinRequest(stream network.Stream) error {
	// Send a challenge to the peer to verify its identity
	nonce, err := createChallenge()
	if err != nil {
		return errors.Wrap(err, "error creating challenge")
	}

	log.Debugf("Sending challenge to peer %s", stream.Conn().RemotePeer())
	challengeMessage := QuaiProtocolMessage{
		Flag: challengeFlag,
		Data: nonce,
	}
	if err := writeQuaiMessage(stream, &challengeMessage); err != nil {
		return errors.Wrap(err, "error sending challenge")
	}

	// Wait for the peer's response
	challengeResponse, err := readQuaiMessageWithTimeout(stream)
	if err != nil {
		return errors.Wrap(err, "error reading challenge response")
	}
	// Verify the peer's response to the challenge
	verified, err := verifyChallengeMsg(stream, nonce, challengeResponse)
	if err != nil {
		return errors.Wrap(err, "error verifying challenge")
	}
	if !verified {
		errorMsg := QuaiProtocolMessage{
			Flag: errorFlag,
			// TODO: add error message and/or error code
		}
		if err := writeQuaiMessage(stream, &errorMsg); err != nil {
			log.Warnf("error sending error message: %s", err)
		}
		stream.Close()
		return errors.New("peer not verified")
	}
	log.Debugf("Peer %s verified!", stream.Conn().RemotePeer())
	// send welcome message
	welcomeMessage := QuaiProtocolMessage{
		Flag: welcomeFlag,
	}
	if err := writeQuaiMessage(stream, &welcomeMessage); err != nil {
		return errors.Wrap(err, "error sending welcome message")
	}
	return nil
}

// Creates a random 32 bytes nonce to be used as a challenge to be
// signed by the peer.
func createChallenge() ([]byte, error) {
	nonce := make([]byte, 32) // 32 bytes random data
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, errors.Wrap(err, "error generating nonce")
	}
	return nonce, nil
}

// Verifies the signature of the challenge message received from the peer.
// The public key used to verify the message is expected to be extractable from the stream's remote peer ID.
func verifyChallengeMsg(stream network.Stream, nonce []byte, challengeResponse *QuaiProtocolMessage) (bool, error) {
	// check if the message is a challenge response
	// TODO: evaluate implementing a more robust way of checking the message type
	if challengeResponse.Flag != challengeResponseFlag {
		errMsg := QuaiProtocolMessage{
			Flag: errorFlag,
		}
		if err := writeQuaiMessage(stream, &errMsg); err != nil {
			log.Warnf("error sending error message: %s", err)
		}
		err := errors.New("expected challenge response message")
		log.Debugf("%s", err)
		return false, err
	}
	// get the signature from the message
	signature := challengeResponse.Data

	// get the public key of the remote peer
	peerID := stream.Conn().RemotePeer()
	pubKey, err := peerID.ExtractPublicKey()
	if err != nil {
		// TODO: consider other ways of getting the public key
		err = errors.Wrap(err, "error extracting public key from peer ID")
		log.Debugf("%s", err)
		return false, err
	}
	log.Debugf("Verifying signature of peer %s (public key type: %s)", peerID, pubKey.Type().String())
	return verifySignature(nonce, signature, pubKey)
}

// Verifies the signature of the given nonce using the given public key.
// This function is used by the responder to verify the signature sent by the peer
func verifySignature(nonce []byte, signature []byte, pubKey crypto.PubKey) (bool, error) {
	hash := sha256.Sum256(nonce)
	verified, err := pubKey.Verify(hash[:], signature)
	if err != nil {
		err = errors.Wrap(err, "error verifying signature")
		log.Debugf("%s", err)
		return false, err
	}
	if !verified {
		log.Debugf("Signature verification failed!")
		return false, nil
	}
	log.Debugf("Signature verified!")
	return true, nil
}

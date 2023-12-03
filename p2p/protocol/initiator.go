package protocol

import (
	"github.com/dominant-strategies/go-quai/log"
	"github.com/libp2p/go-libp2p/core/network"

	"github.com/pkg/errors"
)

// JoinNetwork is a function called by a quai p2p node (an initiator) to join the quai network.
// This function is called during the node's startup process.
// If this function returns an error the node will fail to join the network.
// It attempts to:
//  1. Connect to the bootnodes
//  2. Open a stream to each connected peer using the Quai protocol
//  3. Send a join request through the stream
//  4. Perform the protocol handshake
func JoinNetwork(p QuaiP2PNode) error {
	// Attempt to connect to the bootnodes
	var connected bool
	for _, peer := range p.GetBootPeers() {
		log.Debugf("attempting to connect to peer: %s", peer.ID)
		if err := p.Connect(peer); err != nil {
			log.Warnf("error connecting to peer: %s", err)
			continue
		}
		connected = true
		log.Debugf("connected to peer: %s", peer.ID)
	}
	if !connected {
		return errors.New("could not connect to any bootnode")
	}
	// bool to indicate if the node has been validated by at least one bootnode
	var validated bool

	// Open a stream to each connected peer using the Quai protocol
	for _, peerID := range p.Network().Peers() {
		stream, err := p.NewStream(peerID, ProtocolVersion)
		if err != nil {
			log.Warnf("error opening stream to peer %s: %s", peerID, err)
			continue
		}
		defer stream.Close()
		log.Debugf("opened stream to peer %s", peerID)
		// Send a join request through the stream
		if err := sendJoinRequest(stream); err != nil {
			log.Warnf("error sending join request to peer %s: %s", peerID, err)
			continue
		}
		log.Debugf("sent join request to peer %s", peerID)

		// Wait for and process the response
		// TODO: consider returning an error and a bool to signal the reason for node
		// not being able to join the network
		if err := processJoinResponse(stream, p); err != nil {
			log.Warnf("error processing join response from peer %s: %s", peerID, err)
			continue
		}
		// TODO: having 1 bootnode validate the node is enough?
		if !validated {
			validated = true
			log.Debugf("node validated by peer %s", peerID)
		}
	}
	if !validated {
		return errors.New("node not validated by any bootnode")
	}
	log.Info("node successfully joined the network")
	return nil
}

// Send a join request through the stream to the bootstrap node
func sendJoinRequest(stream network.Stream) error {
	// TODO: implement a more robust join request
	requestMessage := QuaiProtocolMessage{
		Flag: joinFlag,
	}
	return writeQuaiMessage(stream, &requestMessage)
}

// processJoinResponse implements the responder side of the protocol handshake,
// including the challenge-response authentication
func processJoinResponse(stream network.Stream, p QuaiP2PNode) error {
	response, err := readQuaiMessage(stream)
	if err != nil {
		return errors.Wrap(err, "error reading response from stream")
	}

	// get the challenge nonce from the response, sign it and send it back
	if response.Flag != challengeFlag {
		return errors.New("invalid response flag")
	}
	log.Debugf("received challenge from peer %s", stream.Conn().RemotePeer())

	// sign the challenge
	nonce := response.Data
	signature, err := p.SignChallenge(nonce)
	if err != nil {
		return errors.Wrap(err, "error signing challenge")
	}

	// send the signed challenge back to the peer
	responseMessage := QuaiProtocolMessage{
		Flag: challengeResponseFlag,
		Data: signature,
	}

	if err := writeQuaiMessage(stream, &responseMessage); err != nil {
		return errors.Wrap(err, "error sending challenge response")
	}
	log.Debugf("sent challenge response to peer %s", stream.Conn().RemotePeer())

	// wait for the welcome message
	response, err = readQuaiMessage(stream)
	if err != nil {
		return errors.Wrap(err, "error reading response from stream")
	}
	log.Debugf("received challenge response from peer %s", stream.Conn().RemotePeer())
	if response.Flag != welcomeFlag {
		log.Debugf("Join request unsuccesfull from peer %s", stream.Conn().RemotePeer())
		return errors.New("Unsuccesful join process: " + string(response.Data))
	}

	log.Infof("Join request successful: %s", response.Data)
	return nil
}

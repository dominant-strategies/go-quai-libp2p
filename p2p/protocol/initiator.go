package protocol

import (
	"github.com/dominant-strategies/go-quai/log"
	"github.com/libp2p/go-libp2p/core/network"

	"github.com/pkg/errors"
)

// Join the node to the quai p2p network
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

		// Send a join request through the stream
		if err := sendJoinRequest(stream); err != nil {
			log.Warnf("error sending join request to peer %s: %s", peerID, err)
			continue
		}

		// Wait for and process the response
		if err := processJoinResponse(stream); err != nil {
			log.Warnf("error processing join response from peer %s: %s", peerID, err)
			continue
		}
		if !validated {
			validated = true
		}
	}
	if !validated {
		return errors.New("node not validated by any bootnode")
	}
	return nil
}

// Send a join request through the stream using the joinSignal
func sendJoinRequest(stream network.Stream) error {
	requestMessage := QuaiProtocolMessage{
		Flag: joinFlag,
	}
	return writeQuaiMessage(stream, &requestMessage)
}

// Wait for and process the response
func processJoinResponse(stream network.Stream) error {
	response, err := readQuaiMessage(stream)
	if err != nil {
		return errors.Wrap(err, "error reading response from stream")
	}

	if response.Flag != welcomeFlag {
		return errors.New("Unsuccesful join process: " + string(response.Data))
	}

	log.Infof("Join request successful: %s", response.Data)
	return nil
}

package protocol

import (
	"github.com/libp2p/go-libp2p/core/network"
)

// Handles a peer's request to join the Quai p2p network.
func processJoinRequest(stream network.Stream) error {
	// TODO: add signature validation
	// TODO: add node to DHT
	// TODO: consider adding node to peerstore
	joinRequestMessage := QuaiProtocolMessage{
		Flag: welcomeFlag,
	}
	return writeQuaiMessage(stream, &joinRequestMessage)
}

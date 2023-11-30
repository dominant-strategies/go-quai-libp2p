package protocol

import (
	"github.com/libp2p/go-libp2p/core/protocol"
)

const (
	// ProtocolVersion is the current version of the Quai protocol
	ProtocolVersion protocol.ID = "/quai/1.0.0"
	// flag to indicate willingness to join the network from the initiator node
	joinFlag byte = 0x01
	// flag to indicate welcome message from the responder node
	welcomeFlag byte = 0x02
	// flag to indicate a challenge (nonce) from the responder node
	challengeFlag byte = 0x03
	// flag to indicate a response to the challenge from the initiator node
	challengeResponseFlag byte = 0x04
	// TODO: should we define different error codes?
	// flag to indicate error message from the responder node
	errorFlag byte = 0x05
)

package protocol

import (
	"github.com/dominant-strategies/go-quai/log"
	"github.com/libp2p/go-libp2p/core/network"
)

func QuaiProtocolHandler(stream network.Stream) {
	defer stream.Close()

	log.Debugf("Received a new stream from %s", stream.Conn().RemotePeer())

	// if there is a protocol mismatch, close the stream
	if stream.Protocol() != ProtocolVersion {
		log.Warnf("Invalid protocol: %s", stream.Protocol())
		// TODO: add logic to drop the peer
		return
	}

	msg, err := readQuaiMessage(stream)
	if err != nil {
		log.Warnf("Error reading message from stream: %s", err)
		return
	}
	log.Debugf("Received message: %+v", msg)

	switch msg.Flag {
	case joinFlag:
		// process join request
		if err := processJoinRequest(stream); err != nil {
			log.Warnf("Error processing join request: %s", err)
		}
	// TODO: handle other requests
	default:
		log.Warnf("invalid message: %+v", msg)
		// send error response
		responseMessage := QuaiProtocolMessage{
			Flag: errorFlag,
		}
		// TODO: add error message and/or error code

		if err := writeQuaiMessage(stream, &responseMessage); err != nil {
			log.Warnf("error writing error response to stream: %s", err)
		}
		log.Debugf("Sent error response: %+v", responseMessage)
	}

}

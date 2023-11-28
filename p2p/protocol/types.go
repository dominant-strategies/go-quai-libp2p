package protocol

import (
	"bufio"
	"encoding/json"

	"github.com/libp2p/go-libp2p/core/network"
)

// TODO: evaluate further extending the message structure.
// TODO: Should be use protobuf instead of json?
// Structure of the Quai protocol message to communicate between nodes using a stream
type QuaiProtocolMessage struct {
	Flag      byte   `json:"flag"`
	Data      []byte `json:"data"`
	ErrorCode int    `json:"error_code"`
}

const (
	// timeout in seconds before a read/write operation on the stream is considered failed
	TIMEOUT = 10
)

// Reads the message from the stream and returns a QuaiProtocolMessage
func readQuaiMessage(stream network.Stream) (*QuaiProtocolMessage, error) {
	// TODO: should we set a deadline for the read operation?
	// err := stream.SetDeadline(time.Now().Add(TIMEOUT * time.Second))
	// if err != nil {
	// 	return nil, err
	// }
	buf := bufio.NewReader(stream)
	requestJSON, err := buf.ReadString('\n')
	if err != nil {
		return nil, err
	}
	// remove newline from end of message
	requestJSON = requestJSON[:len(requestJSON)-1]

	var requestMessage QuaiProtocolMessage
	err = json.Unmarshal([]byte(requestJSON), &requestMessage)
	if err != nil {
		return nil, err
	}
	return &requestMessage, nil
}

// Writes the given QuaiProtocolMessage to the stream
func writeQuaiMessage(stream network.Stream, message *QuaiProtocolMessage) error {
	// TODO: should we set a deadline for the write operation?
	// err := stream.SetDeadline(time.Now().Add(TIMEOUT * time.Second))
	// if err != nil {
	// 	return err
	// }
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}
	// append newline to message
	messageJSON = append(messageJSON, '\n')
	_, err = stream.Write(messageJSON)
	if err != nil {
		return err
	}
	return nil
}

package protocol

import (
	"bufio"
	"encoding/json"
	"time"

	"github.com/pkg/errors"

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
	TIMEOUT = 10 * time.Second
)

// Reads the message from the stream and returns a QuaiProtocolMessage
func readQuaiMessage(stream network.Stream) (*QuaiProtocolMessage, error) {
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

// helper function to read a QuaiProtocolMessage with a timeout
func readQuaiMessageWithTimeout(stream network.Stream) (*QuaiProtocolMessage, error) {
	msgChan := make(chan *QuaiProtocolMessage, 1)
	errChan := make(chan error, 1)

	go func() {
		msg, err := readQuaiMessage(stream)
		if err != nil {
			errChan <- err
			return
		}
		msgChan <- msg
	}()

	select {
	case msg := <-msgChan:
		return msg, nil // Message received
	case err := <-errChan:
		return nil, err // Error occurred
	case <-time.After(TIMEOUT):
		return nil, errors.New("timeout waiting for message")
	}
}

// Writes the given QuaiProtocolMessage to the stream
func writeQuaiMessage(stream network.Stream, message *QuaiProtocolMessage) error {
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

// helper function to write a QuaiProtocolMessage with a timeout
func writeQuaiMessageWithTimeout(stream network.Stream, message *QuaiProtocolMessage) error {
	errChan := make(chan error, 1)

	go func() {
		err := writeQuaiMessage(stream, message)
		if err != nil {
			errChan <- err
			return
		}
	}()

	select {
	case err := <-errChan:
		return err
	case <-time.After(TIMEOUT):
		return errors.New("timeout waiting for message")
	}
}

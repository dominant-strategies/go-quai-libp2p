package common

import (
	"encoding/binary"
	"io"
	"time"

	"github.com/dominant-strategies/go-quai/log"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/pkg/errors"
)

const (
	// timeout in seconds before a read/write operation on the stream is considered failed
	// TODO: consider making this dynamic based on the network latency
	STREAM_TIMEOUT = 10 * time.Second

	// prefix length size in bytes
	LENGTH_PREFIX_SIZE = 4
)

// Reads the message from the stream and returns a byte of data.
func ReadMessageFromStream(stream network.Stream) ([]byte, error) {
	// Set a read deadline
	// err := stream.SetReadDeadline(time.Now().Add(STREAM_TIMEOUT))
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to set read deadline")
	// }

	// Read the length first
	lengthBuffer := make([]byte, LENGTH_PREFIX_SIZE)
	if _, err := io.ReadFull(stream, lengthBuffer); err != nil {
		return nil, errors.Wrap(err, "failed to read length from stream")
	}
	dataLength := binary.BigEndian.Uint32(lengthBuffer)

	// Read the data
	msg := make([]byte, dataLength)
	if _, err := io.ReadFull(stream, msg); err != nil {
		return nil, errors.Wrap(err, "failed to read message from stream")
	}
	log.Tracef("succesfully read %d bytes from stream", dataLength)
	return msg, nil
}

// Writes the message to the stream.
func WriteMessageToStream(stream network.Stream, msg []byte) error {
	// err := stream.SetWriteDeadline(time.Now().Add(STREAM_TIMEOUT))
	// if err != nil {
	// 	return errors.Wrap(err, "failed to set write deadline")
	// }
	msg = addLengthPrefix(msg)
	b, err := stream.Write(msg)
	if err != nil {
		return errors.Wrap(err, "failed to write message to stream")
	}
	log.Tracef("succesfully wrote %d bytes to stream", b)
	return err
}

// Adds a length prefix to the data.
func addLengthPrefix(data []byte) []byte {
	dataLength := len(data)
	lengthBuffer := make([]byte, LENGTH_PREFIX_SIZE)
	binary.BigEndian.PutUint32(lengthBuffer, uint32(dataLength))
	return append(lengthBuffer, data...)
}

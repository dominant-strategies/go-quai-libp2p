package protocol

import (
	"context"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuaiProtocolHandler(t *testing.T) {

	// Create a mock network and two hosts
	mockedNetwork := mocknet.New()

	host1, err := mockedNetwork.GenPeer()
	require.NoError(t, err)
	defer host1.Close()

	host2, err := mockedNetwork.GenPeer()
	require.NoError(t, err)
	defer host2.Close()

	// Connect the two hosts on the mock network
	err = mockedNetwork.LinkAll()
	require.NoError(t, err)

	tests := []struct {
		name              string
		ProtocolVersion   protocol.ID
		Message           QuaiProtocolMessage
		ExpectedStreamErr bool
		ExpectedFlag      byte
	}{
		{
			name:            "join network success",
			ProtocolVersion: ProtocolVersion,
			Message: QuaiProtocolMessage{
				Flag: joinFlag,
			},
			ExpectedStreamErr: false,
			ExpectedFlag:      welcomeFlag,
		},
		{
			name:            "invalid protocol",
			ProtocolVersion: "invalid",
			Message: QuaiProtocolMessage{
				Flag: joinFlag,
			},
			ExpectedStreamErr: true,
		},
		{
			name:            "invalid flag",
			ProtocolVersion: ProtocolVersion,
			Message: QuaiProtocolMessage{
				Flag: 0x00,
			},
			ExpectedStreamErr: false,
			ExpectedFlag:      errorFlag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Register protocol handler on host2
			host2.SetStreamHandler(ProtocolVersion, QuaiProtocolHandler)

			// Establish a stream from host1 (sender) to host2 (receiver)
			stream, err := host1.NewStream(ctx, peer.ID(host2.ID()), tt.ProtocolVersion)
			if tt.ExpectedStreamErr {
				assert.Error(t, err)
				return
			}
			defer stream.Close()

			// Send a message to host2
			err = writeQuaiMessage(stream, &tt.Message)
			assert.NoError(t, err)

			response, err := readQuaiMessage(stream)
			assert.NoError(t, err)
			assert.Equal(t, tt.ExpectedFlag, response.Flag)
		})
	}

}

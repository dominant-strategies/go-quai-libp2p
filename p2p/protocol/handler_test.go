package protocol

import (
	"context"
	"testing"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/stretchr/testify/assert"
)

func TestQuaiProtocolHandler3(t *testing.T) {
	ctx := context.Background()

	// Create a mock network with 1 bootstrap node
	mnet := generateMockNetwork(t, 1)
	bootstrapNode := mnet.Hosts()[0]
	bootstrapNode.SetStreamHandler(ProtocolVersion, QuaiProtocolHandler)

	// add peer to mock network with a private key
	clientNode := generateTestPeer(t, mnet)

	tests := []struct {
		name                  string
		ProtocolVersion       protocol.ID
		JoinMessage           QuaiProtocolMessage
		Signer                host.Host
		ChallengeResponseFlag byte
		ExpectStreamClose     bool
		ExpectJoinSuccess     bool
	}{
		{
			name:            "join network success",
			ProtocolVersion: ProtocolVersion,
			JoinMessage: QuaiProtocolMessage{
				Flag: joinFlag,
			},
			Signer:                clientNode,
			ChallengeResponseFlag: challengeResponseFlag,
			ExpectStreamClose:     false,
			ExpectJoinSuccess:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Establish a stream from clientNode (sender) to bootstrapNode (receiver)
			stream, err := clientNode.NewStream(ctx, peer.ID(bootstrapNode.ID()), tt.ProtocolVersion)
			if tt.ExpectStreamClose {
				// assert error when creating stream
				assert.Error(t, err)
				return
			}
			defer stream.Close()
			err = writeQuaiMessage(stream, &tt.JoinMessage)
			// assert no error when writing join message to stream
			assert.NoError(t, err)
			// Read response from bootstrap node with challenge
			response, err := readQuaiMessage(stream)
			assert.NoError(t, err)

			// Sign the challenge and send it back to the bootstrap node
			signature := signChallenge(t, response.Data, tt.Signer)

			// Send challenge response to bootstrap node
			challengeResponse := QuaiProtocolMessage{
				Flag: challengeResponseFlag,
				Data: signature,
			}
			err = writeQuaiMessage(stream, &challengeResponse)
			assert.NoError(t, err)

			// Read response from bootstrap node with welcome flag
			response, err = readQuaiMessage(stream)
			assert.NoError(t, err)
			// assert welcome flag in response
			assert.Equal(t, welcomeFlag, response.Flag)

		})
	}
}

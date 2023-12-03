package protocol

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessJoinRequest(t *testing.T) {
	ctx := context.Background()

	// Create a mock network with 1 bootstrap node
	mnet := generateMockNetwork(t, 1)
	bootstrapNode := mnet.Hosts()[0]
	// Set up the bootstrap node to handle join requests
	bootstrapNode.SetStreamHandler(ProtocolVersion, func(stream network.Stream) {
		processJoinRequest(stream)
	})

	// add peer to mock network with a private key
	clientNode := generateTestPeer(t, mnet)

	// add another (bad) peer to mock network with a private key
	badNode := generateTestPeer(t, mnet)

	tests := []struct {
		name                  string
		challengeResponseFlag byte
		clientNode            host.Host
		wantErrResp           bool
	}{
		{
			name:                  "join network success",
			challengeResponseFlag: challengeResponseFlag,
			clientNode:            clientNode,
			wantErrResp:           false,
		},
		{
			name:                  "sign message with wrong private key",
			challengeResponseFlag: challengeResponseFlag,
			clientNode:            badNode,
			wantErrResp:           true,
		},
		{
			name:                  "invalid challenge response flag",
			challengeResponseFlag: 0x00,
			clientNode:            clientNode,
			wantErrResp:           true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream, err := clientNode.NewStream(ctx, bootstrapNode.ID(), ProtocolVersion)
			require.NoError(t, err)
			defer stream.Close()

			// 1. client should receive a QuaiProtocolMessage with the challenge
			challengeMsg, err := readQuaiMessage(stream)
			assert.NoError(t, err)
			assert.Equal(t, challengeFlag, challengeMsg.Flag)

			// 2. client should send a QuaiProtocolMessage with the signed challenge
			// Sign the challenge
			signature := signChallenge(t, challengeMsg.Data, tt.clientNode)

			// Send signature to bootstrap node
			challengeResponseMsg := QuaiProtocolMessage{
				Flag: tt.challengeResponseFlag,
				Data: signature,
			}
			err = writeQuaiMessage(stream, &challengeResponseMsg)
			assert.NoError(t, err)

			// 3. client should receive a QuaiProtocolMessage with the welcome flag
			respMsg, err := readQuaiMessage(stream)
			assert.NoError(t, err)
			if tt.wantErrResp {
				assert.NotEqual(t, welcomeFlag, respMsg.Flag)
				return
			}
			assert.Equal(t, welcomeFlag, respMsg.Flag)
		})
	}

}

func TestVerifySignature(t *testing.T) {

	privKey, pubKey, err := crypto.GenerateEd25519Key(rand.Reader)
	require.NoError(t, err)

	nonce, err := createChallenge()
	require.NoError(t, err)
	hash := sha256.Sum256(nonce)

	signature, err := privKey.Sign(hash[:])
	require.NoError(t, err)

	verified, err := verifySignature(nonce, signature, pubKey)
	assert.NoError(t, err, "Signature verification should succeed")
	assert.True(t, verified, "Signature should be verified")

	// Test with wrong signature
	verified, err = verifySignature(nonce, []byte("wrong signature"), pubKey)
	assert.NoError(t, err, "Signature verification should succeed")
	assert.False(t, verified, "Signature should not be verified")

	// Test with wrong public key
	_, pubKey2, err := crypto.GenerateEd25519Key(rand.Reader)
	require.NoError(t, err)
	verified, err = verifySignature(nonce, signature, pubKey2)
	assert.NoError(t, err, "Signature verification should succeed")
	assert.False(t, verified, "Signature should not be verified")

}

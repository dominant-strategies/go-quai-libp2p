package protocol

import (
	"crypto/rand"
	"crypto/sha256"
	"os"
	"testing"

	"github.com/dominant-strategies/go-quai/log"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Comment / un comment below to see log output while testing
	log.ConfigureLogger(log.WithNullLogger())
	// log.ConfigureLogger(log.WithLevel("debug"))
	os.Exit(m.Run())
}

// creates a new peer and adds it to the mock network.
func generateTestPeer(t *testing.T, mnet mocknet.Mocknet) host.Host {
	t.Helper()
	privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
	require.NoError(t, err)
	clientMultiAddr, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/4001")
	require.NoError(t, err)
	testPeer, err := mnet.AddPeer(privKey, clientMultiAddr)
	require.NoError(t, err)
	err = mnet.LinkAll()
	require.NoError(t, err)
	return testPeer
}

// creates a mock network with the specified number of bootnodes
func generateMockNetwork(t *testing.T, bootnodes int) mocknet.Mocknet {
	t.Helper()
	mnet, err := mocknet.FullMeshLinked(bootnodes)
	require.NoError(t, err)
	return mnet
}

// helper function to sign the challenge using the host's private key
func signChallenge(t *testing.T, nonce []byte, host host.Host) []byte {
	t.Helper()
	hash := sha256.Sum256(nonce)
	sig, err := host.Peerstore().PrivKey(host.ID()).Sign(hash[:])
	require.NoError(t, err)
	return sig
}

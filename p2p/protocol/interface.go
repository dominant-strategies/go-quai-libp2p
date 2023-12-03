package protocol

import (
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
)

// interface required to join the quai protocol network
type QuaiP2PNode interface {
	GetBootPeers() []peer.AddrInfo
	Connect(pi peer.AddrInfo) error
	NewStream(peerID peer.ID, protocolID protocol.ID) (network.Stream, error)
	Network() network.Network
	// Signs the sha256 hash of the nonce and returns the signature
	SignChallenge(nonce []byte) (signature []byte, err error)
}

package node

import (
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/consensus/types"
	"github.com/dominant-strategies/go-quai/p2p/pb"
	"github.com/dominant-strategies/go-quai/p2p/protocol"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/pkg/errors"
)

// Opens a stream to the given peer and requests a block for the given hash and slice.
//
// If a block is not found, an error is returned
func (p *P2PNode) requestBlockFromPeer(hash types.Hash, slice types.SliceID, peerID peer.ID) (*types.Block, error) {
	// Open a stream to the peer using a specific protocol for block requests
	stream, err := p.NewStream(peerID, protocol.ProtocolVersion)
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	// create a block request protobuf message
	blockReq, err := pb.CreateProtoBlockRequest(hash, slice)
	if err != nil {
		return nil, err
	}

	// Send the block request to the peer
	err = common.WriteMessageToStream(stream, blockReq)
	if err != nil {
		return nil, err
	}

	// Read the response from the peer
	blockResponse, err := common.ReadMessageFromStream(stream)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response into a block
	found, pbBlock, err := pb.UnmarshalProtoBlockResponse(blockResponse)
	if err != nil {
		return nil, err
	}

	// If the block was found, return it
	if found {
		var block *types.Block
		block.FromProto(pbBlock)
		return block, nil
	}		

	// If the response does not contain a block, return an error
	return nil, errors.New("block not found")
}

// Creates a Cid from a slice ID to be used as DHT key
func shardToCid(slice types.SliceID) cid.Cid {
	sliceBytes := []byte(slice.String())

	// create a multihash from the slice ID
	mhash, _ := multihash.Encode(sliceBytes, multihash.SHA2_256)

	// create a Cid from the multihash
	return cid.NewCidV1(cid.Raw, mhash)

}

package types

import (
	"github.com/dominant-strategies/go-quai/p2p/pb"
)

type Block struct {
	Hash Hash
}

// Implementation of the pb.ProtoConvertable for Block

// converts a custom go Block type (types.Block) to a protocol buffer Block type (pb.Block)
func (b *Block) ToProto() *pb.Block {
	return &pb.Block{
		Hash: b.Hash.ToProto(),
		// TODO: map other fields
    }
}

// Implementation of the pb.ConvertableFromProto for Block

// converts a protocol buffer Block type (pb.Block) to a custom go Block type (types.Block)
func (b *Block) FromProto(pbBlock *pb.Block) {
	b.Hash.FromProto(pbBlock.Hash)
}

// returns a new instance of the protocol buffer Block type (pb.Block)
func (b *Block) NewProtoInstance() *pb.Block {
	return &pb.Block{}
}


type Slice struct {
	SliceID SliceID
}
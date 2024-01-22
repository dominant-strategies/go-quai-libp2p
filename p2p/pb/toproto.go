package pb

import (
	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/core/types"
)

// Converts a custom Block type to a protobuf Block type
func convertBlockToProto(block *types.Block) *Block {
	panic("TODO: implement")
}

// Converts a custom Header type to a protobuf Header type
func convertHeaderToProto(header *types.Header) *Header {
	if header == nil {
		return nil
	}

	protoHeader := &Header{
		UncleHash:     header.UncleHash().Bytes(),
		Coinbase:      header.Coinbase().Bytes(),
		Root:          header.Root().Bytes(),
		TxHash:        header.TxHash().Bytes(),
		EtxHash:       header.EtxHash().Bytes(),
		EtxRollupHash: header.EtxRollupHash().Bytes(),
		ReceiptHash:   header.ReceiptHash().Bytes(),
		Difficulty:    header.Difficulty().Bytes(),
		GasLimit:      header.GasLimit(),
		GasUsed:       header.GasUsed(),
		BaseFee:       header.BaseFee().Bytes(),
		Location:      header.Location(),
		Time:          header.Time(),
		Extra:         header.Extra(),
		MixHash:       header.MixHash().Bytes(),
		Nonce:         header.Nonce().Bytes(),
	}
	// Convert []*big.Int fields with HierarchyDepth
	for i := 0; i < common.HierarchyDepth; i++ {
		if header.ParentEntropy(i) != nil {
			protoHeader.ParentEntropy = append(protoHeader.ParentEntropy, header.ParentEntropy(i).Bytes())
		}
		if header.ParentDeltaS(i) != nil {
			protoHeader.ParentDeltaS = append(protoHeader.ParentDeltaS, header.ParentDeltaS(i).Bytes())
		}

		if header.Number(i) != nil {
			protoHeader.Number = append(protoHeader.Number, header.Number(i).Bytes())
		}
	}

	// Convert array of parent hashes
	parentHash := make([][]byte, len(header.ParentHashArray()))
	for i, hash := range header.ParentHashArray() {
		parentHash[i] = hash.Bytes()
	}
	protoHeader.ParentHash = parentHash

	// Convert array of manifest hashes
	manifestHash := make([][]byte, len(header.ManifestHashArray()))
	for i, hash := range header.ManifestHashArray() {
		manifestHash[i] = hash.Bytes()
	}
	protoHeader.ManifestHash = manifestHash

	return protoHeader
}

// Converts a custom Transaction type to a protobuf Transaction type
func convertTransactionToProto(transaction *types.Transaction) *Transaction {
	panic("TODO: implement")

}

// Converts a custom Block type to a protobuf Block type
func convertHashToProto(hash common.Hash) *Hash {
	hashBytes := hash.Bytes()
	protoHash := &Hash{
		Hash: hashBytes[:],
	}
	return protoHash
}

// Converts a custom Location type to a protobuf Location type
func convertLocationToProto(location common.Location) *Location {
	protoLocation := Location{
		Location: location,
	}
	return &protoLocation
}

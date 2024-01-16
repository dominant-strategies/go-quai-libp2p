package pb

import (
	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/core/types"
)

// Converts a protobuf location type to a custom location type
func convertProtoToLocation(protoLocation *Location) (*common.Location, error) {
	location := common.Location(protoLocation.GetLocation())
	return &location, nil
}

// Converts a protobuf Hash type to a custom Hash type
func convertProtoToHash(protoHash *Hash) (*common.Hash, error) {
	hash := &common.Hash{}
	hash.SetBytes(protoHash.Hash)
	// TODO: implement
	return hash, nil
}

// Converts a protobuf Block type to a custom Block type
func convertProtoToBlock(protoBlock *Block) (*types.Block, error) {

	txs := make([]*types.Transaction, len(protoBlock.Txs))
	for i, protoTransaction := range protoBlock.Txs {
		transaction, err := convertProtoToTransaction(protoTransaction)
		if err != nil {
			return nil, err
		}
		txs[i] = transaction
	}

	block := types.NewBlock(
		nil,
		txs,
		nil,
		nil,
		nil,
		nil,
		nil,
		0,
	)
	//! TODO: implement
	return block, nil
}

// Converts a protobuf Header type to a custom Header type
func convertProtoToHeader(protoHeader *Header) (*types.Header, error) {
	header := new(types.Header)
	header.SetGasLimit(protoHeader.GasLimit)
	header.SetGasUsed(protoHeader.GasUsed)
	// TODO: implement
	return header, nil
}

// Converts a protobuf Transaction type to a custom Transaction type
func convertProtoToTransaction(protoTransaction *Transaction) (*types.Transaction, error) {
	panic("TODO: implement")
	// return nil, nil
}

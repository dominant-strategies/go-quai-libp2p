package pb

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/core/types"
)

// Converts a protobuf location type to a custom location type
func convertProtoToLocation(protoLocation *Location) common.Location {
	location := common.Location(protoLocation.GetLocation())
	return location
}

// Converts a protobuf Hash type to a custom Hash type
func convertProtoToHash(protoHash *Hash) common.Hash {
	hash := common.Hash{}
	hash.SetBytes(protoHash.Hash)
	return hash
}

// Converts a protobuf Block type to a custom Block type
func convertProtoToBlock(protoBlock *Block) *types.Block {
	panic("TODO: implement")
}

// Converts a protobuf Header type to a custom Header type
func convertProtoToHeader(protoHeader *Header) *types.Header {
	if protoHeader == nil {
		return nil
	}

	header := types.EmptyHeader()

	// Set array of ParentHash
	for i, protoBytes := range protoHeader.ParentHash {
		header.SetParentHash(newHashFromBytes(protoBytes), i)
	}

	// Set UncleHash
	header.SetUncleHash(newHashFromBytes(protoHeader.UncleHash))

	// Set Coinbase
	coinbase, err := newAddressFromBytes(protoHeader.Coinbase)
	if err != nil {
		panic("TODO: implement")
	}
	header.SetCoinbase(*coinbase)

	// Set Root
	header.SetRoot(newHashFromBytes(protoHeader.Root))

	// Set TxHash
	header.SetTxHash(newHashFromBytes(protoHeader.TxHash))

	// Set EtxHash
	header.SetEtxHash(newHashFromBytes(protoHeader.EtxHash))

	// Set EtxRollupHash
	header.SetEtxRollupHash(newHashFromBytes(protoHeader.EtxRollupHash))

	// Set ManifestHash
	for i, protoBytes := range protoHeader.ManifestHash {
		header.SetManifestHash(newHashFromBytes(protoBytes), i)
	}

	// Set ReceiptHash
	header.SetReceiptHash(newHashFromBytes(protoHeader.ReceiptHash))

	// Set Difficulty
	header.SetDifficulty(big.NewInt(0).SetBytes(protoHeader.Difficulty))

	// Set ParentEntropy
	for i, protoBytes := range protoHeader.ParentEntropy {
		header.SetParentEntropy(big.NewInt(0).SetBytes(protoBytes), i)
	}

	// Set ParentDeltaS
	for i, protoBytes := range protoHeader.ParentDeltaS {
		header.SetParentDeltaS(big.NewInt(0).SetBytes(protoBytes), i)
	}

	// Set Number
	for i, protoBytes := range protoHeader.Number {
		header.SetNumber(big.NewInt(0).SetBytes(protoBytes), i)
	}

	// Set GasLimit
	header.SetGasLimit(protoHeader.GasLimit)

	// Set GasUsed
	header.SetGasUsed(protoHeader.GasUsed)

	// Set BaseFee
	header.SetBaseFee(big.NewInt(0).SetBytes(protoHeader.BaseFee))

	// Set Location
	header.SetLocation(common.Location(protoHeader.Location))

	// Set Time
	header.SetTime(protoHeader.Time)

	// Set Extra
	header.SetExtra(protoHeader.Extra)

	// Set MixHash
	header.SetMixHash(newHashFromBytes(protoHeader.MixHash))

	// Set Nonce
	if len(protoHeader.Nonce) == 8 {
		nonce := types.BlockNonce(protoHeader.Nonce)
		header.SetNonce(nonce)
	}

	return header
}

// Converts a protobuf Transaction type to a custom Transaction type
func convertProtoToTransaction(protoTransaction *Transaction) *types.Transaction {
	panic("TODO: implement")
}

// Helper function to convert a slice of bytes to a common.Hash
func newHashFromBytes(hashBytes []byte) common.Hash {
	hash := common.Hash{}
	hash.SetBytes(hashBytes)
	return hash
}

// Helper function to convert a slice of bytes to a *common.Address
func newAddressFromBytes(addressBytes []byte) (*common.Address, error) {
	if len(addressBytes) == 0 {
		return &common.Address{}, nil
	}
	hexEncoded := hex.EncodeToString(addressBytes)
	hexEncoded = "0x" + hexEncoded
	address := common.Address{}
	err := address.UnmarshalText([]byte(hexEncoded))
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling address: %v", err)
	}
	return &address, nil
}

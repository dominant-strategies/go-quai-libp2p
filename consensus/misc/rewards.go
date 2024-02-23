package misc

import (
	"math/big"

	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/core/types"
)

// CalculateReward calculates the coinbase rewards depending on the type of the block
func CalculateReward(header *types.Header) *big.Int {
	if header.Coinbase().IsInQiLedgerScope() {
		return calculateQiReward(header)
	} else {
		return calculateQuaiReward(header)
	}
}

// CalculateQuaiReward calculates the quai that can be recieved for mining a block
func calculateQuaiReward(header *types.Header) *big.Int {
	return common.Big0
}

// CalculateQiReward caculates the qi that can be received for mining a block
func calculateQiReward(header *types.Header) *big.Int {
	return common.Big0
}

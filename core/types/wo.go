package types

import (
	"math/big"

	"github.com/dominant-strategies/go-quai/common"
)

type workObject struct {
	proposedHeader PendingHeader
	parentHash     common.Hash
	parentNumber   big.Int
	difficulty     big.Int
	txHash         common.Hash
	nonce          BlockNonce
}

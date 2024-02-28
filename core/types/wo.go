package types

import (
	"math/big"

	"github.com/dominant-strategies/go-quai/common"
)

type WorkObject struct {
	woHeader WorkObjectHeader
	woBlock  WorkObjectBlock
	tx       Transaction
}

type WorkObjectHeader struct {
	proposedHeaderHash common.Hash
	parentHash         common.Hash
	parentNumber       *big.Int
	difficulty         *big.Int
	txHash             common.Hash
	nonce              BlockNonce
}

type WorkObjectBlock struct {
	header          *Header
	transactions    Transactions
	extTransactions Transactions
	uncles          []*Header
	manifest        BlockManifest
}

func (wo *WorkObject) Header() *Header {
	return wo.woBlock.header
}

func (wo *WorkObject) SetProposedHeader(header *Header) {
	wo.proposedHeader = CopyHeader(header)
}

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
	HeaderHash   common.Hash
	parentHash   common.Hash
	parentNumber *big.Int
	difficulty   *big.Int
	txHash       common.Hash
	nonce        BlockNonce
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

func (wo *WorkObject) Transactions() Transactions {
	return wo.woBlock.transactions
}

func (wo *WorkObject) ExtTransactions() Transactions {
	return wo.woBlock.extTransactions
}

func (wo *WorkObject) Uncles() []*Header {
	return wo.woBlock.uncles
}

func (wo *WorkObject) Manifest() BlockManifest {
	return wo.woBlock.manifest
}

func (wo *WorkObject) ParentHash() common.Hash {
	return wo.woHeader.parentHash
}

func (wo *WorkObject) ParentNumber() *big.Int {
	return wo.woHeader.parentNumber
}

func (wo *WorkObject) Difficulty() *big.Int {
	return wo.woHeader.difficulty
}

func (wo *WorkObject) TxHash() common.Hash {
	return wo.woHeader.txHash
}

func (wo *WorkObject) Nonce() BlockNonce {
	return wo.woHeader.nonce
}

func (wo *WorkObject) HeaderHash() common.Hash {
	return wo.woHeader.HeaderHash
}

func (wo *WorkObject) Tx() Transaction {
	return wo.tx
}

func (wo *WorkObject) SetTx(tx Transaction) {
	wo.tx = tx
}

func (wo *WorkObject) SetHeader(header *Header) {
	wo.woBlock.header = header
}

func (wo *WorkObject) SetTransactions(transactions Transactions) {
	wo.woBlock.transactions = transactions
}

func (wo *WorkObject) SetExtTransactions(transactions Transactions) {
	wo.woBlock.extTransactions = transactions
}

func (wo *WorkObject) SetUncles(uncles []*Header) {
	wo.woBlock.uncles = uncles
}

func (wo *WorkObject) SetManifest(manifest BlockManifest) {
	wo.woBlock.manifest = manifest
}

func (wo *WorkObject) SetParentHash(parentHash common.Hash) {
	wo.woHeader.parentHash = parentHash
}

func (wo *WorkObject) SetParentNumber(parentNumber *big.Int) {
	wo.woHeader.parentNumber = parentNumber
}

func (wo *WorkObject) SetDifficulty(difficulty *big.Int) {
	wo.woHeader.difficulty = difficulty
}

func (wo *WorkObject) SetTxHash(txHash common.Hash) {
	wo.woHeader.txHash = txHash
}

func (wo *WorkObject) SetNonce(nonce BlockNonce) {
	wo.woHeader.nonce = nonce
}

func (wo *WorkObject) SetHeaderHash(headerHash common.Hash) {
	wo.woHeader.HeaderHash = headerHash
}

func NewWorkObject(header *Header, transactions Transactions, extTransactions Transactions, uncles []*Header, manifest BlockManifest, parentHash common.Hash, parentNumber *big.Int, difficulty *big.Int, txHash common.Hash, nonce BlockNonce, headerHash common.Hash, tx Transaction) *WorkObject {
	return &WorkObject{
		woHeader: WorkObjectHeader{
			HeaderHash:   headerHash,
			parentHash:   parentHash,
			parentNumber: parentNumber,
			difficulty:   difficulty,
			txHash:       txHash,
			nonce:        nonce,
		},
		woBlock: WorkObjectBlock{
			header:          CopyHeader(header),
			transactions:    transactions,
			extTransactions: extTransactions,
			uncles:          uncles,
			manifest:        manifest,
		},
		tx: tx,
	}
}

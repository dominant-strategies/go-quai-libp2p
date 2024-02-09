package types

import (
	"errors"
	"math/big"

	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/log"
)

type ExternalTx struct {
	ChainID           *big.Int
	OriginatingTxHash common.Hash
	ETXIndex          uint16
	Gas               uint64
	To                *common.Address `rlp:"nilString"` // nil means contract creation
	Value             *big.Int
	Data              []byte
	AccessList        AccessList
	Sender            common.Address

	// External transactions do not have signatures. The origin chain will
	// emit an ETX, and consequently 'authorization' of this transaction comes
	// from chain consensus and not from an account signature.
	//
	// Before an ETX can be processed at the destination chain, the ETX must
	// become referencable through block manifests, thereby guaranteeing that
	// the origin chain indeed confirmed emission of that ETX.
}

// PendingEtxs are ETXs which have been emitted in a subordinate block. The
// block is not valid in dominant chains, but dominant chains relay the pending
// ETXs to other chains in the network to facilitate ETX forward propagation.
//
// A dominant chain does not have the state to check correctness or acceptability
// of these ETXs in the subordinate chains, but it does need to know that these
// ETXs are valid against a block header which came from a subordinate chain.
// For this reason, we indlude a header from the subordinate chain.
type PendingEtxs struct {
	Header *Header `json:"header" gencodec:"required"`
	// Etxs array contains ETXs from the chain which produced this block, and a
	// subordinate rollup of ETXs for that chain's subordinate (if it has one).
	// Etxs[originCtx] = external transactions in origin CTX
	// (optional) Etxs[originCtx+1] = rollup of ETXs emitted by originCtx+1
	Etxs []Transactions `json:"etxs"   gencodec:"required"`
}

func (p *PendingEtxs) IsValid(hasher TrieHasher, nodeCtx int) bool {
	if p == nil || p.Header == nil || p.Etxs == nil {
		log.Global.WithField("p", p).Info("PendingEtx: p/p.Header/p.Etxs is nil")
		return false
	}
	if len(p.Etxs) < common.HierarchyDepth {
		return false
	}
	// pending ETXs must have originated from our subordinate context.
	singletonCtx := nodeCtx + 1
	rollupCtx := singletonCtx + 1
	// singletonCtx must exist and must match hash
	if singletonCtx >= len(p.Etxs) || DeriveSha(p.Etxs[singletonCtx], hasher) != p.Header.EtxHash() {
		return false
	}
	// rollupCtx may not exist (i.e. if we are a region node), but if it is, the rollup hash must match
	if rollupCtx < len(p.Etxs) && DeriveSha(p.Etxs[rollupCtx], hasher) != p.Header.EtxRollupHash() {
		return false
	}
	return true
}

// ProtoEncode encodes the PendingEtxs to protobuf format.
func (p *PendingEtxs) ProtoEncode() (*ProtoPendingEtxs, error) {
	header, err := p.Header.ProtoEncode()
	if err != nil {
		return nil, err
	}
	etxs := make([]*ProtoTransactions, len(p.Etxs))
	for i, pEtxs := range p.Etxs {
		etxs[i], err = pEtxs.ProtoEncode()
		if err != nil {
			return nil, err
		}
	}
	return &ProtoPendingEtxs{
		Header: header,
		Etxs:   etxs,
	}, nil
}

// ProtoDecode decodes the protobuf to a PendingEtxs representation.
func (p *PendingEtxs) ProtoDecode(protoPendingEtxs *ProtoPendingEtxs, location common.Location) error {
	if protoPendingEtxs.Header == nil {
		return errors.New("header is nil in ProtoDecode")
	}
	p.Header = new(Header)
	err := p.Header.ProtoDecode(protoPendingEtxs.GetHeader())
	if err != nil {
		return err
	}
	p.Etxs = make([]Transactions, len(protoPendingEtxs.Etxs))
	for i, protoEtxs := range protoPendingEtxs.Etxs {
		p.Etxs[i] = make(Transactions, len(protoEtxs.GetTransactions()))
		for j, protoTx := range protoEtxs.GetTransactions() {
			p.Etxs[i][j] = &Transaction{}
			err = p.Etxs[i][j].ProtoDecode(protoTx, location)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *ExternalTx) copy() TxData {
	cpy := &ExternalTx{
		To:                tx.To, // TODO: copy pointed-to address
		Data:              common.CopyBytes(tx.Data),
		Gas:               tx.Gas,
		OriginatingTxHash: tx.OriginatingTxHash,
		ETXIndex:          tx.ETXIndex,
		Sender:            tx.Sender,

		// These are copied below.
		AccessList: make(AccessList, len(tx.AccessList)),
		Value:      new(big.Int),
		ChainID:    new(big.Int),
	}
	copy(cpy.AccessList, tx.AccessList)
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	return cpy
}

// accessors for innerTx.
func (tx *ExternalTx) txType() byte                   { return ExternalTxType }
func (tx *ExternalTx) chainID() *big.Int              { return tx.ChainID }
func (tx *ExternalTx) protected() bool                { return true }
func (tx *ExternalTx) accessList() AccessList         { return tx.AccessList }
func (tx *ExternalTx) data() []byte                   { return tx.Data }
func (tx *ExternalTx) gas() uint64                    { return tx.Gas }
func (tx *ExternalTx) gasFeeCap() *big.Int            { return new(big.Int) } // placeholder
func (tx *ExternalTx) gasTipCap() *big.Int            { return new(big.Int) } // placeholder
func (tx *ExternalTx) gasPrice() *big.Int             { return new(big.Int) } // placeholder
func (tx *ExternalTx) value() *big.Int                { return tx.Value }
func (tx *ExternalTx) to() *common.Address            { return tx.To }
func (tx *ExternalTx) etxSender() common.Address      { return tx.Sender }
func (tx *ExternalTx) originatingTxHash() common.Hash { return tx.OriginatingTxHash }
func (tx *ExternalTx) etxIndex() uint16               { return tx.ETXIndex }
func (tx *ExternalTx) nonce() uint64                  { panic("external TX does not have nonce") }
func (tx *ExternalTx) etxGasLimit() uint64            { panic("external TX does not have etxGasLimit") }
func (tx *ExternalTx) etxGasPrice() *big.Int          { panic("external TX does not have etxGasPrice") }
func (tx *ExternalTx) etxGasTip() *big.Int            { panic("external TX does not have etxGasTip") }
func (tx *ExternalTx) etxData() []byte                { panic("external TX does not have etxData") }
func (tx *ExternalTx) etxAccessList() AccessList      { panic("external TX does not have etxAccessList") }

func (tx *ExternalTx) rawSignatureValues() (v, r, s *big.Int) {
	// Signature values are ignored for external transactions
	return nil, nil, nil
}

func (tx *ExternalTx) setSignatureValues(chainID, v, r, s *big.Int) {
	// Signature values are ignored for external transactions
}

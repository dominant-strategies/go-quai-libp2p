package types

import (
	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/log"
	"github.com/dominant-strategies/go-quai/params"
)

// The EtxSet maps an ETX hash to the ETX and block number in which it became available.
// If no entry exists for a given ETX hash, then that ETX is not available.
type EtxSet map[common.Hash]EtxSetEntry

type EtxSetEntry struct {
	Height uint64
	ETX    Transaction
}

func NewEtxSet() EtxSet {
	return make(EtxSet)
}

// updateInboundEtxs updates the set of inbound ETXs available to be mined into
// a block in this location. This method adds any new ETXs to the set and
// removes expired ETXs.
func (set *EtxSet) Update(newInboundEtxs Transactions, currentHeight uint64, nodeLocation common.Location) {
	// Add new ETX entries to the inbound set
	for _, etx := range newInboundEtxs {
		if etx.To().Location().Equal(nodeLocation) {
			(*set)[etx.Hash()] = EtxSetEntry{currentHeight, *etx}
		} else {
			panic("cannot add ETX destined to other chain to our ETX set")
		}
	}

	// Remove expired ETXs
	for txHash, entry := range *set {
		availableAtBlock := entry.Height
		etxExpirationHeight := availableAtBlock + params.EtxExpirationAge
		if currentHeight > etxExpirationHeight {
			log.Global.WithFields(log.Fields{
				"hash":                txHash,
				"gasTipCap":           entry.ETX.GasTipCap(),
				"gasFeeCap":           entry.ETX.GasFeeCap(),
				"gasLimit":            entry.ETX.Gas(),
				"availableAtBlock":    availableAtBlock,
				"etxExpirationHeight": etxExpirationHeight,
				"currentHeight":       currentHeight,
			}).Warn("ETX expired")
			delete(*set, txHash)
		}
	}
}

// ProtoEncode encodes the EtxSet to protobuf format.
func (set *EtxSet) ProtoEncode() *ProtoEtxSet {
	protoSet := &ProtoEtxSet{}
	for hash, entry := range *set {
		protoHash := hash.ProtoEncode()
		etxHeight := entry.Height
		etx, err := entry.ETX.ProtoEncode()
		if err != nil {
			panic(err)
		}
		protoSet.EtxSet = append(protoSet.EtxSet, &ProtoEtxSetEntry{
			EtxHash: protoHash,
			Height:  &etxHeight,
			Etx:     etx,
		})
	}
	return protoSet
}

// ProtoDecode decodes the EtxSet from protobuf format.
func (set *EtxSet) ProtoDecode(protoSet *ProtoEtxSet, location common.Location) error {
	for _, entry := range protoSet.GetEtxSet() {
		hash := &common.Hash{}
		hash.ProtoDecode(entry.GetEtxHash())
		etx := new(Transaction)
		err := etx.ProtoDecode(entry.GetEtx(), location)
		if err != nil {
			return err
		}
		(*set)[*hash] = EtxSetEntry{
			Height: entry.GetHeight(),
			ETX:    *etx,
		}
	}
	return nil
}

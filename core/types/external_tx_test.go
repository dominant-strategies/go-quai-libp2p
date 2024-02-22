package types

import (
	"testing"
)

func TestPendingEtxsValidity(t *testing.T) {
	pendingEtxs := &PendingEtxs{EmptyHeader(), make(Transactions, 0)}

	t.Log("Len of pendingEtxs", len(pendingEtxs.Etxs))

	pendingEtxs.IsValid(nil)
}

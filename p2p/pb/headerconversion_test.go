package pb

import (
	"crypto/rand"

	"math/big"
	"testing"

	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomHeader(t *testing.T) *types.Header {
	header := types.EmptyHeader()

	// Set array of ParentHash
	for i := 0; i < common.HierarchyDepth; i++ {
		header.SetParentHash(newRandomHash(), i)
	}

	// Set UncleHash
	header.SetUncleHash(newRandomHash())

	// Set Coinbase
	coinbase, err := newAddressFromBytes(newRandomHash().Bytes()[0:20])
	require.NoError(t, err)
	header.SetCoinbase(*coinbase)

	// Set Root
	header.SetRoot(newRandomHash())

	// Set TxHash
	header.SetTxHash(newRandomHash())

	// Set EtxHash
	header.SetEtxHash(newRandomHash())

	// Set EtxRollupHash
	header.SetEtxRollupHash(newRandomHash())

	// Set ManifestHash
	for i := 0; i < common.HierarchyDepth; i++ {
		header.SetManifestHash(newRandomHash(), i)
	}

	// Set ReceiptHash
	header.SetReceiptHash(newRandomHash())

	// Set Difficulty
	header.SetDifficulty(newRandomBigInt())

	// Set ParentEntropy, ParentDeltaS, and Number
	for i := 0; i < common.HierarchyDepth; i++ {
		header.SetParentEntropy(newRandomBigInt(), i)
		header.SetParentDeltaS(newRandomBigInt(), i)
		header.SetNumber(newRandomBigInt(), i)
	}

	// Set GasLimit and GasUsed
	header.SetGasLimit(newRandomBigInt().Uint64())
	header.SetGasUsed(newRandomBigInt().Uint64())

	// Set BaseFee
	header.SetBaseFee(newRandomBigInt())

	// Set Location
	location := common.Location{1, 2}
	header.SetLocation(location)

	// Set Time
	header.SetTime(newRandomBigInt().Uint64())

	// Set Extra
	header.SetExtra([]byte("extra"))

	// Set MixHash
	header.SetMixHash(newRandomHash())

	return header
}

// helper function to return a random hash
func newRandomHash() common.Hash {
	hashBytes := make([]byte, 32)
	rand.Read(hashBytes)
	hash := common.Hash{}
	hash.SetBytes(hashBytes)
	return hash
}

// helper function to return a random big int
func newRandomBigInt() *big.Int {
	bigInt := big.NewInt(0)
	bigInt.SetBytes(newRandomHash().Bytes())
	return bigInt
}

func TestHeaderConversion(t *testing.T) {
	tests := []struct {
		name        string
		inputHeader *types.Header
	}{
		{
			name:        "nil header",
			inputHeader: nil,
		},
		{
			name:        "empty header",
			inputHeader: types.EmptyHeader(),
		},
		{
			name:        "random header",
			inputHeader: createRandomHeader(t),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Convert to protobuf and back
			protoHeader := convertHeaderToProto(tc.inputHeader)
			resultHeader := convertProtoToHeader(protoHeader)

			compareHeaders(t, tc.inputHeader, resultHeader)

		})
	}
}

// helper function to compare two headers field by field
func compareHeaders(t *testing.T, h1, h2 *types.Header) {
	if h1 == nil && h2 == nil {
		return
	} else if h1 == nil || h2 == nil {
		assert.Fail(t, "One header is nil")
		return
	}

	for i := 0; i < common.HierarchyDepth; i++ {
		assert.Equal(t, h1.ParentHash(i), h2.ParentHash(i), "ParentHashes do not match")
		assert.Equal(t, h1.ParentEntropy(i), h2.ParentEntropy(i), "ParentEntropy do not match")
		assert.Equal(t, h1.ParentDeltaS(i), h2.ParentDeltaS(i), "ParentDeltaS do not match")
		assert.Equal(t, h1.Number(i), h2.Number(i), "Numbers do not match")
		assert.Equal(t, h1.ManifestHash(i), h2.ManifestHash(i), "ManifestHashes do not match")
	}

	assert.Equal(t, h1.UncleHash(), h2.UncleHash(), "UncleHashes do not match")
	assert.Equal(t, h1.Coinbase(), h2.Coinbase(), "Coinbases do not match")
	assert.Equal(t, h1.Root(), h2.Root(), "Roots do not match")
	assert.Equal(t, h1.TxHash(), h2.TxHash(), "TxHashes do not match")
	assert.Equal(t, h1.EtxHash(), h2.EtxHash(), "EtxHashes do not match")
	assert.Equal(t, h1.EtxRollupHash(), h2.EtxRollupHash(), "EtxRollupHashes do not match")
	assert.Equal(t, h1.ReceiptHash(), h2.ReceiptHash(), "ReceiptHashes do not match")
	assert.Equal(t, h1.Difficulty(), h2.Difficulty(), "Difficulties do not match")
	assert.Equal(t, h1.GasLimit(), h2.GasLimit(), "GasLimits do not match")
	assert.Equal(t, h1.GasUsed(), h2.GasUsed(), "GasUseds do not match")
	assert.Equal(t, h1.BaseFee(), h2.BaseFee(), "BaseFees do not match")
	assert.Equal(t, h1.Time(), h2.Time(), "Times do not match")
	assert.Equal(t, h1.Extra(), h2.Extra(), "Extras do not match")
	assert.Equal(t, h1.MixHash(), h2.MixHash(), "MixHashes do not match")
	assert.Equal(t, h1.Nonce(), h2.Nonce(), "Nonces do not match")
}

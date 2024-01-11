package pb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalUnmarshalProtoMessage(t *testing.T) {
	// Create a mock Block
	pbBlock := &Block{
		Hash: &Hash{Hash: []byte("mockHash")},
	}

	// Marshal the Block
	data, err := MarshalProtoMessage(pbBlock)
	assert.NoError(t, err)

	// Unmarshal the data back into a Block
	var unmarshaledBlock Block
	err = UnmarshalProtoMessage(data, &unmarshaledBlock)
	assert.NoError(t, err)

	// Marshal the unmarshaledBlock
	newData, err := MarshalProtoMessage(&unmarshaledBlock)
	assert.NoError(t, err)

	// Check if the marshaled data of both blocks are equal
	assert.Equal(t, data, newData)
}

type mockBlock struct {
	Hash [32]byte
}

func (m *mockBlock) ToProto() *Block {
	return &Block{
		Hash: &Hash{Hash: m.Hash[:]},
	}
}

func (m *mockBlock) FromProto(pbMsg *Block) {
	pbHash := pbMsg.Hash
	copy(m.Hash[:], pbHash.Hash)
}

func (m *mockBlock) NewProtoInstance() *Block {
	return &Block{}
}

func TestConvertAndMarshal(t *testing.T) {
	mockHash := [32]byte{0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c}
	mockBlock1 := &mockBlock{Hash: mockHash}

	// Convert the mockBlock to a protobuf Block and marshal it
	data, err := ConvertAndMarshal(mockBlock1)
	assert.NoError(t, err)

	// Unmarshal the data back into a protobuf Block and convert it to a mockBlock
	mock2 := &mockBlock{}
	err = UnmarshalAndConvert(data, mock2)
	assert.NoError(t, err)

	// Check if the hashes are equal
	assert.Equal(t, mockBlock1.Hash, mock2.Hash)

}

package pubsubManager

import (
	"context"
	"testing"
	"time"

	"github.com/dominant-strategies/go-quai/consensus/types"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPublishSubscribeBlock(t *testing.T) {
	ctx := context.Background()

	// Create a mock network
	mnet := mocknet.New()

	// Setup two nodes on the mock network
	host1, err := mnet.GenPeer()
	require.NoError(t, err)
	defer host1.Close()

	host2, err := mnet.GenPeer()
	require.NoError(t, err)
	defer host2.Close()

	// Connect the two nodes
	err = mnet.LinkAll()
	require.NoError(t, err)
	err = mnet.ConnectAllButSelf()
	require.NoError(t, err)

	// Setup GossipSub
	gm1, err := NewGossipSubManager(ctx, host1)
	require.NoError(t, err)
	gm2, err := NewGossipSubManager(ctx, host2)
	require.NoError(t, err)

	// Subscribe to the block topic on both nodes
	slice := types.SliceID{
		Region: 0,
		Zone:   0,
	}
	block := types.Block{}
	err = gm1.Subscribe(slice, block)
	assert.NoError(t, err)
	err = gm2.Subscribe(slice, block)
	assert.NoError(t, err)
	time.Sleep(time.Second) // Allow time for subscription to be established

	// define a callback function to handle received data via gossipsub
	dataChan := make(chan interface{})
	cb := func(data interface{}) {
		dataChan <- data
	}

	// Start the gossipsub manager on second node
	gm2.Start(cb)

	// Create a mock block to publish
	mockHash := types.Hash{0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c, 0x7e, 0x1c, 0x7c}
	mockBlock := types.Block{
		Hash: mockHash,
	}

	// Publish the block on the first node
	err = gm1.Broadcast(slice, mockBlock)
	assert.NoError(t, err)

	// Wait for the message to be received on the second node
	select {
	case msg := <-dataChan:
		// Check if the received message is a block
		receivedBlock, ok := msg.(types.Block)
		assert.True(t, ok)
		// Check if the received block is equal to the mock block
		assert.Equal(t, mockBlock, receivedBlock)

	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for message")
	}
}

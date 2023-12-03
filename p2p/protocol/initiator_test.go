package protocol

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dominant-strategies/go-quai/p2p/protocol/mocks"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/stretchr/testify/assert"
)

func TestJoinNetwork(t *testing.T) {
	// vars used in mock stubs
	ctx := context.Background()
	mockErr := fmt.Errorf("mock error")

	// Create a mock network with 2 bootnodes
	mnet := generateMockNetwork(t, 2)
	bootNodes := mnet.Hosts()

	// set protocol handler
	for _, bootNode := range bootNodes {
		bootNode.SetStreamHandler(ProtocolVersion, QuaiProtocolHandler)
	}

	bootNodeAddr1 := peer.AddrInfo{
		ID:    bootNodes[0].ID(),
		Addrs: bootNodes[0].Addrs(),
	}

	bootNodeAddr2 := peer.AddrInfo{
		ID:    bootNodes[1].ID(),
		Addrs: bootNodes[1].Addrs(),
	}

	testNode := generateTestPeer(t, mnet)

	tests := []struct {
		name     string
		mockStub func(*mocks.MockQuaiP2PNode)
		WantErr  bool
	}{
		{
			name: "join network successfully",
			mockStub: func(mockedQuaiNode *mocks.MockQuaiP2PNode) {
				mockedQuaiNode.EXPECT().GetBootPeers().Return([]peer.AddrInfo{bootNodeAddr1, bootNodeAddr2}).Times(1)
				mockedQuaiNode.EXPECT().Connect(gomock.Any()).Return(nil).Times(2)
				mockedQuaiNode.EXPECT().Network().Return(testNode.Network()).Times(1)
				mockedQuaiNode.EXPECT().NewStream(gomock.Eq(bootNodeAddr1.ID), gomock.Eq(ProtocolVersion)).Return(testNode.NewStream(ctx, bootNodeAddr1.ID, ProtocolVersion)).Times(1)
				mockedQuaiNode.EXPECT().NewStream(gomock.Eq(bootNodeAddr2.ID), gomock.Eq(ProtocolVersion)).Return(testNode.NewStream(ctx, bootNodeAddr2.ID, ProtocolVersion)).Times(1)
				mockedQuaiNode.EXPECT().SignChallenge(gomock.Any()).DoAndReturn(
					func(nonce []byte) ([]byte, error) {
						return signChallenge(t, nonce, testNode), nil
					}).Times(2)
			},
			WantErr: false,
		},
		{
			name: "connect fails with all bootnodes",
			mockStub: func(mockedQuaiNode *mocks.MockQuaiP2PNode) {
				mockedQuaiNode.EXPECT().GetBootPeers().Return([]peer.AddrInfo{bootNodeAddr1, bootNodeAddr2}).Times(1)
				mockedQuaiNode.EXPECT().Connect(gomock.Any()).Return(mockErr).Times(2)
			},
			WantErr: true,
		},
		{
			name: "new stream fails with 1 bootnode, succeeds with other",
			mockStub: func(mockedQuaiNode *mocks.MockQuaiP2PNode) {
				mockedQuaiNode.EXPECT().GetBootPeers().Return([]peer.AddrInfo{bootNodeAddr1, bootNodeAddr2}).Times(1)
				mockedQuaiNode.EXPECT().Connect(gomock.Any()).Return(nil).Times(2)
				mockedQuaiNode.EXPECT().Network().Return(testNode.Network()).Times(1)
				mockedQuaiNode.EXPECT().NewStream(gomock.Eq(bootNodeAddr1.ID), gomock.Eq(ProtocolVersion)).Return(nil, mockErr).Times(1)
				mockedQuaiNode.EXPECT().NewStream(gomock.Eq(bootNodeAddr2.ID), gomock.Eq(ProtocolVersion)).Return(testNode.NewStream(ctx, bootNodeAddr2.ID, ProtocolVersion)).Times(1)
				mockedQuaiNode.EXPECT().SignChallenge(gomock.Any()).DoAndReturn(
					func(nonce []byte) ([]byte, error) {
						return signChallenge(t, nonce, testNode), nil
					}).Times(1)
			},
			WantErr: false,
		},
		{
			name: "new stream fails with all bootnodes",
			mockStub: func(mockedQuaiNode *mocks.MockQuaiP2PNode) {
				mockedQuaiNode.EXPECT().GetBootPeers().Return([]peer.AddrInfo{bootNodeAddr1, bootNodeAddr2}).Times(1)
				mockedQuaiNode.EXPECT().Connect(gomock.Any()).Return(nil).Times(2)
				mockedQuaiNode.EXPECT().Network().Return(testNode.Network()).Times(1)
				mockedQuaiNode.EXPECT().NewStream(gomock.Any(), gomock.Any()).Return(nil, mockErr).Times(2)
			},
			WantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockedQuaiNode := mocks.NewMockQuaiP2PNode(ctrl)
			tt.mockStub(mockedQuaiNode)
			err := JoinNetwork(mockedQuaiNode)
			if tt.WantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

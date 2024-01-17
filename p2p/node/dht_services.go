package node

import (
	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/p2p/pubsubManager"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

// Creates a Cid from the topic to be used as DHT key
func topicToCid(topic string) cid.Cid {
	// convert the topic to a byte slice
	topicBytes := []byte(topic)

	// create a multihash from the slice ID
	mhash, _ := multihash.Encode(topicBytes, multihash.SHA2_256)

	// create a Cid from the multihash
	return cid.NewCidV1(cid.Raw, mhash)

}

// Announces in the DHT that this node is a provider of the given data type
// with the topic derived from the slice ID and data type.
func (p *P2PNode) announceToDHT(location common.Location, dataType interface{}, data interface{}) error {
	topic, err := pubsubManager.TopicName(location, dataType)
	if err != nil {
		return err
	}
	cid := topicToCid(topic)
	return p.dht.Provide(p.ctx, cid, true)

}

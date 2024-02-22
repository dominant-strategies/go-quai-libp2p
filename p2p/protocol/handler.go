package protocol

import (
	"errors"
	"io"
	"math/big"

	"github.com/libp2p/go-libp2p/core/network"

	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/core/types"
	"github.com/dominant-strategies/go-quai/log"
	"github.com/dominant-strategies/go-quai/p2p/pb"
	"github.com/dominant-strategies/go-quai/trie"
)

// QuaiProtocolHandler handles all the incoming requests and responds with corresponding data
func QuaiProtocolHandler(stream network.Stream, node QuaiP2PNode) {
	defer stream.Close()

	log.Global.Debugf("Received a new stream from %s", stream.Conn().RemotePeer())

	// if there is a protocol mismatch, close the stream
	if stream.Protocol() != ProtocolVersion {
		log.Global.Warnf("Invalid protocol: %s", stream.Protocol())
		// TODO: add logic to drop the peer
		return
	}

	// go node.ReadResponses(stream)

	// Enter the read loop for the stream and handle messages
	for {
		data, err := common.ReadMessageFromStream(stream)
		if err != nil {
			if errors.Is(err, network.ErrReset) || errors.Is(err, io.EOF) {
				stream.Close()
				return
			}

			log.Global.Errorf("error reading message from stream: %s", err)
			// TODO: handle error
			continue
		}

		quaiMsg, err := pb.DecodeQuaiMessage(data)
		if err != nil {
			log.Global.Errorf("error decoding quai message: %s", err)
			continue
		}

		switch {
		case quaiMsg.GetRequest() != nil:
			handleRequest(quaiMsg.GetRequest(), stream, node)

		case quaiMsg.GetResponse() != nil:
			handleResponse(quaiMsg.GetResponse(), stream, node)

		default:
			log.Global.Errorf("unsupported quai message type: %T", quaiMsg)
		}
	}
	log.Global.Tracef("Exiting Quai Protocol Handler")
}

func handleRequest(quaiMsg *pb.QuaiRequestMessage, stream network.Stream, node QuaiP2PNode) {
	id, decodedType, loc, query, err := pb.DecodeQuaiRequest(quaiMsg)
	if err != nil {
		log.Global.Errorf("error decoding quai request: %s", err)
		// TODO: handle error

	}
	switch query.(type) {
	case *common.Hash:
		log.Global.Debugf("Received request id: %d for %T, location %v hash %s from peer %s", id, decodedType, loc, query, stream.Conn().RemotePeer())
	case *big.Int:
		log.Global.Debugf("Received request id: %d for %T, location %v number %s from peer %s", id, decodedType, loc, query, stream.Conn().RemotePeer())
	default:
		log.Global.Errorf("unsupported request input data field type: %T", query)
	}

	switch decodedType.(type) {
	case *types.Block:
		requestedHash := query.(*common.Hash)
		err = handleBlockRequest(id, loc, *requestedHash, stream, node)
		if err != nil {
			log.Global.Errorf("error handling block request: %s", err)
			// TODO: handle error
			return
		}
	case *types.Header:
		requestedHash := query.(*common.Hash)
		err = handleHeaderRequest(id, loc, *requestedHash, stream, node)
		if err != nil {
			log.Global.Errorf("error handling header request: %s", err)
			// TODO: handle error
			return
		}
	case *types.Transaction:
		requestedHash := query.(*common.Hash)
		err = handleTransactionRequest(id, loc, *requestedHash, stream, node)
		if err != nil {
			log.Global.Errorf("error handling transaction request: %s", err)
			// TODO: handle error
			return
		}
	case *common.Hash:
		number := query.(*big.Int)
		err = handleBlockNumberRequest(id, loc, number, stream, node)
		if err != nil {
			log.Global.Errorf("error handling block number request: %s", err)
			return
		}
	case trie.TrieNodeRequest:
		requestedHash := query.(*common.Hash)
		err := handleTrieNodeRequest(id, loc, *requestedHash, stream, node)
		if err != nil {
			log.Global.Errorf("error handling trie node request: %s", err)
		}
	default:
		log.Global.Errorf("unsupported request data type: %T", decodedType)
		// TODO: handle error
		return

	}
}

func handleResponse(quaiResp *pb.QuaiResponseMessage, stream network.Stream, node QuaiP2PNode) {
	recvdID, recvdType, err := pb.DecodeQuaiResponse(quaiResp)
	if err != nil {
		log.Global.WithField(
			"err", err,
		).Errorf("error decoding quai response: %s", err)
		return
	}

	dataChan, err := node.GetRequestManager().GetRequestChan(recvdID)
	if err != nil {
		log.Global.WithFields(log.Fields{
			"requestID": recvdID,
			"err":       err,
		}).Error("error associating request ID with data channel")
		return
	}
	dataChan <- recvdType
}

// Seeks the block in the cache or database and sends it to the peer in a pb.QuaiResponseMessage
func handleBlockRequest(id uint32, loc common.Location, hash common.Hash, stream network.Stream, node QuaiP2PNode) error {
	// check if we have the block in our cache or database
	block := node.GetBlock(hash, loc)
	if block == nil {
		log.Global.Debugf("block not found")
		return nil
	}
	log.Global.Debugf("block found %s", block.Hash())
	// create a Quai Message Response with the block
	data, err := pb.EncodeQuaiResponse(id, loc, block)
	if err != nil {
		return err
	}
	err = common.WriteMessageToStream(stream, data)
	if err != nil {
		return err
	}
	log.Global.Errorf("Sent block %s to peer %s", block.Hash(), stream.Conn().RemotePeer())
	return nil
}

// Seeks the header in the cache or database and sends it to the peer in a pb.QuaiResponseMessage
func handleHeaderRequest(id uint32, loc common.Location, hash common.Hash, stream network.Stream, node QuaiP2PNode) error {
	header := node.GetHeader(hash, loc)
	if header == nil {
		log.Global.Debugf("header not found")
		// TODO: handle header not found
		return nil
	}
	log.Global.Debugf("header found %s", header.Hash())
	// create a Quai Message Response with the header
	data, err := pb.EncodeQuaiResponse(id, loc, header)
	if err != nil {
		return err
	}
	err = common.WriteMessageToStream(stream, data)
	if err != nil {
		return err
	}
	log.Global.Debugf("Sent header %s to peer %s", header.Hash(), stream.Conn().RemotePeer())
	return nil
}

func handleTransactionRequest(id uint32, loc common.Location, hash common.Hash, stream network.Stream, node QuaiP2PNode) error {
	panic("TODO: implement")
}

// Seeks the block in the cache or database and sends it to the peer in a pb.QuaiResponseMessage
func handleBlockNumberRequest(id uint32, loc common.Location, number *big.Int, stream network.Stream, node QuaiP2PNode) error {
	// check if we have the block in our cache or database
	blockHash := node.GetBlockHashByNumber(number, loc)
	if blockHash == nil {
		log.Global.Tracef("block not found")
		return nil
	}
	log.Global.Tracef("block found %s", blockHash)
	// create a Quai Message Response with the block
	data, err := pb.EncodeQuaiResponse(id, loc, blockHash)
	if err != nil {
		return err
	}

	err = common.WriteMessageToStream(stream, data)
	if err != nil {
		return err
	}
	log.Global.Tracef("Sent block hash %s to peer %s", blockHash, stream.Conn().RemotePeer())
	return nil
}

func handleTrieNodeRequest(id uint32, loc common.Location, hash common.Hash, stream network.Stream, node QuaiP2PNode) error {
	trieNode := node.GetTrieNode(hash, loc)
	if trieNode == nil {
		log.Global.Tracef("trie node not found")
		return nil
	}
	log.Global.Tracef("trie node found")
	data, err := pb.EncodeQuaiResponse(id, loc, trieNode)
	if err != nil {
		return err
	}
	err = common.WriteMessageToStream(stream, data)
	if err != nil {
		return err
	}
	log.Global.Tracef("Sent trie node to peer %s", stream.Conn().RemotePeer())
	return nil
}

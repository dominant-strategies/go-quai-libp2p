package pb

import (
	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/core/types"
	"github.com/dominant-strategies/go-quai/log"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// Unmarshals a serialized protobuf slice of bytes into a protocol buffer type
func UnmarshalProtoMessage(data []byte, msg proto.Message) error {
	if err := proto.Unmarshal(data, msg); err != nil {
		log.Errorf("error unmarshaling proto message: %s", err)
		return err
	}
	return nil
}

// Marshals a protocol buffer type into a serialized protobuf slice of bytes
func MarshalProtoMessage(msg proto.Message) ([]byte, error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Errorf("error marshaling proto message: %s", err)
		return nil, err
	}
	log.Tracef("marshaled proto message - data length: %d, data=%v", len(data), data)
	return data, nil
}

// EncodeRequestMessage creates a marshaled protobuf message for a Quai Request.
// Returns the serialized protobuf message.
func EncodeQuaiRequest(action QuaiRequestMessage_ActionType, location *common.Location, hash *common.Hash) ([]byte, error) {
	log.Tracef("encoding quai request: action=%v, location=%s, hash=%v", action, location, hash)

	protoHash, err := convertHashToProto(hash)
	if err != nil {
		log.Errorf("error converting hash to proto: %s", err)
		return nil, err
	}

	protoLocation, err := convertLocationToProto(location)
	if err != nil {
		log.Errorf("error converting slice to proto: %s", err)
		return nil, err
	}

	request := &Request{
		Hash:     protoHash,
		Location: protoLocation,
	}

	quaiMsg := &QuaiRequestMessage{
		Action:  action,
		Request: request,
	}

	log.Tracef("encoded quai request - quaiMsg=%+v", quaiMsg)

	return MarshalProtoMessage(quaiMsg)

}

// DecodeRequestMessage unmarshals a protobuf message into a Quai Request.
// Returns the action type, sliceID, and hash.
func DecodeQuaiRequest(data []byte) (action QuaiRequestMessage_ActionType, location *common.Location, hash *common.Hash, err error) {
	var quaiMsg QuaiRequestMessage
	log.Tracef("decoding quai request - data length: %d,  data=%v", len(data), data)
	err = UnmarshalProtoMessage(data, &quaiMsg)
	if err != nil {
		return QuaiRequestMessage_UNKNOWN, nil, nil, err
	}

	action = quaiMsg.Action

	switch action {
	case QuaiRequestMessage_REQUEST_BLOCK, QuaiRequestMessage_REQUEST_HEADER, QuaiRequestMessage_REQUEST_TRANSACTION:
		request := quaiMsg.GetRequest()
		protoHash := request.GetHash()
		protoLocation := request.GetLocation()

		location, err = convertProtoToLocation(protoLocation)
		if err != nil {
			return QuaiRequestMessage_UNKNOWN, nil, nil, err
		}

		hash, err = convertProtoToHash(protoHash)
		if err != nil {
			return QuaiRequestMessage_UNKNOWN, nil, nil, err
		}

	default:
		return QuaiRequestMessage_UNKNOWN, nil, nil, errors.Errorf("unsupported action type: %v", action)
	}

	return action, location, hash, nil
}

// EncodeResponse creates a marshaled protobuf message for a Quai Response.
// Returns the serialized protobuf message.
func EncodeQuaiResponse(action QuaiResponseMessage_ActionType, data interface{}) ([]byte, error) {

	var quaiMsg *QuaiResponseMessage
	response, err := convertDataToProtoResponse(action, data)
	if err != nil {
		return nil, err
	}
	quaiMsg = &QuaiResponseMessage{
		Action:   action,
		Response: response,
	}

	return MarshalProtoMessage(quaiMsg)
}

// Unmarshals a serialized protobuf message into a Quai Response message.
// Returns the action type and the decoded type (i.e. *types.Header, *types.Block, etc).
func DecodeQuaiResponse(data []byte) (action QuaiResponseMessage_ActionType, response interface{}, err error) {
	var quaiMsg QuaiResponseMessage
	err = UnmarshalProtoMessage(data, &quaiMsg)
	if err != nil {
		return QuaiResponseMessage_UNKNOWN, nil, err
	}

	action = quaiMsg.Action

	switch action {
	case QuaiResponseMessage_RESPONSE_BLOCK:
		protoBlock := quaiMsg.Response.GetBlock()
		block, err := convertProtoToBlock(protoBlock)
		if err != nil {
			return QuaiResponseMessage_UNKNOWN, nil, err
		}
		response = block

	case QuaiResponseMessage_RESPONSE_HEADER:
		protoHeader := quaiMsg.Response.GetHeader()
		header, err := convertProtoToHeader(protoHeader)
		if err != nil {
			return QuaiResponseMessage_UNKNOWN, nil, err
		}
		response = header

	case QuaiResponseMessage_RESPONSE_TRANSACTION:
		protoTransaction := quaiMsg.Response.GetTransaction()
		transaction, err := convertProtoToTransaction(protoTransaction)
		if err != nil {
			return QuaiResponseMessage_UNKNOWN, nil, err
		}
		response = transaction
	default:
		return QuaiResponseMessage_UNKNOWN, nil, errors.Errorf("unsupported action type: %v", action)
	}

	return action, response, nil
}

// Converts a custom go type to a proto type and marhsals it into a protobuf message
func ConvertAndMarshal(data interface{}) ([]byte, error) {
	switch data := data.(type) {
	case *types.Block:
		log.Tracef("marshalling block: %+v", data)
		protoBlock, err := convertBlockToProto(data)
		if err != nil {
			return nil, err
		}
		return MarshalProtoMessage(protoBlock)
	case *types.Transaction:
		log.Tracef("marshalling transaction: %+v", data)
		protoTransaction, err := convertTransactionToProto(data)
		if err != nil {
			return nil, err
		}
		return MarshalProtoMessage(protoTransaction)
	case *types.Header:
		log.Tracef("marshalling header: %+v", data)
		protoHeader, err := convertHeaderToProto(data)
		if err != nil {
			return nil, err
		}
		return MarshalProtoMessage(protoHeader)
	default:
		log.Debugf("unsupported data type: %T", data)
		return nil, errors.New("unsupported data type")
	}
}

// Unmarshals a protobuf message into a proto type and converts it to a custom go type
func UnmarshalAndConvert(data []byte, dataPtr interface{}) error {
	switch dataPtr := dataPtr.(type) {
	case *types.Block:
		protoBlock := new(Block)
		err := UnmarshalProtoMessage(data, protoBlock)
		if err != nil {
			return err
		}
		block, err := convertProtoToBlock(protoBlock)
		if err != nil {
			return err
		}
		*dataPtr = *block
		return nil
	case *types.Transaction:
		protoTransaction := new(Transaction)
		err := UnmarshalProtoMessage(data, protoTransaction)
		if err != nil {
			return err
		}
		transaction, err := convertProtoToTransaction(protoTransaction)
		if err != nil {
			return err
		}
		*dataPtr = *transaction
		return nil
	case *types.Header:
		protoHeader := new(Header)
		err := UnmarshalProtoMessage(data, protoHeader)
		if err != nil {
			return err
		}
		header, err := convertProtoToHeader(protoHeader)
		if err != nil {
			return err
		}
		*dataPtr = *header
		return nil
	default:
		log.Debugf("unsupported data type: %T", dataPtr)
		return errors.New("unsupported data type")
	}
}

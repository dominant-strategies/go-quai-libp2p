package pb

import (
	"github.com/dominant-strategies/go-quai/log"
	"google.golang.org/protobuf/proto"
)

// Unmarshals a serialized protobuf slice of bytes into a protocol buffer type
func UnmarshalProtoMessage(data []byte, msg proto.Message) error {
	if err := proto.Unmarshal(data, msg); err != nil {
		return err
	}
	return nil
}

// Marshals a protocol buffer type into a serialized protobuf slice of bytes
func MarshalProtoMessage(msg proto.Message) ([]byte, error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ProtoConvertable is an interface for types that shall be converted to a protobuf message.
type ProtoConvertable[T proto.Message] interface {
	ToProto() T
}

// ConvertToProto converts a ProtoConvertable type to a protobuf message.
func ConvertToProto[P proto.Message](c ProtoConvertable[P]) P {
	return c.ToProto()
}

// ConvertableFromProto is an interface for types that shall be converted from a protobuf message.
type ConvertableFromProto[P proto.Message] interface {
	FromProto(protoMsg P)
	NewProtoInstance() P
}

// UnmarshalAndConvert takes a slice of bytes (protobuf serialized data) and
// unmarshals it into a protobuf message, then converts it to a custom type
// using the FromProto method. The appropriate protobuf type is determined by the NewProtoInstance interface ConvertableFromProto method.
func UnmarshalAndConvert[T ConvertableFromProto[P], P proto.Message](data []byte, target T) error {
	log.Tracef("Unmarshalling protobuf message: %+v", data)
	protoMsg := target.NewProtoInstance() // Create a new instance of the protobuf message type
	if err := proto.Unmarshal(data, protoMsg); err != nil {
		return err
	}
	log.Tracef("Unmarshalled protobuf message: %+v", protoMsg)
	target.FromProto(protoMsg)
	log.Tracef("Converted protobuf message to custom type: %+v", target)
	return nil
}

// ConvertAndMarshal takes a custom type and converts it to a protobuf message using the ToProto method,
// then marshals it into a slice of bytes (protobuf serialized data).
func ConvertAndMarshal[T ProtoConvertable[P], P proto.Message](target T) ([]byte, error) {
	protoMsg := target.ToProto()
	log.Tracef("Converted custom type to protobuf message: %+v", protoMsg)
	data, err := proto.Marshal(protoMsg)
	if err != nil {
		return nil, err
	}
	log.Tracef("Marshalled protobuf message: %+v", data)
	return data, nil
}

// Creates a marshaled protobuf message for a block request.
func CreateProtoBlockRequest(hash ProtoConvertable[*Hash], slice ProtoConvertable[*SliceID]) ([]byte, error) {
	blockReq := &BlockRequest{
		Hash:    hash.ToProto(),
		SliceId: slice.ToProto(),
	}
	return MarshalProtoMessage(blockReq)
}

// Unmarhsalls a protobuf block response message, and returns a boolean indicating whether the block was found.
// If the block was found, it returns a *pb.Block, otherwise it returns nil.
func UnmarshalProtoBlockResponse(data []byte) (bool, *Block, error) {
	var blockResponse BlockResponse
	err := UnmarshalProtoMessage(data, &blockResponse)
	if err != nil {
		return false, nil, err
	}
	if blockResponse.Found {
		return true, blockResponse.Block, nil
	}
	return false, nil, nil
}

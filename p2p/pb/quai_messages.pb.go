// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: p2p/pb/quai_messages.proto

package pb

import (
	common "github.com/dominant-strategies/go-quai/common"
	types "github.com/dominant-strategies/go-quai/core/types"
	trie "github.com/dominant-strategies/go-quai/trie"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GossipSub messages for broadcasting blocks and transactions
type GossipBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *types.ProtoBlock `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *GossipBlock) Reset() {
	*x = GossipBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipBlock) ProtoMessage() {}

func (x *GossipBlock) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipBlock.ProtoReflect.Descriptor instead.
func (*GossipBlock) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GossipBlock) GetBlock() *types.ProtoBlock {
	if x != nil {
		return x.Block
	}
	return nil
}

type GossipTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *types.ProtoTransaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *GossipTransaction) Reset() {
	*x = GossipTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipTransaction) ProtoMessage() {}

func (x *GossipTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipTransaction.ProtoReflect.Descriptor instead.
func (*GossipTransaction) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GossipTransaction) GetTransaction() *types.ProtoTransaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

// QuaiRequestMessage is the main 'envelope' for QuaiProtocol request messages
type QuaiRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Location *common.ProtoLocation `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// Types that are assignable to Data:
	//
	//	*QuaiRequestMessage_Hash
	//	*QuaiRequestMessage_Number
	Data isQuaiRequestMessage_Data `protobuf_oneof:"data"`
	// Types that are assignable to Request:
	//
	//	*QuaiRequestMessage_Block
	//	*QuaiRequestMessage_Header
	//	*QuaiRequestMessage_Transaction
	//	*QuaiRequestMessage_BlockHash
	//	*QuaiRequestMessage_TrieNode
	Request isQuaiRequestMessage_Request `protobuf_oneof:"request"`
}

func (x *QuaiRequestMessage) Reset() {
	*x = QuaiRequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuaiRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuaiRequestMessage) ProtoMessage() {}

func (x *QuaiRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuaiRequestMessage.ProtoReflect.Descriptor instead.
func (*QuaiRequestMessage) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{2}
}

func (x *QuaiRequestMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QuaiRequestMessage) GetLocation() *common.ProtoLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (m *QuaiRequestMessage) GetData() isQuaiRequestMessage_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *QuaiRequestMessage) GetHash() *common.ProtoHash {
	if x, ok := x.GetData().(*QuaiRequestMessage_Hash); ok {
		return x.Hash
	}
	return nil
}

func (x *QuaiRequestMessage) GetNumber() []byte {
	if x, ok := x.GetData().(*QuaiRequestMessage_Number); ok {
		return x.Number
	}
	return nil
}

func (m *QuaiRequestMessage) GetRequest() isQuaiRequestMessage_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *QuaiRequestMessage) GetBlock() *types.ProtoBlock {
	if x, ok := x.GetRequest().(*QuaiRequestMessage_Block); ok {
		return x.Block
	}
	return nil
}

func (x *QuaiRequestMessage) GetHeader() *types.ProtoHeader {
	if x, ok := x.GetRequest().(*QuaiRequestMessage_Header); ok {
		return x.Header
	}
	return nil
}

func (x *QuaiRequestMessage) GetTransaction() *types.ProtoTransaction {
	if x, ok := x.GetRequest().(*QuaiRequestMessage_Transaction); ok {
		return x.Transaction
	}
	return nil
}

func (x *QuaiRequestMessage) GetBlockHash() *common.ProtoHash {
	if x, ok := x.GetRequest().(*QuaiRequestMessage_BlockHash); ok {
		return x.BlockHash
	}
	return nil
}

func (x *QuaiRequestMessage) GetTrieNode() *trie.ProtoTrieNode {
	if x, ok := x.GetRequest().(*QuaiRequestMessage_TrieNode); ok {
		return x.TrieNode
	}
	return nil
}

type isQuaiRequestMessage_Data interface {
	isQuaiRequestMessage_Data()
}

type QuaiRequestMessage_Hash struct {
	Hash *common.ProtoHash `protobuf:"bytes,3,opt,name=hash,proto3,oneof"`
}

type QuaiRequestMessage_Number struct {
	Number []byte `protobuf:"bytes,7,opt,name=number,proto3,oneof"`
}

func (*QuaiRequestMessage_Hash) isQuaiRequestMessage_Data() {}

func (*QuaiRequestMessage_Number) isQuaiRequestMessage_Data() {}

type isQuaiRequestMessage_Request interface {
	isQuaiRequestMessage_Request()
}

type QuaiRequestMessage_Block struct {
	Block *types.ProtoBlock `protobuf:"bytes,4,opt,name=block,proto3,oneof"`
}

type QuaiRequestMessage_Header struct {
	Header *types.ProtoHeader `protobuf:"bytes,5,opt,name=header,proto3,oneof"`
}

type QuaiRequestMessage_Transaction struct {
	Transaction *types.ProtoTransaction `protobuf:"bytes,6,opt,name=transaction,proto3,oneof"`
}

type QuaiRequestMessage_BlockHash struct {
	BlockHash *common.ProtoHash `protobuf:"bytes,8,opt,name=blockHash,proto3,oneof"`
}

type QuaiRequestMessage_TrieNode struct {
	TrieNode *trie.ProtoTrieNode `protobuf:"bytes,9,opt,name=trieNode,proto3,oneof"`
}

func (*QuaiRequestMessage_Block) isQuaiRequestMessage_Request() {}

func (*QuaiRequestMessage_Header) isQuaiRequestMessage_Request() {}

func (*QuaiRequestMessage_Transaction) isQuaiRequestMessage_Request() {}

func (*QuaiRequestMessage_BlockHash) isQuaiRequestMessage_Request() {}

func (*QuaiRequestMessage_TrieNode) isQuaiRequestMessage_Request() {}

// QuaiResponseMessage is the main 'envelope' for QuaiProtocol response messages
type QuaiResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Location *common.ProtoLocation `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// Types that are assignable to Response:
	//
	//	*QuaiResponseMessage_Block
	//	*QuaiResponseMessage_Header
	//	*QuaiResponseMessage_Transaction
	//	*QuaiResponseMessage_BlockHash
	//	*QuaiResponseMessage_TrieNode
	Response isQuaiResponseMessage_Response `protobuf_oneof:"response"`
}

func (x *QuaiResponseMessage) Reset() {
	*x = QuaiResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuaiResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuaiResponseMessage) ProtoMessage() {}

func (x *QuaiResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuaiResponseMessage.ProtoReflect.Descriptor instead.
func (*QuaiResponseMessage) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{3}
}

func (x *QuaiResponseMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QuaiResponseMessage) GetLocation() *common.ProtoLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (m *QuaiResponseMessage) GetResponse() isQuaiResponseMessage_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *QuaiResponseMessage) GetBlock() *types.ProtoBlock {
	if x, ok := x.GetResponse().(*QuaiResponseMessage_Block); ok {
		return x.Block
	}
	return nil
}

func (x *QuaiResponseMessage) GetHeader() *types.ProtoHeader {
	if x, ok := x.GetResponse().(*QuaiResponseMessage_Header); ok {
		return x.Header
	}
	return nil
}

func (x *QuaiResponseMessage) GetTransaction() *types.ProtoTransaction {
	if x, ok := x.GetResponse().(*QuaiResponseMessage_Transaction); ok {
		return x.Transaction
	}
	return nil
}

func (x *QuaiResponseMessage) GetBlockHash() *common.ProtoHash {
	if x, ok := x.GetResponse().(*QuaiResponseMessage_BlockHash); ok {
		return x.BlockHash
	}
	return nil
}

func (x *QuaiResponseMessage) GetTrieNode() *trie.ProtoTrieNode {
	if x, ok := x.GetResponse().(*QuaiResponseMessage_TrieNode); ok {
		return x.TrieNode
	}
	return nil
}

type isQuaiResponseMessage_Response interface {
	isQuaiResponseMessage_Response()
}

type QuaiResponseMessage_Block struct {
	Block *types.ProtoBlock `protobuf:"bytes,3,opt,name=block,proto3,oneof"`
}

type QuaiResponseMessage_Header struct {
	Header *types.ProtoHeader `protobuf:"bytes,4,opt,name=header,proto3,oneof"`
}

type QuaiResponseMessage_Transaction struct {
	Transaction *types.ProtoTransaction `protobuf:"bytes,5,opt,name=transaction,proto3,oneof"`
}

type QuaiResponseMessage_BlockHash struct {
	BlockHash *common.ProtoHash `protobuf:"bytes,6,opt,name=blockHash,proto3,oneof"`
}

type QuaiResponseMessage_TrieNode struct {
	TrieNode *trie.ProtoTrieNode `protobuf:"bytes,7,opt,name=trieNode,proto3,oneof"`
}

func (*QuaiResponseMessage_Block) isQuaiResponseMessage_Response() {}

func (*QuaiResponseMessage_Header) isQuaiResponseMessage_Response() {}

func (*QuaiResponseMessage_Transaction) isQuaiResponseMessage_Response() {}

func (*QuaiResponseMessage_BlockHash) isQuaiResponseMessage_Response() {}

func (*QuaiResponseMessage_TrieNode) isQuaiResponseMessage_Response() {}

type QuaiMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*QuaiMessage_Request
	//	*QuaiMessage_Response
	Payload isQuaiMessage_Payload `protobuf_oneof:"payload"`
}

func (x *QuaiMessage) Reset() {
	*x = QuaiMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuaiMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuaiMessage) ProtoMessage() {}

func (x *QuaiMessage) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuaiMessage.ProtoReflect.Descriptor instead.
func (*QuaiMessage) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{4}
}

func (m *QuaiMessage) GetPayload() isQuaiMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *QuaiMessage) GetRequest() *QuaiRequestMessage {
	if x, ok := x.GetPayload().(*QuaiMessage_Request); ok {
		return x.Request
	}
	return nil
}

func (x *QuaiMessage) GetResponse() *QuaiResponseMessage {
	if x, ok := x.GetPayload().(*QuaiMessage_Response); ok {
		return x.Response
	}
	return nil
}

type isQuaiMessage_Payload interface {
	isQuaiMessage_Payload()
}

type QuaiMessage_Request struct {
	Request *QuaiRequestMessage `protobuf:"bytes,1,opt,name=request,proto3,oneof"`
}

type QuaiMessage_Response struct {
	Response *QuaiResponseMessage `protobuf:"bytes,2,opt,name=response,proto3,oneof"`
}

func (*QuaiMessage_Request) isQuaiMessage_Payload() {}

func (*QuaiMessage_Response) isQuaiMessage_Payload() {}

var File_p2p_pb_quai_messages_proto protoreflect.FileDescriptor

var file_p2p_pb_quai_messages_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x75, 0x61, 0x69, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x71, 0x75,
	0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x74, 0x72, 0x69, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x74, 0x72, 0x69, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36,
	0x0a, 0x0b, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x27, 0x0a,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x4e, 0x0a, 0x11, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa9, 0x03, 0x0a, 0x12, 0x51, 0x75, 0x61, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x27, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x61, 0x73,
	0x68, 0x48, 0x00, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x01, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2c,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x48, 0x01, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x48,
	0x01, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x31, 0x0a, 0x08,
	0x74, 0x72, 0x69, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x74, 0x72, 0x69, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x72, 0x69, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x48, 0x01, 0x52, 0x08, 0x74, 0x72, 0x69, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x42,
	0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xe0, 0x02, 0x0a, 0x13, 0x51, 0x75, 0x61, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x00, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x61, 0x73, 0x68, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x31, 0x0a, 0x08, 0x74, 0x72, 0x69, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x72, 0x69, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x72, 0x69, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52,
	0x08, 0x74, 0x72, 0x69, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x51, 0x75, 0x61, 0x69, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x51, 0x75, 0x61, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x51, 0x75, 0x61, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42,
	0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f,
	0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65,
	0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x71, 0x75, 0x61, 0x69, 0x2f, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_p2p_pb_quai_messages_proto_rawDescOnce sync.Once
	file_p2p_pb_quai_messages_proto_rawDescData = file_p2p_pb_quai_messages_proto_rawDesc
)

func file_p2p_pb_quai_messages_proto_rawDescGZIP() []byte {
	file_p2p_pb_quai_messages_proto_rawDescOnce.Do(func() {
		file_p2p_pb_quai_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_pb_quai_messages_proto_rawDescData)
	})
	return file_p2p_pb_quai_messages_proto_rawDescData
}

var file_p2p_pb_quai_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_p2p_pb_quai_messages_proto_goTypes = []interface{}{
	(*GossipBlock)(nil),            // 0: quaiprotocol.GossipBlock
	(*GossipTransaction)(nil),      // 1: quaiprotocol.GossipTransaction
	(*QuaiRequestMessage)(nil),     // 2: quaiprotocol.QuaiRequestMessage
	(*QuaiResponseMessage)(nil),    // 3: quaiprotocol.QuaiResponseMessage
	(*QuaiMessage)(nil),            // 4: quaiprotocol.QuaiMessage
	(*types.ProtoBlock)(nil),       // 5: block.ProtoBlock
	(*types.ProtoTransaction)(nil), // 6: block.ProtoTransaction
	(*common.ProtoLocation)(nil),   // 7: common.ProtoLocation
	(*common.ProtoHash)(nil),       // 8: common.ProtoHash
	(*types.ProtoHeader)(nil),      // 9: block.ProtoHeader
	(*trie.ProtoTrieNode)(nil),     // 10: trie.ProtoTrieNode
}
var file_p2p_pb_quai_messages_proto_depIdxs = []int32{
	5,  // 0: quaiprotocol.GossipBlock.block:type_name -> block.ProtoBlock
	6,  // 1: quaiprotocol.GossipTransaction.transaction:type_name -> block.ProtoTransaction
	7,  // 2: quaiprotocol.QuaiRequestMessage.location:type_name -> common.ProtoLocation
	8,  // 3: quaiprotocol.QuaiRequestMessage.hash:type_name -> common.ProtoHash
	5,  // 4: quaiprotocol.QuaiRequestMessage.block:type_name -> block.ProtoBlock
	9,  // 5: quaiprotocol.QuaiRequestMessage.header:type_name -> block.ProtoHeader
	6,  // 6: quaiprotocol.QuaiRequestMessage.transaction:type_name -> block.ProtoTransaction
	8,  // 7: quaiprotocol.QuaiRequestMessage.blockHash:type_name -> common.ProtoHash
	10, // 8: quaiprotocol.QuaiRequestMessage.trieNode:type_name -> trie.ProtoTrieNode
	7,  // 9: quaiprotocol.QuaiResponseMessage.location:type_name -> common.ProtoLocation
	5,  // 10: quaiprotocol.QuaiResponseMessage.block:type_name -> block.ProtoBlock
	9,  // 11: quaiprotocol.QuaiResponseMessage.header:type_name -> block.ProtoHeader
	6,  // 12: quaiprotocol.QuaiResponseMessage.transaction:type_name -> block.ProtoTransaction
	8,  // 13: quaiprotocol.QuaiResponseMessage.blockHash:type_name -> common.ProtoHash
	10, // 14: quaiprotocol.QuaiResponseMessage.trieNode:type_name -> trie.ProtoTrieNode
	2,  // 15: quaiprotocol.QuaiMessage.request:type_name -> quaiprotocol.QuaiRequestMessage
	3,  // 16: quaiprotocol.QuaiMessage.response:type_name -> quaiprotocol.QuaiResponseMessage
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_p2p_pb_quai_messages_proto_init() }
func file_p2p_pb_quai_messages_proto_init() {
	if File_p2p_pb_quai_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p2p_pb_quai_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GossipBlock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_pb_quai_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GossipTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_pb_quai_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuaiRequestMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_pb_quai_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuaiResponseMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_pb_quai_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuaiMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_p2p_pb_quai_messages_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*QuaiRequestMessage_Hash)(nil),
		(*QuaiRequestMessage_Number)(nil),
		(*QuaiRequestMessage_Block)(nil),
		(*QuaiRequestMessage_Header)(nil),
		(*QuaiRequestMessage_Transaction)(nil),
		(*QuaiRequestMessage_BlockHash)(nil),
		(*QuaiRequestMessage_TrieNode)(nil),
	}
	file_p2p_pb_quai_messages_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*QuaiResponseMessage_Block)(nil),
		(*QuaiResponseMessage_Header)(nil),
		(*QuaiResponseMessage_Transaction)(nil),
		(*QuaiResponseMessage_BlockHash)(nil),
		(*QuaiResponseMessage_TrieNode)(nil),
	}
	file_p2p_pb_quai_messages_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*QuaiMessage_Request)(nil),
		(*QuaiMessage_Response)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_p2p_pb_quai_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_p2p_pb_quai_messages_proto_goTypes,
		DependencyIndexes: file_p2p_pb_quai_messages_proto_depIdxs,
		MessageInfos:      file_p2p_pb_quai_messages_proto_msgTypes,
	}.Build()
	File_p2p_pb_quai_messages_proto = out.File
	file_p2p_pb_quai_messages_proto_rawDesc = nil
	file_p2p_pb_quai_messages_proto_goTypes = nil
	file_p2p_pb_quai_messages_proto_depIdxs = nil
}

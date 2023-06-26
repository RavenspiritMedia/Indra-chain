// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: operator.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventType int32

const (
	// For initially added transactions
	EventType_ADDED EventType = 0
	// For enqueued transactions in the account queue
	EventType_ENQUEUED EventType = 1
	// For promoted transactions
	EventType_PROMOTED EventType = 2
	// For dropped transactions
	EventType_DROPPED EventType = 3
	// For demoted transactions
	EventType_DEMOTED EventType = 4
	// For pruned promoted transactions
	EventType_PRUNED_PROMOTED EventType = 5
	// For pruned enqueued transactions
	EventType_PRUNED_ENQUEUED EventType = 6
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "ADDED",
		1: "ENQUEUED",
		2: "PROMOTED",
		3: "DROPPED",
		4: "DEMOTED",
		5: "PRUNED_PROMOTED",
		6: "PRUNED_ENQUEUED",
	}
	EventType_value = map[string]int32{
		"ADDED":           0,
		"ENQUEUED":        1,
		"PROMOTED":        2,
		"DROPPED":         3,
		"DEMOTED":         4,
		"PRUNED_PROMOTED": 5,
		"PRUNED_ENQUEUED": 6,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_operator_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_operator_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_operator_proto_rawDescGZIP(), []int{0}
}

type AddTxnReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw  *anypb.Any `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	From string     `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *AddTxnReq) Reset() {
	*x = AddTxnReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTxnReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTxnReq) ProtoMessage() {}

func (x *AddTxnReq) ProtoReflect() protoreflect.Message {
	mi := &file_operator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTxnReq.ProtoReflect.Descriptor instead.
func (*AddTxnReq) Descriptor() ([]byte, []int) {
	return file_operator_proto_rawDescGZIP(), []int{0}
}

func (x *AddTxnReq) GetRaw() *anypb.Any {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *AddTxnReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

type AddTxnResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash string `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
}

func (x *AddTxnResp) Reset() {
	*x = AddTxnResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTxnResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTxnResp) ProtoMessage() {}

func (x *AddTxnResp) ProtoReflect() protoreflect.Message {
	mi := &file_operator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTxnResp.ProtoReflect.Descriptor instead.
func (*AddTxnResp) Descriptor() ([]byte, []int) {
	return file_operator_proto_rawDescGZIP(), []int{1}
}

func (x *AddTxnResp) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type TxnPoolStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length uint64 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *TxnPoolStatusResp) Reset() {
	*x = TxnPoolStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnPoolStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnPoolStatusResp) ProtoMessage() {}

func (x *TxnPoolStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_operator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnPoolStatusResp.ProtoReflect.Descriptor instead.
func (*TxnPoolStatusResp) Descriptor() ([]byte, []int) {
	return file_operator_proto_rawDescGZIP(), []int{2}
}

func (x *TxnPoolStatusResp) GetLength() uint64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested event types
	Types []EventType `protobuf:"varint,1,rep,packed,name=types,proto3,enum=v1.EventType" json:"types,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_operator_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeRequest) GetTypes() []EventType {
	if x != nil {
		return x.Types
	}
	return nil
}

type TxPoolEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   EventType `protobuf:"varint,1,opt,name=type,proto3,enum=v1.EventType" json:"type,omitempty"`
	TxHash string    `protobuf:"bytes,2,opt,name=txHash,proto3" json:"txHash,omitempty"`
}

func (x *TxPoolEvent) Reset() {
	*x = TxPoolEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxPoolEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxPoolEvent) ProtoMessage() {}

func (x *TxPoolEvent) ProtoReflect() protoreflect.Message {
	mi := &file_operator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxPoolEvent.ProtoReflect.Descriptor instead.
func (*TxPoolEvent) Descriptor() ([]byte, []int) {
	return file_operator_proto_rawDescGZIP(), []int{4}
}

func (x *TxPoolEvent) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_ADDED
}

func (x *TxPoolEvent) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

var File_operator_proto protoreflect.FileDescriptor

var file_operator_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x09,
	0x41, 0x64, 0x64, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x03, 0x72, 0x61, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x03, 0x72, 0x61,
	0x77, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x24, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x54, 0x78, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22, 0x2b, 0x0a, 0x11, 0x54,
	0x78, 0x6e, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x37, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x22, 0x48, 0x0a, 0x0b, 0x54, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x2a, 0x76, 0x0a, 0x09, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x44, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x44, 0x52, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x55,
	0x4e, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x13,
	0x0a, 0x0f, 0x50, 0x52, 0x55, 0x4e, 0x45, 0x44, 0x5f, 0x45, 0x4e, 0x51, 0x55, 0x45, 0x55, 0x45,
	0x44, 0x10, 0x06, 0x32, 0xa9, 0x01, 0x0a, 0x0f, 0x54, 0x78, 0x6e, 0x50, 0x6f, 0x6f, 0x6c, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x78, 0x6e, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x27, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x54, 0x78, 0x6e, 0x12, 0x0d, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x09, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x42,
	0x0f, 0x5a, 0x0d, 0x2f, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operator_proto_rawDescOnce sync.Once
	file_operator_proto_rawDescData = file_operator_proto_rawDesc
)

func file_operator_proto_rawDescGZIP() []byte {
	file_operator_proto_rawDescOnce.Do(func() {
		file_operator_proto_rawDescData = protoimpl.X.CompressGZIP(file_operator_proto_rawDescData)
	})
	return file_operator_proto_rawDescData
}

var file_operator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_operator_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_operator_proto_goTypes = []interface{}{
	(EventType)(0),            // 0: v1.EventType
	(*AddTxnReq)(nil),         // 1: v1.AddTxnReq
	(*AddTxnResp)(nil),        // 2: v1.AddTxnResp
	(*TxnPoolStatusResp)(nil), // 3: v1.TxnPoolStatusResp
	(*SubscribeRequest)(nil),  // 4: v1.SubscribeRequest
	(*TxPoolEvent)(nil),       // 5: v1.TxPoolEvent
	(*anypb.Any)(nil),         // 6: google.protobuf.Any
	(*emptypb.Empty)(nil),     // 7: google.protobuf.Empty
}
var file_operator_proto_depIdxs = []int32{
	6, // 0: v1.AddTxnReq.raw:type_name -> google.protobuf.Any
	0, // 1: v1.SubscribeRequest.types:type_name -> v1.EventType
	0, // 2: v1.TxPoolEvent.type:type_name -> v1.EventType
	7, // 3: v1.TxnPoolOperator.Status:input_type -> google.protobuf.Empty
	1, // 4: v1.TxnPoolOperator.AddTxn:input_type -> v1.AddTxnReq
	4, // 5: v1.TxnPoolOperator.Subscribe:input_type -> v1.SubscribeRequest
	3, // 6: v1.TxnPoolOperator.Status:output_type -> v1.TxnPoolStatusResp
	2, // 7: v1.TxnPoolOperator.AddTxn:output_type -> v1.AddTxnResp
	5, // 8: v1.TxnPoolOperator.Subscribe:output_type -> v1.TxPoolEvent
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_operator_proto_init() }
func file_operator_proto_init() {
	if File_operator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_operator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTxnReq); i {
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
		file_operator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTxnResp); i {
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
		file_operator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnPoolStatusResp); i {
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
		file_operator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_operator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxPoolEvent); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_operator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operator_proto_goTypes,
		DependencyIndexes: file_operator_proto_depIdxs,
		EnumInfos:         file_operator_proto_enumTypes,
		MessageInfos:      file_operator_proto_msgTypes,
	}.Build()
	File_operator_proto = out.File
	file_operator_proto_rawDesc = nil
	file_operator_proto_goTypes = nil
	file_operator_proto_depIdxs = nil
}

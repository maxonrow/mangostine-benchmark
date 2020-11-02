// version of protocol buffer used

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: mangostine.proto

// package name for the buffer will be used later

package mangostine_v0

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// argument
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// data type and position of data
	Sender    string  `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver  []*Send `protobuf:"bytes,2,rep,name=receiver,proto3" json:"receiver,omitempty"`
	PublicKey []byte  `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Signature []byte  `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mangostine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_mangostine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_mangostine_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Transaction) GetReceiver() []*Send {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *Transaction) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Transaction) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Send struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To     string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Send) Reset() {
	*x = Send{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mangostine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Send) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Send) ProtoMessage() {}

func (x *Send) ProtoReflect() protoreflect.Message {
	mi := &file_mangostine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Send.ProtoReflect.Descriptor instead.
func (*Send) Descriptor() ([]byte, []int) {
	return file_mangostine_proto_rawDescGZIP(), []int{1}
}

func (x *Send) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Send) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type RawTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawTx []byte `protobuf:"bytes,1,opt,name=rawTx,proto3" json:"rawTx,omitempty"`
}

func (x *RawTransaction) Reset() {
	*x = RawTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mangostine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawTransaction) ProtoMessage() {}

func (x *RawTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_mangostine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawTransaction.ProtoReflect.Descriptor instead.
func (*RawTransaction) Descriptor() ([]byte, []int) {
	return file_mangostine_proto_rawDescGZIP(), []int{2}
}

func (x *RawTransaction) GetRawTx() []byte {
	if x != nil {
		return x.RawTx
	}
	return nil
}

// return value
type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// data type and position of data
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mangostine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mangostine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_mangostine_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_mangostine_proto protoreflect.FileDescriptor

var file_mangostine_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x76,
	0x30, 0x22, 0x93, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61,
	0x6e, 0x67, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x2e, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x0e, 0x52, 0x61, 0x77, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x77,
	0x54, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x61, 0x77, 0x54, 0x78, 0x22,
	0x2f, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x9f, 0x01, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x48, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x78, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x65, 0x5f,
	0x76, 0x30, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x22,
	0x2e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x76, 0x30, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x61, 0x77, 0x54, 0x78, 0x12,
	0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x76, 0x30, 0x2e,
	0x52, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x22,
	0x2e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x76, 0x30, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mangostine_proto_rawDescOnce sync.Once
	file_mangostine_proto_rawDescData = file_mangostine_proto_rawDesc
)

func file_mangostine_proto_rawDescGZIP() []byte {
	file_mangostine_proto_rawDescOnce.Do(func() {
		file_mangostine_proto_rawDescData = protoimpl.X.CompressGZIP(file_mangostine_proto_rawDescData)
	})
	return file_mangostine_proto_rawDescData
}

var file_mangostine_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mangostine_proto_goTypes = []interface{}{
	(*Transaction)(nil),         // 0: mangostine_v0.Transaction
	(*Send)(nil),                // 1: mangostine_v0.Send
	(*RawTransaction)(nil),      // 2: mangostine_v0.RawTransaction
	(*TransactionResponse)(nil), // 3: mangostine_v0.TransactionResponse
}
var file_mangostine_proto_depIdxs = []int32{
	1, // 0: mangostine_v0.Transaction.receiver:type_name -> mangostine_v0.Send
	0, // 1: mangostine_v0.Rpc.SendTx:input_type -> mangostine_v0.Transaction
	2, // 2: mangostine_v0.Rpc.SendRawTx:input_type -> mangostine_v0.RawTransaction
	3, // 3: mangostine_v0.Rpc.SendTx:output_type -> mangostine_v0.TransactionResponse
	3, // 4: mangostine_v0.Rpc.SendRawTx:output_type -> mangostine_v0.TransactionResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mangostine_proto_init() }
func file_mangostine_proto_init() {
	if File_mangostine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mangostine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_mangostine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Send); i {
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
		file_mangostine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawTransaction); i {
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
		file_mangostine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
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
			RawDescriptor: file_mangostine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mangostine_proto_goTypes,
		DependencyIndexes: file_mangostine_proto_depIdxs,
		MessageInfos:      file_mangostine_proto_msgTypes,
	}.Build()
	File_mangostine_proto = out.File
	file_mangostine_proto_rawDesc = nil
	file_mangostine_proto_goTypes = nil
	file_mangostine_proto_depIdxs = nil
}

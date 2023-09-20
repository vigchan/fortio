// regenerate the .pb.go file after any change using
// protoc ping.proto --go_out=plugins=grpc:.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: ping.proto

package fgrpc

import (
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

type PingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq            int64  `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`                       // sequence number
	Ts             int64  `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"`                         // src send ts / dest receive ts
	Payload        string `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`                // extra packet data
	DelayNanos     int64  `protobuf:"varint,4,opt,name=delayNanos,proto3" json:"delayNanos,omitempty"`         // delay the response by x nanoseconds
	ResponseLength int64  `protobuf:"varint,5,opt,name=responseLength,proto3" json:"responseLength,omitempty"` // Request the length of the response payload
	ResponseBytes  []byte `protobuf:"bytes,6,opt,name=responseBytes,proto3" json:"responseBytes,omitempty"`    // The response payload in bytes
}

func (x *PingMessage) Reset() {
	*x = PingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingMessage) ProtoMessage() {}

func (x *PingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingMessage.ProtoReflect.Descriptor instead.
func (*PingMessage) Descriptor() ([]byte, []int) {
	return file_ping_proto_rawDescGZIP(), []int{0}
}

func (x *PingMessage) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *PingMessage) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *PingMessage) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *PingMessage) GetDelayNanos() int64 {
	if x != nil {
		return x.DelayNanos
	}
	return 0
}

func (x *PingMessage) GetResponseLength() int64 {
	if x != nil {
		return x.ResponseLength
	}
	return 0
}

func (x *PingMessage) GetResponseBytes() []byte {
	if x != nil {
		return x.ResponseBytes
	}
	return nil
}

var File_ping_proto protoreflect.FileDescriptor

var file_ping_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x67,
	0x72, 0x70, 0x63, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6e, 0x6f, 0x73, 0x12,
	0x26, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x32, 0x3e, 0x0a,
	0x0a, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x66, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x66, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2e, 0x2f, 0x66, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ping_proto_rawDescOnce sync.Once
	file_ping_proto_rawDescData = file_ping_proto_rawDesc
)

func file_ping_proto_rawDescGZIP() []byte {
	file_ping_proto_rawDescOnce.Do(func() {
		file_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_ping_proto_rawDescData)
	})
	return file_ping_proto_rawDescData
}

var file_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ping_proto_goTypes = []interface{}{
	(*PingMessage)(nil), // 0: fgrpc.PingMessage
}
var file_ping_proto_depIdxs = []int32{
	0, // 0: fgrpc.PingServer.Ping:input_type -> fgrpc.PingMessage
	0, // 1: fgrpc.PingServer.Ping:output_type -> fgrpc.PingMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ping_proto_init() }
func file_ping_proto_init() {
	if File_ping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingMessage); i {
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
			RawDescriptor: file_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ping_proto_goTypes,
		DependencyIndexes: file_ping_proto_depIdxs,
		MessageInfos:      file_ping_proto_msgTypes,
	}.Build()
	File_ping_proto = out.File
	file_ping_proto_rawDesc = nil
	file_ping_proto_goTypes = nil
	file_ping_proto_depIdxs = nil
}

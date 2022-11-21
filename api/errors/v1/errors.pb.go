// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: errors/v1/errors.proto

package errorsv1

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

// The `Status` type defines a logical error model that is suitable for
// different programming environments, including REST APIs and RPC APIs.
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A simple error code that can be easily handled by the client. The
	// actual error code is defined by `google.rpc.Code`.
	Code Code `protobuf:"varint,1,opt,name=code,proto3,enum=errors.v1.Code" json:"code,omitempty"`
	// A developer-facing human-readable error message in English. It should
	// both explain the error and offer an actionable resolution to it.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_v1_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_errors_v1_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_errors_v1_errors_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_CODE_UNSPECIFIED
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_errors_v1_errors_proto protoreflect.FileDescriptor

var file_errors_v1_errors_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x46, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x89, 0x01, 0x0a, 0x15, 0x69, 0x6f, 0x2e,
	0x61, 0x31, 0x33, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x1e, 0x61, 0x31, 0x33, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x15, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_v1_errors_proto_rawDescOnce sync.Once
	file_errors_v1_errors_proto_rawDescData = file_errors_v1_errors_proto_rawDesc
)

func file_errors_v1_errors_proto_rawDescGZIP() []byte {
	file_errors_v1_errors_proto_rawDescOnce.Do(func() {
		file_errors_v1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_v1_errors_proto_rawDescData)
	})
	return file_errors_v1_errors_proto_rawDescData
}

var file_errors_v1_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_v1_errors_proto_goTypes = []interface{}{
	(*Error)(nil), // 0: errors.v1.Error
	(Code)(0),     // 1: errors.v1.Code
}
var file_errors_v1_errors_proto_depIdxs = []int32{
	1, // 0: errors.v1.Error.code:type_name -> errors.v1.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_errors_v1_errors_proto_init() }
func file_errors_v1_errors_proto_init() {
	if File_errors_v1_errors_proto != nil {
		return
	}
	file_errors_v1_error_codes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_errors_v1_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_errors_v1_errors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_v1_errors_proto_goTypes,
		DependencyIndexes: file_errors_v1_errors_proto_depIdxs,
		MessageInfos:      file_errors_v1_errors_proto_msgTypes,
	}.Build()
	File_errors_v1_errors_proto = out.File
	file_errors_v1_errors_proto_rawDesc = nil
	file_errors_v1_errors_proto_goTypes = nil
	file_errors_v1_errors_proto_depIdxs = nil
}

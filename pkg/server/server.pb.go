// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: pkg/server/server.proto

package server

import (
	database "github.com/mxplusb/pleiades/api/v1/database"
	raft "github.com/mxplusb/pleiades/api/v1/raft"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_pkg_server_server_proto protoreflect.FileDescriptor

var file_pkg_server_server_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x72, 0x61, 0x66,
	0x74, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x5f,
	0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x91, 0x06, 0x0a, 0x0c, 0x53,
	0x68, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0a, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x17, 0x2e, 0x72, 0x61, 0x66, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x54, 0x0a, 0x12, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x1f, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x51, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x57, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x61,
	0x66, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x36, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x61, 0x72, 0x64, 0x12, 0x15, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x68,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x45, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x1a, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x42, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x19, 0x2e,
	0x72, 0x61, 0x66, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x4a, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x72, 0x61, 0x66, 0x74,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3f, 0x0a,
	0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x18, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x53, 0x74,
	0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xea,
	0x01, 0x0a, 0x08, 0x52, 0x61, 0x66, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x45, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x1a, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x72, 0x61, 0x66, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x12, 0x15, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x61, 0x66,
	0x74, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x2a, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x11, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x53,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x61, 0x66,
	0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xf2, 0x01, 0x0a, 0x0c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x50, 0x0a, 0x0e,
	0x4e, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x56,
	0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x12, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x32, 0x89, 0x04, 0x0a, 0x0e, 0x4b, 0x56, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x4c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x50,
	0x75, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x50, 0x75, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x75, 0x74, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b,
	0x61, 0x31, 0x33, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6c, 0x65, 0x69, 0x61, 0x64, 0x65, 0x73,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_pkg_server_server_proto_goTypes = []interface{}{
	(*raft.AddReplicaRequest)(nil),           // 0: raft.AddReplicaRequest
	(*raft.AddReplicaObserverRequest)(nil),   // 1: raft.AddReplicaObserverRequest
	(*raft.AddReplicaWitnessRequest)(nil),    // 2: raft.AddReplicaWitnessRequest
	(*raft.GetLeaderIdRequest)(nil),          // 3: raft.GetLeaderIdRequest
	(*raft.GetShardMembersRequest)(nil),      // 4: raft.GetShardMembersRequest
	(*raft.NewShardRequest)(nil),             // 5: raft.NewShardRequest
	(*raft.RemoveDataRequest)(nil),           // 6: raft.RemoveDataRequest
	(*raft.DeleteReplicaRequest)(nil),        // 7: raft.DeleteReplicaRequest
	(*raft.StartReplicaRequest)(nil),         // 8: raft.StartReplicaRequest
	(*raft.StopReplicaRequest)(nil),          // 9: raft.StopReplicaRequest
	(*raft.CompactRequest)(nil),              // 10: raft.CompactRequest
	(*raft.GetHostConfigRequest)(nil),        // 11: raft.GetHostConfigRequest
	(*raft.SnapshotRequest)(nil),             // 12: raft.SnapshotRequest
	(*raft.StopRequest)(nil),                 // 13: raft.StopRequest
	(*database.NewTransactionRequest)(nil),   // 14: database.NewTransactionRequest
	(*database.CloseTransactionRequest)(nil), // 15: database.CloseTransactionRequest
	(*database.CommitRequest)(nil),           // 16: database.CommitRequest
	(*database.CreateAccountRequest)(nil),    // 17: database.CreateAccountRequest
	(*database.DeleteBucketRequest)(nil),     // 18: database.DeleteBucketRequest
	(*database.CreateBucketRequest)(nil),     // 19: database.CreateBucketRequest
	(*database.GetKeyRequest)(nil),           // 20: database.GetKeyRequest
	(*database.PutKeyRequest)(nil),           // 21: database.PutKeyRequest
	(*database.DeleteKeyRequest)(nil),        // 22: database.DeleteKeyRequest
	(*raft.AddReplicaReply)(nil),             // 23: raft.AddReplicaReply
	(*raft.AddReplicaObserverReply)(nil),     // 24: raft.AddReplicaObserverReply
	(*raft.AddReplicaWitnessReply)(nil),      // 25: raft.AddReplicaWitnessReply
	(*raft.GetLeaderIdReply)(nil),            // 26: raft.GetLeaderIdReply
	(*raft.GetShardMembersReply)(nil),        // 27: raft.GetShardMembersReply
	(*raft.NewShardReply)(nil),               // 28: raft.NewShardReply
	(*raft.RemoveDataReply)(nil),             // 29: raft.RemoveDataReply
	(*raft.DeleteReplicaReply)(nil),          // 30: raft.DeleteReplicaReply
	(*raft.StartReplicaReply)(nil),           // 31: raft.StartReplicaReply
	(*raft.StopReplicaReply)(nil),            // 32: raft.StopReplicaReply
	(*raft.CompactReply)(nil),                // 33: raft.CompactReply
	(*raft.GetHostConfigReply)(nil),          // 34: raft.GetHostConfigReply
	(*raft.SnapshotReply)(nil),               // 35: raft.SnapshotReply
	(*raft.StopReply)(nil),                   // 36: raft.StopReply
	(*database.NewTransactionReply)(nil),     // 37: database.NewTransactionReply
	(*database.CloseTransactionReply)(nil),   // 38: database.CloseTransactionReply
	(*database.CommitReply)(nil),             // 39: database.CommitReply
	(*database.CreateAccountReply)(nil),      // 40: database.CreateAccountReply
	(*database.DeleteBucketReply)(nil),       // 41: database.DeleteBucketReply
	(*database.CreateBucketReply)(nil),       // 42: database.CreateBucketReply
	(*database.GetKeyReply)(nil),             // 43: database.GetKeyReply
	(*database.PutKeyReply)(nil),             // 44: database.PutKeyReply
	(*database.DeleteKeyReply)(nil),          // 45: database.DeleteKeyReply
}
var file_pkg_server_server_proto_depIdxs = []int32{
	0,  // 0: server.ShardManager.AddReplica:input_type -> raft.AddReplicaRequest
	1,  // 1: server.ShardManager.AddReplicaObserver:input_type -> raft.AddReplicaObserverRequest
	2,  // 2: server.ShardManager.AddReplicaWitness:input_type -> raft.AddReplicaWitnessRequest
	3,  // 3: server.ShardManager.GetLeaderId:input_type -> raft.GetLeaderIdRequest
	4,  // 4: server.ShardManager.GetShardMembers:input_type -> raft.GetShardMembersRequest
	5,  // 5: server.ShardManager.NewShard:input_type -> raft.NewShardRequest
	6,  // 6: server.ShardManager.RemoveData:input_type -> raft.RemoveDataRequest
	7,  // 7: server.ShardManager.RemoveReplica:input_type -> raft.DeleteReplicaRequest
	8,  // 8: server.ShardManager.StartReplica:input_type -> raft.StartReplicaRequest
	8,  // 9: server.ShardManager.StartReplicaObserver:input_type -> raft.StartReplicaRequest
	9,  // 10: server.ShardManager.StopReplica:input_type -> raft.StopReplicaRequest
	10, // 11: server.RaftHost.Compact:input_type -> raft.CompactRequest
	11, // 12: server.RaftHost.GetHostConfig:input_type -> raft.GetHostConfigRequest
	12, // 13: server.RaftHost.Snapshot:input_type -> raft.SnapshotRequest
	13, // 14: server.RaftHost.Stop:input_type -> raft.StopRequest
	14, // 15: server.Transactions.NewTransaction:input_type -> database.NewTransactionRequest
	15, // 16: server.Transactions.CloseTransaction:input_type -> database.CloseTransactionRequest
	16, // 17: server.Transactions.Commit:input_type -> database.CommitRequest
	17, // 18: server.KVStoreService.CreateAccount:input_type -> database.CreateAccountRequest
	18, // 19: server.KVStoreService.DeleteAccount:input_type -> database.DeleteBucketRequest
	19, // 20: server.KVStoreService.CreateBucket:input_type -> database.CreateBucketRequest
	18, // 21: server.KVStoreService.DeleteBucket:input_type -> database.DeleteBucketRequest
	20, // 22: server.KVStoreService.GetKey:input_type -> database.GetKeyRequest
	21, // 23: server.KVStoreService.PutKey:input_type -> database.PutKeyRequest
	22, // 24: server.KVStoreService.DeleteKey:input_type -> database.DeleteKeyRequest
	23, // 25: server.ShardManager.AddReplica:output_type -> raft.AddReplicaReply
	24, // 26: server.ShardManager.AddReplicaObserver:output_type -> raft.AddReplicaObserverReply
	25, // 27: server.ShardManager.AddReplicaWitness:output_type -> raft.AddReplicaWitnessReply
	26, // 28: server.ShardManager.GetLeaderId:output_type -> raft.GetLeaderIdReply
	27, // 29: server.ShardManager.GetShardMembers:output_type -> raft.GetShardMembersReply
	28, // 30: server.ShardManager.NewShard:output_type -> raft.NewShardReply
	29, // 31: server.ShardManager.RemoveData:output_type -> raft.RemoveDataReply
	30, // 32: server.ShardManager.RemoveReplica:output_type -> raft.DeleteReplicaReply
	31, // 33: server.ShardManager.StartReplica:output_type -> raft.StartReplicaReply
	31, // 34: server.ShardManager.StartReplicaObserver:output_type -> raft.StartReplicaReply
	32, // 35: server.ShardManager.StopReplica:output_type -> raft.StopReplicaReply
	33, // 36: server.RaftHost.Compact:output_type -> raft.CompactReply
	34, // 37: server.RaftHost.GetHostConfig:output_type -> raft.GetHostConfigReply
	35, // 38: server.RaftHost.Snapshot:output_type -> raft.SnapshotReply
	36, // 39: server.RaftHost.Stop:output_type -> raft.StopReply
	37, // 40: server.Transactions.NewTransaction:output_type -> database.NewTransactionReply
	38, // 41: server.Transactions.CloseTransaction:output_type -> database.CloseTransactionReply
	39, // 42: server.Transactions.Commit:output_type -> database.CommitReply
	40, // 43: server.KVStoreService.CreateAccount:output_type -> database.CreateAccountReply
	41, // 44: server.KVStoreService.DeleteAccount:output_type -> database.DeleteBucketReply
	42, // 45: server.KVStoreService.CreateBucket:output_type -> database.CreateBucketReply
	41, // 46: server.KVStoreService.DeleteBucket:output_type -> database.DeleteBucketReply
	43, // 47: server.KVStoreService.GetKey:output_type -> database.GetKeyReply
	44, // 48: server.KVStoreService.PutKey:output_type -> database.PutKeyReply
	45, // 49: server.KVStoreService.DeleteKey:output_type -> database.DeleteKeyReply
	25, // [25:50] is the sub-list for method output_type
	0,  // [0:25] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_server_server_proto_init() }
func file_pkg_server_server_proto_init() {
	if File_pkg_server_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_server_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_pkg_server_server_proto_goTypes,
		DependencyIndexes: file_pkg_server_server_proto_depIdxs,
	}.Build()
	File_pkg_server_server_proto = out.File
	file_pkg_server_server_proto_rawDesc = nil
	file_pkg_server_server_proto_goTypes = nil
	file_pkg_server_server_proto_depIdxs = nil
}

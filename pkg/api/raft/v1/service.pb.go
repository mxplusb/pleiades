// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: raft/v1/service.proto

package raftv1

import (
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

var File_raft_v1_service_proto protoreflect.FileDescriptor

var file_raft_v1_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x68,
	0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72, 0x61, 0x66, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x84, 0x07, 0x0a, 0x0c, 0x53, 0x68, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x12, 0x1a, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x12, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x22, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x11, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12,
	0x21, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x54, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x61,
	0x72, 0x64, 0x12, 0x18, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77,
	0x53, 0x68, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e,
	0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12,
	0x1d, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x1c,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x14, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x61, 0x66, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12,
	0x1b, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x91, 0x02, 0x0a, 0x0b, 0x48,
	0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48,
	0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x2e, 0x72, 0x61, 0x66, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x12, 0x18, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x53, 0x74, 0x6f,
	0x70, 0x12, 0x14, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa8,
	0x01, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x6d, 0x78,
	0x70, 0x6c, 0x75, 0x73, 0x62, 0x2e, 0x70, 0x6c, 0x65, 0x69, 0x61, 0x64, 0x65, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x78, 0x70, 0x6c, 0x75, 0x73, 0x62, 0x2f, 0x70,
	0x6c, 0x65, 0x69, 0x61, 0x64, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x61, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x61, 0x66, 0x74, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x52, 0x61, 0x66, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x07, 0x52, 0x61, 0x66, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x52, 0x61, 0x66, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x08, 0x52, 0x61, 0x66, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_raft_v1_service_proto_goTypes = []interface{}{
	(*AddReplicaRequest)(nil),            // 0: raft.v1.AddReplicaRequest
	(*AddReplicaObserverRequest)(nil),    // 1: raft.v1.AddReplicaObserverRequest
	(*AddReplicaWitnessRequest)(nil),     // 2: raft.v1.AddReplicaWitnessRequest
	(*GetLeaderIdRequest)(nil),           // 3: raft.v1.GetLeaderIdRequest
	(*GetShardMembersRequest)(nil),       // 4: raft.v1.GetShardMembersRequest
	(*NewShardRequest)(nil),              // 5: raft.v1.NewShardRequest
	(*RemoveDataRequest)(nil),            // 6: raft.v1.RemoveDataRequest
	(*RemoveReplicaRequest)(nil),         // 7: raft.v1.RemoveReplicaRequest
	(*StartReplicaRequest)(nil),          // 8: raft.v1.StartReplicaRequest
	(*StartReplicaObserverRequest)(nil),  // 9: raft.v1.StartReplicaObserverRequest
	(*StopReplicaRequest)(nil),           // 10: raft.v1.StopReplicaRequest
	(*CompactRequest)(nil),               // 11: raft.v1.CompactRequest
	(*GetHostConfigRequest)(nil),         // 12: raft.v1.GetHostConfigRequest
	(*SnapshotRequest)(nil),              // 13: raft.v1.SnapshotRequest
	(*StopRequest)(nil),                  // 14: raft.v1.StopRequest
	(*AddReplicaResponse)(nil),           // 15: raft.v1.AddReplicaResponse
	(*AddReplicaObserverResponse)(nil),   // 16: raft.v1.AddReplicaObserverResponse
	(*AddReplicaWitnessResponse)(nil),    // 17: raft.v1.AddReplicaWitnessResponse
	(*GetLeaderIdResponse)(nil),          // 18: raft.v1.GetLeaderIdResponse
	(*GetShardMembersResponse)(nil),      // 19: raft.v1.GetShardMembersResponse
	(*NewShardResponse)(nil),             // 20: raft.v1.NewShardResponse
	(*RemoveDataResponse)(nil),           // 21: raft.v1.RemoveDataResponse
	(*RemoveReplicaResponse)(nil),        // 22: raft.v1.RemoveReplicaResponse
	(*StartReplicaResponse)(nil),         // 23: raft.v1.StartReplicaResponse
	(*StartReplicaObserverResponse)(nil), // 24: raft.v1.StartReplicaObserverResponse
	(*StopReplicaResponse)(nil),          // 25: raft.v1.StopReplicaResponse
	(*CompactResponse)(nil),              // 26: raft.v1.CompactResponse
	(*GetHostConfigResponse)(nil),        // 27: raft.v1.GetHostConfigResponse
	(*SnapshotResponse)(nil),             // 28: raft.v1.SnapshotResponse
	(*StopResponse)(nil),                 // 29: raft.v1.StopResponse
}
var file_raft_v1_service_proto_depIdxs = []int32{
	0,  // 0: raft.v1.ShardService.AddReplica:input_type -> raft.v1.AddReplicaRequest
	1,  // 1: raft.v1.ShardService.AddReplicaObserver:input_type -> raft.v1.AddReplicaObserverRequest
	2,  // 2: raft.v1.ShardService.AddReplicaWitness:input_type -> raft.v1.AddReplicaWitnessRequest
	3,  // 3: raft.v1.ShardService.GetLeaderId:input_type -> raft.v1.GetLeaderIdRequest
	4,  // 4: raft.v1.ShardService.GetShardMembers:input_type -> raft.v1.GetShardMembersRequest
	5,  // 5: raft.v1.ShardService.NewShard:input_type -> raft.v1.NewShardRequest
	6,  // 6: raft.v1.ShardService.RemoveData:input_type -> raft.v1.RemoveDataRequest
	7,  // 7: raft.v1.ShardService.RemoveReplica:input_type -> raft.v1.RemoveReplicaRequest
	8,  // 8: raft.v1.ShardService.StartReplica:input_type -> raft.v1.StartReplicaRequest
	9,  // 9: raft.v1.ShardService.StartReplicaObserver:input_type -> raft.v1.StartReplicaObserverRequest
	10, // 10: raft.v1.ShardService.StopReplica:input_type -> raft.v1.StopReplicaRequest
	11, // 11: raft.v1.HostService.Compact:input_type -> raft.v1.CompactRequest
	12, // 12: raft.v1.HostService.GetHostConfig:input_type -> raft.v1.GetHostConfigRequest
	13, // 13: raft.v1.HostService.Snapshot:input_type -> raft.v1.SnapshotRequest
	14, // 14: raft.v1.HostService.Stop:input_type -> raft.v1.StopRequest
	15, // 15: raft.v1.ShardService.AddReplica:output_type -> raft.v1.AddReplicaResponse
	16, // 16: raft.v1.ShardService.AddReplicaObserver:output_type -> raft.v1.AddReplicaObserverResponse
	17, // 17: raft.v1.ShardService.AddReplicaWitness:output_type -> raft.v1.AddReplicaWitnessResponse
	18, // 18: raft.v1.ShardService.GetLeaderId:output_type -> raft.v1.GetLeaderIdResponse
	19, // 19: raft.v1.ShardService.GetShardMembers:output_type -> raft.v1.GetShardMembersResponse
	20, // 20: raft.v1.ShardService.NewShard:output_type -> raft.v1.NewShardResponse
	21, // 21: raft.v1.ShardService.RemoveData:output_type -> raft.v1.RemoveDataResponse
	22, // 22: raft.v1.ShardService.RemoveReplica:output_type -> raft.v1.RemoveReplicaResponse
	23, // 23: raft.v1.ShardService.StartReplica:output_type -> raft.v1.StartReplicaResponse
	24, // 24: raft.v1.ShardService.StartReplicaObserver:output_type -> raft.v1.StartReplicaObserverResponse
	25, // 25: raft.v1.ShardService.StopReplica:output_type -> raft.v1.StopReplicaResponse
	26, // 26: raft.v1.HostService.Compact:output_type -> raft.v1.CompactResponse
	27, // 27: raft.v1.HostService.GetHostConfig:output_type -> raft.v1.GetHostConfigResponse
	28, // 28: raft.v1.HostService.Snapshot:output_type -> raft.v1.SnapshotResponse
	29, // 29: raft.v1.HostService.Stop:output_type -> raft.v1.StopResponse
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_raft_v1_service_proto_init() }
func file_raft_v1_service_proto_init() {
	if File_raft_v1_service_proto != nil {
		return
	}
	file_raft_v1_raft_host_proto_init()
	file_raft_v1_raft_shard_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_raft_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_raft_v1_service_proto_goTypes,
		DependencyIndexes: file_raft_v1_service_proto_depIdxs,
	}.Build()
	File_raft_v1_service_proto = out.File
	file_raft_v1_service_proto_rawDesc = nil
	file_raft_v1_service_proto_goTypes = nil
	file_raft_v1_service_proto_depIdxs = nil
}

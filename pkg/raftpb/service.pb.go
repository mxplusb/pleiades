// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: raftpb/service.proto

package raftpb

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

var File_raftpb_service_proto protoreflect.FileDescriptor

var file_raftpb_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x1a, 0x16,
	0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2f, 0x72,
	0x61, 0x66, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xf8, 0x06, 0x0a, 0x0c, 0x53, 0x68, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x43, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x19,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x61, 0x66, 0x74,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x72, 0x61,
	0x66, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x58, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x20, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x57, 0x69, 0x74, 0x6e, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x74,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x57, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x72, 0x61,
	0x66, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12, 0x57, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90,
	0x02, 0x01, 0x12, 0x3d, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x61, 0x72, 0x64, 0x12, 0x17,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62,
	0x2e, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x43, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x19, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x61, 0x66,
	0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x1c, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x12, 0x1b, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x61, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x12, 0x1a, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8e, 0x02, 0x0a, 0x0b, 0x48,
	0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12, 0x3d, 0x0a, 0x08, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x17, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70,
	0x12, 0x13, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x70, 0x62, 0x2e, 0x53,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x78, 0x70, 0x6c, 0x75, 0x73,
	0x62, 0x2f, 0x70, 0x6c, 0x65, 0x69, 0x61, 0x64, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72,
	0x61, 0x66, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_raftpb_service_proto_goTypes = []interface{}{
	(*AddReplicaRequest)(nil),            // 0: raftpb.AddReplicaRequest
	(*AddReplicaObserverRequest)(nil),    // 1: raftpb.AddReplicaObserverRequest
	(*AddReplicaWitnessRequest)(nil),     // 2: raftpb.AddReplicaWitnessRequest
	(*GetLeaderIdRequest)(nil),           // 3: raftpb.GetLeaderIdRequest
	(*GetShardMembersRequest)(nil),       // 4: raftpb.GetShardMembersRequest
	(*NewShardRequest)(nil),              // 5: raftpb.NewShardRequest
	(*RemoveDataRequest)(nil),            // 6: raftpb.RemoveDataRequest
	(*RemoveReplicaRequest)(nil),         // 7: raftpb.RemoveReplicaRequest
	(*StartReplicaRequest)(nil),          // 8: raftpb.StartReplicaRequest
	(*StartReplicaObserverRequest)(nil),  // 9: raftpb.StartReplicaObserverRequest
	(*StopReplicaRequest)(nil),           // 10: raftpb.StopReplicaRequest
	(*CompactRequest)(nil),               // 11: raftpb.CompactRequest
	(*GetHostConfigRequest)(nil),         // 12: raftpb.GetHostConfigRequest
	(*SnapshotRequest)(nil),              // 13: raftpb.SnapshotRequest
	(*StopRequest)(nil),                  // 14: raftpb.StopRequest
	(*AddReplicaResponse)(nil),           // 15: raftpb.AddReplicaResponse
	(*AddReplicaObserverResponse)(nil),   // 16: raftpb.AddReplicaObserverResponse
	(*AddReplicaWitnessResponse)(nil),    // 17: raftpb.AddReplicaWitnessResponse
	(*GetLeaderIdResponse)(nil),          // 18: raftpb.GetLeaderIdResponse
	(*GetShardMembersResponse)(nil),      // 19: raftpb.GetShardMembersResponse
	(*NewShardResponse)(nil),             // 20: raftpb.NewShardResponse
	(*RemoveDataResponse)(nil),           // 21: raftpb.RemoveDataResponse
	(*RemoveReplicaResponse)(nil),        // 22: raftpb.RemoveReplicaResponse
	(*StartReplicaResponse)(nil),         // 23: raftpb.StartReplicaResponse
	(*StartReplicaObserverResponse)(nil), // 24: raftpb.StartReplicaObserverResponse
	(*StopReplicaResponse)(nil),          // 25: raftpb.StopReplicaResponse
	(*CompactResponse)(nil),              // 26: raftpb.CompactResponse
	(*GetHostConfigResponse)(nil),        // 27: raftpb.GetHostConfigResponse
	(*SnapshotResponse)(nil),             // 28: raftpb.SnapshotResponse
	(*StopResponse)(nil),                 // 29: raftpb.StopResponse
}
var file_raftpb_service_proto_depIdxs = []int32{
	0,  // 0: raftpb.ShardService.AddReplica:input_type -> raftpb.AddReplicaRequest
	1,  // 1: raftpb.ShardService.AddReplicaObserver:input_type -> raftpb.AddReplicaObserverRequest
	2,  // 2: raftpb.ShardService.AddReplicaWitness:input_type -> raftpb.AddReplicaWitnessRequest
	3,  // 3: raftpb.ShardService.GetLeaderId:input_type -> raftpb.GetLeaderIdRequest
	4,  // 4: raftpb.ShardService.GetShardMembers:input_type -> raftpb.GetShardMembersRequest
	5,  // 5: raftpb.ShardService.NewShard:input_type -> raftpb.NewShardRequest
	6,  // 6: raftpb.ShardService.RemoveData:input_type -> raftpb.RemoveDataRequest
	7,  // 7: raftpb.ShardService.RemoveReplica:input_type -> raftpb.RemoveReplicaRequest
	8,  // 8: raftpb.ShardService.StartReplica:input_type -> raftpb.StartReplicaRequest
	9,  // 9: raftpb.ShardService.StartReplicaObserver:input_type -> raftpb.StartReplicaObserverRequest
	10, // 10: raftpb.ShardService.StopReplica:input_type -> raftpb.StopReplicaRequest
	11, // 11: raftpb.HostService.Compact:input_type -> raftpb.CompactRequest
	12, // 12: raftpb.HostService.GetHostConfig:input_type -> raftpb.GetHostConfigRequest
	13, // 13: raftpb.HostService.Snapshot:input_type -> raftpb.SnapshotRequest
	14, // 14: raftpb.HostService.Stop:input_type -> raftpb.StopRequest
	15, // 15: raftpb.ShardService.AddReplica:output_type -> raftpb.AddReplicaResponse
	16, // 16: raftpb.ShardService.AddReplicaObserver:output_type -> raftpb.AddReplicaObserverResponse
	17, // 17: raftpb.ShardService.AddReplicaWitness:output_type -> raftpb.AddReplicaWitnessResponse
	18, // 18: raftpb.ShardService.GetLeaderId:output_type -> raftpb.GetLeaderIdResponse
	19, // 19: raftpb.ShardService.GetShardMembers:output_type -> raftpb.GetShardMembersResponse
	20, // 20: raftpb.ShardService.NewShard:output_type -> raftpb.NewShardResponse
	21, // 21: raftpb.ShardService.RemoveData:output_type -> raftpb.RemoveDataResponse
	22, // 22: raftpb.ShardService.RemoveReplica:output_type -> raftpb.RemoveReplicaResponse
	23, // 23: raftpb.ShardService.StartReplica:output_type -> raftpb.StartReplicaResponse
	24, // 24: raftpb.ShardService.StartReplicaObserver:output_type -> raftpb.StartReplicaObserverResponse
	25, // 25: raftpb.ShardService.StopReplica:output_type -> raftpb.StopReplicaResponse
	26, // 26: raftpb.HostService.Compact:output_type -> raftpb.CompactResponse
	27, // 27: raftpb.HostService.GetHostConfig:output_type -> raftpb.GetHostConfigResponse
	28, // 28: raftpb.HostService.Snapshot:output_type -> raftpb.SnapshotResponse
	29, // 29: raftpb.HostService.Stop:output_type -> raftpb.StopResponse
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_raftpb_service_proto_init() }
func file_raftpb_service_proto_init() {
	if File_raftpb_service_proto != nil {
		return
	}
	file_raftpb_raft_host_proto_init()
	file_raftpb_raft_shard_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_raftpb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_raftpb_service_proto_goTypes,
		DependencyIndexes: file_raftpb_service_proto_depIdxs,
	}.Build()
	File_raftpb_service_proto = out.File
	file_raftpb_service_proto_rawDesc = nil
	file_raftpb_service_proto_goTypes = nil
	file_raftpb_service_proto_depIdxs = nil
}
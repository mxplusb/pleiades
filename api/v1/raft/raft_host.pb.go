// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: api/v1/raft/raft_host.proto

package raft

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

type CompactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// replicaId is a non-zero value used to identify a node within a Raft cluster.
	ReplicaId uint64 `protobuf:"varint,1,opt,name=replicaId,proto3" json:"replicaId,omitempty"`
	// shardId is the unique value used to identify a Raft cluster.
	ShardId uint64 `protobuf:"varint,2,opt,name=shardId,proto3" json:"shardId,omitempty"`
}

func (x *CompactRequest) Reset() {
	*x = CompactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_raft_raft_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactRequest) ProtoMessage() {}

func (x *CompactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_raft_raft_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactRequest.ProtoReflect.Descriptor instead.
func (*CompactRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_raft_raft_host_proto_rawDescGZIP(), []int{0}
}

func (x *CompactRequest) GetReplicaId() uint64 {
	if x != nil {
		return x.ReplicaId
	}
	return 0
}

func (x *CompactRequest) GetShardId() uint64 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

type CompactReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompactReply) Reset() {
	*x = CompactReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_raft_raft_host_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactReply) ProtoMessage() {}

func (x *CompactReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_raft_raft_host_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactReply.ProtoReflect.Descriptor instead.
func (*CompactReply) Descriptor() ([]byte, []int) {
	return file_api_v1_raft_raft_host_proto_rawDescGZIP(), []int{1}
}

type LeaderTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// shardId is the unique value used to identify a Raft cluster.
	ShardId      uint64 `protobuf:"varint,1,opt,name=shardId,proto3" json:"shardId,omitempty"`
	TargetNodeId string `protobuf:"bytes,2,opt,name=targetNodeId,proto3" json:"targetNodeId,omitempty"`
}

func (x *LeaderTransferRequest) Reset() {
	*x = LeaderTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_raft_raft_host_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderTransferRequest) ProtoMessage() {}

func (x *LeaderTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_raft_raft_host_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderTransferRequest.ProtoReflect.Descriptor instead.
func (*LeaderTransferRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_raft_raft_host_proto_rawDescGZIP(), []int{2}
}

func (x *LeaderTransferRequest) GetShardId() uint64 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *LeaderTransferRequest) GetTargetNodeId() string {
	if x != nil {
		return x.TargetNodeId
	}
	return ""
}

type LeaderTransferReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaderTransferReply) Reset() {
	*x = LeaderTransferReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_raft_raft_host_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderTransferReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderTransferReply) ProtoMessage() {}

func (x *LeaderTransferReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_raft_raft_host_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderTransferReply.ProtoReflect.Descriptor instead.
func (*LeaderTransferReply) Descriptor() ([]byte, []int) {
	return file_api_v1_raft_raft_host_proto_rawDescGZIP(), []int{3}
}

type SnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// shardId is the unique value used to identify a Raft cluster.
	ShardId uint64 `protobuf:"varint,1,opt,name=shardId,proto3" json:"shardId,omitempty"`
	Timeout int64  `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *SnapshotRequest) Reset() {
	*x = SnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_raft_raft_host_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotRequest) ProtoMessage() {}

func (x *SnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_raft_raft_host_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotRequest.ProtoReflect.Descriptor instead.
func (*SnapshotRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_raft_raft_host_proto_rawDescGZIP(), []int{4}
}

func (x *SnapshotRequest) GetShardId() uint64 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *SnapshotRequest) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type SnapshotReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SnapshotReply) Reset() {
	*x = SnapshotReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_raft_raft_host_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotReply) ProtoMessage() {}

func (x *SnapshotReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_raft_raft_host_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotReply.ProtoReflect.Descriptor instead.
func (*SnapshotReply) Descriptor() ([]byte, []int) {
	return file_api_v1_raft_raft_host_proto_rawDescGZIP(), []int{5}
}

type StopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_raft_raft_host_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_raft_raft_host_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_raft_raft_host_proto_rawDescGZIP(), []int{6}
}

type StopReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopReply) Reset() {
	*x = StopReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_raft_raft_host_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopReply) ProtoMessage() {}

func (x *StopReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_raft_raft_host_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopReply.ProtoReflect.Descriptor instead.
func (*StopReply) Descriptor() ([]byte, []int) {
	return file_api_v1_raft_raft_host_proto_rawDescGZIP(), []int{7}
}

type GetHostConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHostConfigRequest) Reset() {
	*x = GetHostConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_raft_raft_host_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostConfigRequest) ProtoMessage() {}

func (x *GetHostConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_raft_raft_host_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostConfigRequest.ProtoReflect.Descriptor instead.
func (*GetHostConfigRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_raft_raft_host_proto_rawDescGZIP(), []int{8}
}

type GetHostConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHostConfigReply) Reset() {
	*x = GetHostConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_raft_raft_host_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostConfigReply) ProtoMessage() {}

func (x *GetHostConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_raft_raft_host_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostConfigReply.ProtoReflect.Descriptor instead.
func (*GetHostConfigReply) Descriptor() ([]byte, []int) {
	return file_api_v1_raft_raft_host_proto_rawDescGZIP(), []int{9}
}

type GetClusterMembershipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId uint64 `protobuf:"varint,1,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
}

func (x *GetClusterMembershipRequest) Reset() {
	*x = GetClusterMembershipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_raft_raft_host_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterMembershipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterMembershipRequest) ProtoMessage() {}

func (x *GetClusterMembershipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_raft_raft_host_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterMembershipRequest.ProtoReflect.Descriptor instead.
func (*GetClusterMembershipRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_raft_raft_host_proto_rawDescGZIP(), []int{10}
}

func (x *GetClusterMembershipRequest) GetClusterId() uint64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

var File_api_v1_raft_raft_host_proto protoreflect.FileDescriptor

var file_api_v1_raft_raft_host_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x72, 0x61,
	0x66, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72,
	0x61, 0x66, 0x74, 0x22, 0x48, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0x0e, 0x0a,
	0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x55, 0x0a,
	0x15, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x45, 0x0a, 0x0f, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x0b, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48, 0x6f,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3b, 0x0a,
	0x1b, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x42, 0x1e, 0x5a, 0x1c, 0x61, 0x31,
	0x33, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6c, 0x65, 0x69, 0x61, 0x64, 0x65, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_v1_raft_raft_host_proto_rawDescOnce sync.Once
	file_api_v1_raft_raft_host_proto_rawDescData = file_api_v1_raft_raft_host_proto_rawDesc
)

func file_api_v1_raft_raft_host_proto_rawDescGZIP() []byte {
	file_api_v1_raft_raft_host_proto_rawDescOnce.Do(func() {
		file_api_v1_raft_raft_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_raft_raft_host_proto_rawDescData)
	})
	return file_api_v1_raft_raft_host_proto_rawDescData
}

var file_api_v1_raft_raft_host_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_v1_raft_raft_host_proto_goTypes = []interface{}{
	(*CompactRequest)(nil),              // 0: raft.CompactRequest
	(*CompactReply)(nil),                // 1: raft.CompactReply
	(*LeaderTransferRequest)(nil),       // 2: raft.LeaderTransferRequest
	(*LeaderTransferReply)(nil),         // 3: raft.LeaderTransferReply
	(*SnapshotRequest)(nil),             // 4: raft.SnapshotRequest
	(*SnapshotReply)(nil),               // 5: raft.SnapshotReply
	(*StopRequest)(nil),                 // 6: raft.StopRequest
	(*StopReply)(nil),                   // 7: raft.StopReply
	(*GetHostConfigRequest)(nil),        // 8: raft.GetHostConfigRequest
	(*GetHostConfigReply)(nil),          // 9: raft.GetHostConfigReply
	(*GetClusterMembershipRequest)(nil), // 10: raft.GetClusterMembershipRequest
}
var file_api_v1_raft_raft_host_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_raft_raft_host_proto_init() }
func file_api_v1_raft_raft_host_proto_init() {
	if File_api_v1_raft_raft_host_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_raft_raft_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompactRequest); i {
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
		file_api_v1_raft_raft_host_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompactReply); i {
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
		file_api_v1_raft_raft_host_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderTransferRequest); i {
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
		file_api_v1_raft_raft_host_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderTransferReply); i {
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
		file_api_v1_raft_raft_host_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotRequest); i {
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
		file_api_v1_raft_raft_host_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotReply); i {
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
		file_api_v1_raft_raft_host_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRequest); i {
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
		file_api_v1_raft_raft_host_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopReply); i {
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
		file_api_v1_raft_raft_host_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostConfigRequest); i {
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
		file_api_v1_raft_raft_host_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostConfigReply); i {
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
		file_api_v1_raft_raft_host_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterMembershipRequest); i {
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
			RawDescriptor: file_api_v1_raft_raft_host_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_raft_raft_host_proto_goTypes,
		DependencyIndexes: file_api_v1_raft_raft_host_proto_depIdxs,
		MessageInfos:      file_api_v1_raft_raft_host_proto_msgTypes,
	}.Build()
	File_api_v1_raft_raft_host_proto = out.File
	file_api_v1_raft_raft_host_proto_rawDesc = nil
	file_api_v1_raft_raft_host_proto_goTypes = nil
	file_api_v1_raft_raft_host_proto_depIdxs = nil
}
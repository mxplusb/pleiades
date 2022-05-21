// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: raft.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RaftConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId                  uint64 `protobuf:"varint,1,opt,name=NodeId,proto3" json:"NodeId,omitempty"`
	ClusterId               uint64 `protobuf:"varint,2,opt,name=ClusterId,proto3" json:"ClusterId,omitempty"`
	CheckQuorum             bool   `protobuf:"varint,3,opt,name=CheckQuorum,proto3" json:"CheckQuorum,omitempty"`
	ElectionRoundTripTime   uint64 `protobuf:"varint,4,opt,name=ElectionRoundTripTime,proto3" json:"ElectionRoundTripTime,omitempty"`
	HeartbeatRoundTripTime  uint64 `protobuf:"varint,5,opt,name=HeartbeatRoundTripTime,proto3" json:"HeartbeatRoundTripTime,omitempty"`
	SnapshotEntries         uint64 `protobuf:"varint,6,opt,name=SnapshotEntries,proto3" json:"SnapshotEntries,omitempty"`
	CompactionOverhead      uint64 `protobuf:"varint,7,opt,name=CompactionOverhead,proto3" json:"CompactionOverhead,omitempty"`
	OrderedConfigChange     bool   `protobuf:"varint,8,opt,name=OrderedConfigChange,proto3" json:"OrderedConfigChange,omitempty"`
	MaxInMemLogSize         uint64 `protobuf:"varint,9,opt,name=MaxInMemLogSize,proto3" json:"MaxInMemLogSize,omitempty"`
	SnapshotCompressionType uint64 `protobuf:"varint,10,opt,name=SnapshotCompressionType,proto3" json:"SnapshotCompressionType,omitempty"`
	EntryCompressionType    uint64 `protobuf:"varint,11,opt,name=EntryCompressionType,proto3" json:"EntryCompressionType,omitempty"`
	DisableAutoCompactions  bool   `protobuf:"varint,12,opt,name=DisableAutoCompactions,proto3" json:"DisableAutoCompactions,omitempty"`
	IsObserver              bool   `protobuf:"varint,13,opt,name=IsObserver,proto3" json:"IsObserver,omitempty"`
	IsWitness               bool   `protobuf:"varint,14,opt,name=IsWitness,proto3" json:"IsWitness,omitempty"`
	Quiesce                 bool   `protobuf:"varint,15,opt,name=Quiesce,proto3" json:"Quiesce,omitempty"`
}

func (x *RaftConfig) Reset() {
	*x = RaftConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftConfig) ProtoMessage() {}

func (x *RaftConfig) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftConfig.ProtoReflect.Descriptor instead.
func (*RaftConfig) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{0}
}

func (x *RaftConfig) GetNodeId() uint64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *RaftConfig) GetClusterId() uint64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *RaftConfig) GetCheckQuorum() bool {
	if x != nil {
		return x.CheckQuorum
	}
	return false
}

func (x *RaftConfig) GetElectionRoundTripTime() uint64 {
	if x != nil {
		return x.ElectionRoundTripTime
	}
	return 0
}

func (x *RaftConfig) GetHeartbeatRoundTripTime() uint64 {
	if x != nil {
		return x.HeartbeatRoundTripTime
	}
	return 0
}

func (x *RaftConfig) GetSnapshotEntries() uint64 {
	if x != nil {
		return x.SnapshotEntries
	}
	return 0
}

func (x *RaftConfig) GetCompactionOverhead() uint64 {
	if x != nil {
		return x.CompactionOverhead
	}
	return 0
}

func (x *RaftConfig) GetOrderedConfigChange() bool {
	if x != nil {
		return x.OrderedConfigChange
	}
	return false
}

func (x *RaftConfig) GetMaxInMemLogSize() uint64 {
	if x != nil {
		return x.MaxInMemLogSize
	}
	return 0
}

func (x *RaftConfig) GetSnapshotCompressionType() uint64 {
	if x != nil {
		return x.SnapshotCompressionType
	}
	return 0
}

func (x *RaftConfig) GetEntryCompressionType() uint64 {
	if x != nil {
		return x.EntryCompressionType
	}
	return 0
}

func (x *RaftConfig) GetDisableAutoCompactions() bool {
	if x != nil {
		return x.DisableAutoCompactions
	}
	return false
}

func (x *RaftConfig) GetIsObserver() bool {
	if x != nil {
		return x.IsObserver
	}
	return false
}

func (x *RaftConfig) GetIsWitness() bool {
	if x != nil {
		return x.IsWitness
	}
	return false
}

func (x *RaftConfig) GetQuiesce() bool {
	if x != nil {
		return x.Quiesce
	}
	return false
}

type NewRaftConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable bool        `protobuf:"varint,1,opt,name=Enable,proto3" json:"Enable,omitempty"`
	Name   string      `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Config *RaftConfig `protobuf:"bytes,3,opt,name=Config,proto3" json:"Config,omitempty"`
}

func (x *NewRaftConfigRequest) Reset() {
	*x = NewRaftConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRaftConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRaftConfigRequest) ProtoMessage() {}

func (x *NewRaftConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRaftConfigRequest.ProtoReflect.Descriptor instead.
func (*NewRaftConfigRequest) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{1}
}

func (x *NewRaftConfigRequest) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *NewRaftConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewRaftConfigRequest) GetConfig() *RaftConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type InvalidRaftConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *InvalidRaftConfigResponse) Reset() {
	*x = InvalidRaftConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidRaftConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidRaftConfigResponse) ProtoMessage() {}

func (x *InvalidRaftConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidRaftConfigResponse.ProtoReflect.Descriptor instead.
func (*InvalidRaftConfigResponse) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{2}
}

func (x *InvalidRaftConfigResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *InvalidRaftConfigResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ValidRaftConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateApplied       *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=DateApplied,proto3" json:"DateApplied,omitempty"`
	Name              string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ActiveRaftConfigs int32                  `protobuf:"varint,3,opt,name=ActiveRaftConfigs,proto3" json:"ActiveRaftConfigs,omitempty"`
}

func (x *ValidRaftConfigResponse) Reset() {
	*x = ValidRaftConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidRaftConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidRaftConfigResponse) ProtoMessage() {}

func (x *ValidRaftConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidRaftConfigResponse.ProtoReflect.Descriptor instead.
func (*ValidRaftConfigResponse) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{3}
}

func (x *ValidRaftConfigResponse) GetDateApplied() *timestamppb.Timestamp {
	if x != nil {
		return x.DateApplied
	}
	return nil
}

func (x *ValidRaftConfigResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ValidRaftConfigResponse) GetActiveRaftConfigs() int32 {
	if x != nil {
		return x.ActiveRaftConfigs
	}
	return 0
}

type RaftConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid  bool        `protobuf:"varint,1,opt,name=Valid,proto3" json:"Valid,omitempty"`
	Name   string      `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Config *RaftConfig `protobuf:"bytes,3,opt,name=Config,proto3" json:"Config,omitempty"`
}

func (x *RaftConfigResponse) Reset() {
	*x = RaftConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftConfigResponse) ProtoMessage() {}

func (x *RaftConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftConfigResponse.ProtoReflect.Descriptor instead.
func (*RaftConfigResponse) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{4}
}

func (x *RaftConfigResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *RaftConfigResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RaftConfigResponse) GetConfig() *RaftConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type GetRaftConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetRaftConfigRequest) Reset() {
	*x = GetRaftConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRaftConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRaftConfigRequest) ProtoMessage() {}

func (x *GetRaftConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRaftConfigRequest.ProtoReflect.Descriptor instead.
func (*GetRaftConfigRequest) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{5}
}

func (x *GetRaftConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetRaftConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configuration *RaftConfig `protobuf:"bytes,1,opt,name=Configuration,proto3" json:"Configuration,omitempty"`
}

func (x *GetRaftConfigResponse) Reset() {
	*x = GetRaftConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRaftConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRaftConfigResponse) ProtoMessage() {}

func (x *GetRaftConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRaftConfigResponse.ProtoReflect.Descriptor instead.
func (*GetRaftConfigResponse) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{6}
}

func (x *GetRaftConfigResponse) GetConfiguration() *RaftConfig {
	if x != nil {
		return x.Configuration
	}
	return nil
}

type ListRaftConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRaftConfigs) Reset() {
	*x = ListRaftConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRaftConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRaftConfigs) ProtoMessage() {}

func (x *ListRaftConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRaftConfigs.ProtoReflect.Descriptor instead.
func (*ListRaftConfigs) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{7}
}

type ListRaftConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailableConfigs map[string]*RaftConfig `protobuf:"bytes,1,rep,name=AvailableConfigs,proto3" json:"AvailableConfigs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListRaftConfigsResponse) Reset() {
	*x = ListRaftConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRaftConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRaftConfigsResponse) ProtoMessage() {}

func (x *ListRaftConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRaftConfigsResponse.ProtoReflect.Descriptor instead.
func (*ListRaftConfigsResponse) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{8}
}

func (x *ListRaftConfigsResponse) GetAvailableConfigs() map[string]*RaftConfig {
	if x != nil {
		return x.AvailableConfigs
	}
	return nil
}

var File_raft_proto protoreflect.FileDescriptor

var file_raft_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x05,
	0x0a, 0x0a, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x51, 0x75, 0x6f, 0x72, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x51, 0x75,
	0x6f, 0x72, 0x75, 0x6d, 0x12, 0x34, 0x0a, 0x15, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x69, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x15, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x54, 0x72, 0x69, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x69, 0x70,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x69, 0x70, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x12,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x76, 0x65, 0x72, 0x68, 0x65,
	0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x76, 0x65, 0x72, 0x68, 0x65, 0x61, 0x64, 0x12, 0x30, 0x0a, 0x13,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x4d, 0x61, 0x78, 0x49, 0x6e, 0x4d, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x4d, 0x61, 0x78, 0x49, 0x6e, 0x4d, 0x65,
	0x6d, 0x4c, 0x6f, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x14, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x41, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41,
	0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x49, 0x73, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x49, 0x73, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x49, 0x73, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x51, 0x75, 0x69, 0x65, 0x73, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x51,
	0x75, 0x69, 0x65, 0x73, 0x63, 0x65, 0x22, 0x67, 0x0a, 0x14, 0x4e, 0x65, 0x77, 0x52, 0x61, 0x66,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x52, 0x61, 0x66,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x45, 0x0a, 0x19, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x17, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61,
	0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x11, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x22, 0x63, 0x0a, 0x12, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x22, 0xc7, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x66, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a,
	0x0a, 0x10, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x61, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x1a, 0x50, 0x0a, 0x15, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1b, 0x5a, 0x19,
	0x72, 0x33, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6c, 0x65, 0x69, 0x61, 0x64, 0x65, 0x73, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_raft_proto_rawDescOnce sync.Once
	file_raft_proto_rawDescData = file_raft_proto_rawDesc
)

func file_raft_proto_rawDescGZIP() []byte {
	file_raft_proto_rawDescOnce.Do(func() {
		file_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_raft_proto_rawDescData)
	})
	return file_raft_proto_rawDescData
}

var file_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_raft_proto_goTypes = []interface{}{
	(*RaftConfig)(nil),                // 0: RaftConfig
	(*NewRaftConfigRequest)(nil),      // 1: NewRaftConfigRequest
	(*InvalidRaftConfigResponse)(nil), // 2: InvalidRaftConfigResponse
	(*ValidRaftConfigResponse)(nil),   // 3: ValidRaftConfigResponse
	(*RaftConfigResponse)(nil),        // 4: RaftConfigResponse
	(*GetRaftConfigRequest)(nil),      // 5: GetRaftConfigRequest
	(*GetRaftConfigResponse)(nil),     // 6: GetRaftConfigResponse
	(*ListRaftConfigs)(nil),           // 7: ListRaftConfigs
	(*ListRaftConfigsResponse)(nil),   // 8: ListRaftConfigsResponse
	nil,                               // 9: ListRaftConfigsResponse.AvailableConfigsEntry
	(*timestamppb.Timestamp)(nil),     // 10: google.protobuf.Timestamp
}
var file_raft_proto_depIdxs = []int32{
	0,  // 0: NewRaftConfigRequest.Config:type_name -> RaftConfig
	10, // 1: ValidRaftConfigResponse.DateApplied:type_name -> google.protobuf.Timestamp
	0,  // 2: RaftConfigResponse.Config:type_name -> RaftConfig
	0,  // 3: GetRaftConfigResponse.Configuration:type_name -> RaftConfig
	9,  // 4: ListRaftConfigsResponse.AvailableConfigs:type_name -> ListRaftConfigsResponse.AvailableConfigsEntry
	0,  // 5: ListRaftConfigsResponse.AvailableConfigsEntry.value:type_name -> RaftConfig
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_raft_proto_init() }
func file_raft_proto_init() {
	if File_raft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftConfig); i {
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
		file_raft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRaftConfigRequest); i {
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
		file_raft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidRaftConfigResponse); i {
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
		file_raft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidRaftConfigResponse); i {
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
		file_raft_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftConfigResponse); i {
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
		file_raft_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRaftConfigRequest); i {
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
		file_raft_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRaftConfigResponse); i {
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
		file_raft_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRaftConfigs); i {
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
		file_raft_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRaftConfigsResponse); i {
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
			RawDescriptor: file_raft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_raft_proto_goTypes,
		DependencyIndexes: file_raft_proto_depIdxs,
		MessageInfos:      file_raft_proto_msgTypes,
	}.Build()
	File_raft_proto = out.File
	file_raft_proto_rawDesc = nil
	file_raft_proto_goTypes = nil
	file_raft_proto_depIdxs = nil
}

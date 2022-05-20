// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: raft_server.proto

package servers

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "r3t.io/pleiades/pkg/protos"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RaftConfigServiceClient is the client API for RaftConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftConfigServiceClient interface {
	AddConfiguration(ctx context.Context, in *protos.NewRaftConfigRequest, opts ...grpc.CallOption) (*protos.RaftConfigResponse, error)
	GetConfiguration(ctx context.Context, in *protos.GetRaftConfigRequest, opts ...grpc.CallOption) (*protos.GetRaftConfigResponse, error)
	ListConfigurations(ctx context.Context, in *protos.ListRaftConfigs, opts ...grpc.CallOption) (*protos.ListRaftConfigsResponse, error)
}

type raftConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftConfigServiceClient(cc grpc.ClientConnInterface) RaftConfigServiceClient {
	return &raftConfigServiceClient{cc}
}

func (c *raftConfigServiceClient) AddConfiguration(ctx context.Context, in *protos.NewRaftConfigRequest, opts ...grpc.CallOption) (*protos.RaftConfigResponse, error) {
	out := new(protos.RaftConfigResponse)
	err := c.cc.Invoke(ctx, "/RaftConfigService/AddConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftConfigServiceClient) GetConfiguration(ctx context.Context, in *protos.GetRaftConfigRequest, opts ...grpc.CallOption) (*protos.GetRaftConfigResponse, error) {
	out := new(protos.GetRaftConfigResponse)
	err := c.cc.Invoke(ctx, "/RaftConfigService/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftConfigServiceClient) ListConfigurations(ctx context.Context, in *protos.ListRaftConfigs, opts ...grpc.CallOption) (*protos.ListRaftConfigsResponse, error) {
	out := new(protos.ListRaftConfigsResponse)
	err := c.cc.Invoke(ctx, "/RaftConfigService/ListConfigurations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftConfigServiceServer is the server API for RaftConfigService service.
// All implementations must embed UnimplementedRaftConfigServiceServer
// for forward compatibility
type RaftConfigServiceServer interface {
	AddConfiguration(context.Context, *protos.NewRaftConfigRequest) (*protos.RaftConfigResponse, error)
	GetConfiguration(context.Context, *protos.GetRaftConfigRequest) (*protos.GetRaftConfigResponse, error)
	ListConfigurations(context.Context, *protos.ListRaftConfigs) (*protos.ListRaftConfigsResponse, error)
	mustEmbedUnimplementedRaftConfigServiceServer()
}

// UnimplementedRaftConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRaftConfigServiceServer struct {
}

func (UnimplementedRaftConfigServiceServer) AddConfiguration(context.Context, *protos.NewRaftConfigRequest) (*protos.RaftConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddConfiguration not implemented")
}
func (UnimplementedRaftConfigServiceServer) GetConfiguration(context.Context, *protos.GetRaftConfigRequest) (*protos.GetRaftConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedRaftConfigServiceServer) ListConfigurations(context.Context, *protos.ListRaftConfigs) (*protos.ListRaftConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfigurations not implemented")
}
func (UnimplementedRaftConfigServiceServer) mustEmbedUnimplementedRaftConfigServiceServer() {}

// UnsafeRaftConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftConfigServiceServer will
// result in compilation errors.
type UnsafeRaftConfigServiceServer interface {
	mustEmbedUnimplementedRaftConfigServiceServer()
}

func RegisterRaftConfigServiceServer(s grpc.ServiceRegistrar, srv RaftConfigServiceServer) {
	s.RegisterService(&RaftConfigService_ServiceDesc, srv)
}

func _RaftConfigService_AddConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.NewRaftConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftConfigServiceServer).AddConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RaftConfigService/AddConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftConfigServiceServer).AddConfiguration(ctx, req.(*protos.NewRaftConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftConfigService_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.GetRaftConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftConfigServiceServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RaftConfigService/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftConfigServiceServer).GetConfiguration(ctx, req.(*protos.GetRaftConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftConfigService_ListConfigurations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.ListRaftConfigs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftConfigServiceServer).ListConfigurations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RaftConfigService/ListConfigurations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftConfigServiceServer).ListConfigurations(ctx, req.(*protos.ListRaftConfigs))
	}
	return interceptor(ctx, in, info, handler)
}

// RaftConfigService_ServiceDesc is the grpc.ServiceDesc for RaftConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaftConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RaftConfigService",
	HandlerType: (*RaftConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddConfiguration",
			Handler:    _RaftConfigService_AddConfiguration_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _RaftConfigService_GetConfiguration_Handler,
		},
		{
			MethodName: "ListConfigurations",
			Handler:    _RaftConfigService_ListConfigurations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raft_server.proto",
}

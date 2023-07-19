// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: raftpb/service.proto

package raftpbconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	raftpb "github.com/mxplusb/pleiades/pkg/raftpb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion1_7_0

const (
	// ShardServiceName is the fully-qualified name of the ShardService service.
	ShardServiceName = "raftpb.ShardService"
	// HostServiceName is the fully-qualified name of the HostService service.
	HostServiceName = "raftpb.HostService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ShardServiceAddReplicaProcedure is the fully-qualified name of the ShardService's AddReplica RPC.
	ShardServiceAddReplicaProcedure = "/raftpb.ShardService/AddReplica"
	// ShardServiceAddReplicaObserverProcedure is the fully-qualified name of the ShardService's
	// AddReplicaObserver RPC.
	ShardServiceAddReplicaObserverProcedure = "/raftpb.ShardService/AddReplicaObserver"
	// ShardServiceAddReplicaWitnessProcedure is the fully-qualified name of the ShardService's
	// AddReplicaWitness RPC.
	ShardServiceAddReplicaWitnessProcedure = "/raftpb.ShardService/AddReplicaWitness"
	// ShardServiceGetLeaderIdProcedure is the fully-qualified name of the ShardService's GetLeaderId
	// RPC.
	ShardServiceGetLeaderIdProcedure = "/raftpb.ShardService/GetLeaderId"
	// ShardServiceGetShardMembersProcedure is the fully-qualified name of the ShardService's
	// GetShardMembers RPC.
	ShardServiceGetShardMembersProcedure = "/raftpb.ShardService/GetShardMembers"
	// ShardServiceNewShardProcedure is the fully-qualified name of the ShardService's NewShard RPC.
	ShardServiceNewShardProcedure = "/raftpb.ShardService/NewShard"
	// ShardServiceRemoveDataProcedure is the fully-qualified name of the ShardService's RemoveData RPC.
	ShardServiceRemoveDataProcedure = "/raftpb.ShardService/RemoveData"
	// ShardServiceRemoveReplicaProcedure is the fully-qualified name of the ShardService's
	// RemoveReplica RPC.
	ShardServiceRemoveReplicaProcedure = "/raftpb.ShardService/RemoveReplica"
	// ShardServiceStartReplicaProcedure is the fully-qualified name of the ShardService's StartReplica
	// RPC.
	ShardServiceStartReplicaProcedure = "/raftpb.ShardService/StartReplica"
	// ShardServiceStartReplicaObserverProcedure is the fully-qualified name of the ShardService's
	// StartReplicaObserver RPC.
	ShardServiceStartReplicaObserverProcedure = "/raftpb.ShardService/StartReplicaObserver"
	// ShardServiceStopReplicaProcedure is the fully-qualified name of the ShardService's StopReplica
	// RPC.
	ShardServiceStopReplicaProcedure = "/raftpb.ShardService/StopReplica"
	// HostServiceCompactProcedure is the fully-qualified name of the HostService's Compact RPC.
	HostServiceCompactProcedure = "/raftpb.HostService/Compact"
	// HostServiceGetHostConfigProcedure is the fully-qualified name of the HostService's GetHostConfig
	// RPC.
	HostServiceGetHostConfigProcedure = "/raftpb.HostService/GetHostConfig"
	// HostServiceSnapshotProcedure is the fully-qualified name of the HostService's Snapshot RPC.
	HostServiceSnapshotProcedure = "/raftpb.HostService/Snapshot"
	// HostServiceStopProcedure is the fully-qualified name of the HostService's Stop RPC.
	HostServiceStopProcedure = "/raftpb.HostService/Stop"
)

// ShardServiceClient is a client for the raftpb.ShardService service.
type ShardServiceClient interface {
	AddReplica(context.Context, *connect_go.Request[raftpb.AddReplicaRequest]) (*connect_go.Response[raftpb.AddReplicaResponse], error)
	AddReplicaObserver(context.Context, *connect_go.Request[raftpb.AddReplicaObserverRequest]) (*connect_go.Response[raftpb.AddReplicaObserverResponse], error)
	AddReplicaWitness(context.Context, *connect_go.Request[raftpb.AddReplicaWitnessRequest]) (*connect_go.Response[raftpb.AddReplicaWitnessResponse], error)
	GetLeaderId(context.Context, *connect_go.Request[raftpb.GetLeaderIdRequest]) (*connect_go.Response[raftpb.GetLeaderIdResponse], error)
	GetShardMembers(context.Context, *connect_go.Request[raftpb.GetShardMembersRequest]) (*connect_go.Response[raftpb.GetShardMembersResponse], error)
	NewShard(context.Context, *connect_go.Request[raftpb.NewShardRequest]) (*connect_go.Response[raftpb.NewShardResponse], error)
	RemoveData(context.Context, *connect_go.Request[raftpb.RemoveDataRequest]) (*connect_go.Response[raftpb.RemoveDataResponse], error)
	RemoveReplica(context.Context, *connect_go.Request[raftpb.RemoveReplicaRequest]) (*connect_go.Response[raftpb.RemoveReplicaResponse], error)
	StartReplica(context.Context, *connect_go.Request[raftpb.StartReplicaRequest]) (*connect_go.Response[raftpb.StartReplicaResponse], error)
	StartReplicaObserver(context.Context, *connect_go.Request[raftpb.StartReplicaObserverRequest]) (*connect_go.Response[raftpb.StartReplicaObserverResponse], error)
	StopReplica(context.Context, *connect_go.Request[raftpb.StopReplicaRequest]) (*connect_go.Response[raftpb.StopReplicaResponse], error)
}

// NewShardServiceClient constructs a client for the raftpb.ShardService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewShardServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ShardServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &shardServiceClient{
		addReplica: connect_go.NewClient[raftpb.AddReplicaRequest, raftpb.AddReplicaResponse](
			httpClient,
			baseURL+ShardServiceAddReplicaProcedure,
			opts...,
		),
		addReplicaObserver: connect_go.NewClient[raftpb.AddReplicaObserverRequest, raftpb.AddReplicaObserverResponse](
			httpClient,
			baseURL+ShardServiceAddReplicaObserverProcedure,
			opts...,
		),
		addReplicaWitness: connect_go.NewClient[raftpb.AddReplicaWitnessRequest, raftpb.AddReplicaWitnessResponse](
			httpClient,
			baseURL+ShardServiceAddReplicaWitnessProcedure,
			opts...,
		),
		getLeaderId: connect_go.NewClient[raftpb.GetLeaderIdRequest, raftpb.GetLeaderIdResponse](
			httpClient,
			baseURL+ShardServiceGetLeaderIdProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		getShardMembers: connect_go.NewClient[raftpb.GetShardMembersRequest, raftpb.GetShardMembersResponse](
			httpClient,
			baseURL+ShardServiceGetShardMembersProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		newShard: connect_go.NewClient[raftpb.NewShardRequest, raftpb.NewShardResponse](
			httpClient,
			baseURL+ShardServiceNewShardProcedure,
			opts...,
		),
		removeData: connect_go.NewClient[raftpb.RemoveDataRequest, raftpb.RemoveDataResponse](
			httpClient,
			baseURL+ShardServiceRemoveDataProcedure,
			opts...,
		),
		removeReplica: connect_go.NewClient[raftpb.RemoveReplicaRequest, raftpb.RemoveReplicaResponse](
			httpClient,
			baseURL+ShardServiceRemoveReplicaProcedure,
			opts...,
		),
		startReplica: connect_go.NewClient[raftpb.StartReplicaRequest, raftpb.StartReplicaResponse](
			httpClient,
			baseURL+ShardServiceStartReplicaProcedure,
			opts...,
		),
		startReplicaObserver: connect_go.NewClient[raftpb.StartReplicaObserverRequest, raftpb.StartReplicaObserverResponse](
			httpClient,
			baseURL+ShardServiceStartReplicaObserverProcedure,
			opts...,
		),
		stopReplica: connect_go.NewClient[raftpb.StopReplicaRequest, raftpb.StopReplicaResponse](
			httpClient,
			baseURL+ShardServiceStopReplicaProcedure,
			opts...,
		),
	}
}

// shardServiceClient implements ShardServiceClient.
type shardServiceClient struct {
	addReplica           *connect_go.Client[raftpb.AddReplicaRequest, raftpb.AddReplicaResponse]
	addReplicaObserver   *connect_go.Client[raftpb.AddReplicaObserverRequest, raftpb.AddReplicaObserverResponse]
	addReplicaWitness    *connect_go.Client[raftpb.AddReplicaWitnessRequest, raftpb.AddReplicaWitnessResponse]
	getLeaderId          *connect_go.Client[raftpb.GetLeaderIdRequest, raftpb.GetLeaderIdResponse]
	getShardMembers      *connect_go.Client[raftpb.GetShardMembersRequest, raftpb.GetShardMembersResponse]
	newShard             *connect_go.Client[raftpb.NewShardRequest, raftpb.NewShardResponse]
	removeData           *connect_go.Client[raftpb.RemoveDataRequest, raftpb.RemoveDataResponse]
	removeReplica        *connect_go.Client[raftpb.RemoveReplicaRequest, raftpb.RemoveReplicaResponse]
	startReplica         *connect_go.Client[raftpb.StartReplicaRequest, raftpb.StartReplicaResponse]
	startReplicaObserver *connect_go.Client[raftpb.StartReplicaObserverRequest, raftpb.StartReplicaObserverResponse]
	stopReplica          *connect_go.Client[raftpb.StopReplicaRequest, raftpb.StopReplicaResponse]
}

// AddReplica calls raftpb.ShardService.AddReplica.
func (c *shardServiceClient) AddReplica(ctx context.Context, req *connect_go.Request[raftpb.AddReplicaRequest]) (*connect_go.Response[raftpb.AddReplicaResponse], error) {
	return c.addReplica.CallUnary(ctx, req)
}

// AddReplicaObserver calls raftpb.ShardService.AddReplicaObserver.
func (c *shardServiceClient) AddReplicaObserver(ctx context.Context, req *connect_go.Request[raftpb.AddReplicaObserverRequest]) (*connect_go.Response[raftpb.AddReplicaObserverResponse], error) {
	return c.addReplicaObserver.CallUnary(ctx, req)
}

// AddReplicaWitness calls raftpb.ShardService.AddReplicaWitness.
func (c *shardServiceClient) AddReplicaWitness(ctx context.Context, req *connect_go.Request[raftpb.AddReplicaWitnessRequest]) (*connect_go.Response[raftpb.AddReplicaWitnessResponse], error) {
	return c.addReplicaWitness.CallUnary(ctx, req)
}

// GetLeaderId calls raftpb.ShardService.GetLeaderId.
func (c *shardServiceClient) GetLeaderId(ctx context.Context, req *connect_go.Request[raftpb.GetLeaderIdRequest]) (*connect_go.Response[raftpb.GetLeaderIdResponse], error) {
	return c.getLeaderId.CallUnary(ctx, req)
}

// GetShardMembers calls raftpb.ShardService.GetShardMembers.
func (c *shardServiceClient) GetShardMembers(ctx context.Context, req *connect_go.Request[raftpb.GetShardMembersRequest]) (*connect_go.Response[raftpb.GetShardMembersResponse], error) {
	return c.getShardMembers.CallUnary(ctx, req)
}

// NewShard calls raftpb.ShardService.NewShard.
func (c *shardServiceClient) NewShard(ctx context.Context, req *connect_go.Request[raftpb.NewShardRequest]) (*connect_go.Response[raftpb.NewShardResponse], error) {
	return c.newShard.CallUnary(ctx, req)
}

// RemoveData calls raftpb.ShardService.RemoveData.
func (c *shardServiceClient) RemoveData(ctx context.Context, req *connect_go.Request[raftpb.RemoveDataRequest]) (*connect_go.Response[raftpb.RemoveDataResponse], error) {
	return c.removeData.CallUnary(ctx, req)
}

// RemoveReplica calls raftpb.ShardService.RemoveReplica.
func (c *shardServiceClient) RemoveReplica(ctx context.Context, req *connect_go.Request[raftpb.RemoveReplicaRequest]) (*connect_go.Response[raftpb.RemoveReplicaResponse], error) {
	return c.removeReplica.CallUnary(ctx, req)
}

// StartReplica calls raftpb.ShardService.StartReplica.
func (c *shardServiceClient) StartReplica(ctx context.Context, req *connect_go.Request[raftpb.StartReplicaRequest]) (*connect_go.Response[raftpb.StartReplicaResponse], error) {
	return c.startReplica.CallUnary(ctx, req)
}

// StartReplicaObserver calls raftpb.ShardService.StartReplicaObserver.
func (c *shardServiceClient) StartReplicaObserver(ctx context.Context, req *connect_go.Request[raftpb.StartReplicaObserverRequest]) (*connect_go.Response[raftpb.StartReplicaObserverResponse], error) {
	return c.startReplicaObserver.CallUnary(ctx, req)
}

// StopReplica calls raftpb.ShardService.StopReplica.
func (c *shardServiceClient) StopReplica(ctx context.Context, req *connect_go.Request[raftpb.StopReplicaRequest]) (*connect_go.Response[raftpb.StopReplicaResponse], error) {
	return c.stopReplica.CallUnary(ctx, req)
}

// ShardServiceHandler is an implementation of the raftpb.ShardService service.
type ShardServiceHandler interface {
	AddReplica(context.Context, *connect_go.Request[raftpb.AddReplicaRequest]) (*connect_go.Response[raftpb.AddReplicaResponse], error)
	AddReplicaObserver(context.Context, *connect_go.Request[raftpb.AddReplicaObserverRequest]) (*connect_go.Response[raftpb.AddReplicaObserverResponse], error)
	AddReplicaWitness(context.Context, *connect_go.Request[raftpb.AddReplicaWitnessRequest]) (*connect_go.Response[raftpb.AddReplicaWitnessResponse], error)
	GetLeaderId(context.Context, *connect_go.Request[raftpb.GetLeaderIdRequest]) (*connect_go.Response[raftpb.GetLeaderIdResponse], error)
	GetShardMembers(context.Context, *connect_go.Request[raftpb.GetShardMembersRequest]) (*connect_go.Response[raftpb.GetShardMembersResponse], error)
	NewShard(context.Context, *connect_go.Request[raftpb.NewShardRequest]) (*connect_go.Response[raftpb.NewShardResponse], error)
	RemoveData(context.Context, *connect_go.Request[raftpb.RemoveDataRequest]) (*connect_go.Response[raftpb.RemoveDataResponse], error)
	RemoveReplica(context.Context, *connect_go.Request[raftpb.RemoveReplicaRequest]) (*connect_go.Response[raftpb.RemoveReplicaResponse], error)
	StartReplica(context.Context, *connect_go.Request[raftpb.StartReplicaRequest]) (*connect_go.Response[raftpb.StartReplicaResponse], error)
	StartReplicaObserver(context.Context, *connect_go.Request[raftpb.StartReplicaObserverRequest]) (*connect_go.Response[raftpb.StartReplicaObserverResponse], error)
	StopReplica(context.Context, *connect_go.Request[raftpb.StopReplicaRequest]) (*connect_go.Response[raftpb.StopReplicaResponse], error)
}

// NewShardServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewShardServiceHandler(svc ShardServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	shardServiceAddReplicaHandler := connect_go.NewUnaryHandler(
		ShardServiceAddReplicaProcedure,
		svc.AddReplica,
		opts...,
	)
	shardServiceAddReplicaObserverHandler := connect_go.NewUnaryHandler(
		ShardServiceAddReplicaObserverProcedure,
		svc.AddReplicaObserver,
		opts...,
	)
	shardServiceAddReplicaWitnessHandler := connect_go.NewUnaryHandler(
		ShardServiceAddReplicaWitnessProcedure,
		svc.AddReplicaWitness,
		opts...,
	)
	shardServiceGetLeaderIdHandler := connect_go.NewUnaryHandler(
		ShardServiceGetLeaderIdProcedure,
		svc.GetLeaderId,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	shardServiceGetShardMembersHandler := connect_go.NewUnaryHandler(
		ShardServiceGetShardMembersProcedure,
		svc.GetShardMembers,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	shardServiceNewShardHandler := connect_go.NewUnaryHandler(
		ShardServiceNewShardProcedure,
		svc.NewShard,
		opts...,
	)
	shardServiceRemoveDataHandler := connect_go.NewUnaryHandler(
		ShardServiceRemoveDataProcedure,
		svc.RemoveData,
		opts...,
	)
	shardServiceRemoveReplicaHandler := connect_go.NewUnaryHandler(
		ShardServiceRemoveReplicaProcedure,
		svc.RemoveReplica,
		opts...,
	)
	shardServiceStartReplicaHandler := connect_go.NewUnaryHandler(
		ShardServiceStartReplicaProcedure,
		svc.StartReplica,
		opts...,
	)
	shardServiceStartReplicaObserverHandler := connect_go.NewUnaryHandler(
		ShardServiceStartReplicaObserverProcedure,
		svc.StartReplicaObserver,
		opts...,
	)
	shardServiceStopReplicaHandler := connect_go.NewUnaryHandler(
		ShardServiceStopReplicaProcedure,
		svc.StopReplica,
		opts...,
	)
	return "/raftpb.ShardService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ShardServiceAddReplicaProcedure:
			shardServiceAddReplicaHandler.ServeHTTP(w, r)
		case ShardServiceAddReplicaObserverProcedure:
			shardServiceAddReplicaObserverHandler.ServeHTTP(w, r)
		case ShardServiceAddReplicaWitnessProcedure:
			shardServiceAddReplicaWitnessHandler.ServeHTTP(w, r)
		case ShardServiceGetLeaderIdProcedure:
			shardServiceGetLeaderIdHandler.ServeHTTP(w, r)
		case ShardServiceGetShardMembersProcedure:
			shardServiceGetShardMembersHandler.ServeHTTP(w, r)
		case ShardServiceNewShardProcedure:
			shardServiceNewShardHandler.ServeHTTP(w, r)
		case ShardServiceRemoveDataProcedure:
			shardServiceRemoveDataHandler.ServeHTTP(w, r)
		case ShardServiceRemoveReplicaProcedure:
			shardServiceRemoveReplicaHandler.ServeHTTP(w, r)
		case ShardServiceStartReplicaProcedure:
			shardServiceStartReplicaHandler.ServeHTTP(w, r)
		case ShardServiceStartReplicaObserverProcedure:
			shardServiceStartReplicaObserverHandler.ServeHTTP(w, r)
		case ShardServiceStopReplicaProcedure:
			shardServiceStopReplicaHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedShardServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedShardServiceHandler struct{}

func (UnimplementedShardServiceHandler) AddReplica(context.Context, *connect_go.Request[raftpb.AddReplicaRequest]) (*connect_go.Response[raftpb.AddReplicaResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.ShardService.AddReplica is not implemented"))
}

func (UnimplementedShardServiceHandler) AddReplicaObserver(context.Context, *connect_go.Request[raftpb.AddReplicaObserverRequest]) (*connect_go.Response[raftpb.AddReplicaObserverResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.ShardService.AddReplicaObserver is not implemented"))
}

func (UnimplementedShardServiceHandler) AddReplicaWitness(context.Context, *connect_go.Request[raftpb.AddReplicaWitnessRequest]) (*connect_go.Response[raftpb.AddReplicaWitnessResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.ShardService.AddReplicaWitness is not implemented"))
}

func (UnimplementedShardServiceHandler) GetLeaderId(context.Context, *connect_go.Request[raftpb.GetLeaderIdRequest]) (*connect_go.Response[raftpb.GetLeaderIdResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.ShardService.GetLeaderId is not implemented"))
}

func (UnimplementedShardServiceHandler) GetShardMembers(context.Context, *connect_go.Request[raftpb.GetShardMembersRequest]) (*connect_go.Response[raftpb.GetShardMembersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.ShardService.GetShardMembers is not implemented"))
}

func (UnimplementedShardServiceHandler) NewShard(context.Context, *connect_go.Request[raftpb.NewShardRequest]) (*connect_go.Response[raftpb.NewShardResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.ShardService.NewShard is not implemented"))
}

func (UnimplementedShardServiceHandler) RemoveData(context.Context, *connect_go.Request[raftpb.RemoveDataRequest]) (*connect_go.Response[raftpb.RemoveDataResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.ShardService.RemoveData is not implemented"))
}

func (UnimplementedShardServiceHandler) RemoveReplica(context.Context, *connect_go.Request[raftpb.RemoveReplicaRequest]) (*connect_go.Response[raftpb.RemoveReplicaResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.ShardService.RemoveReplica is not implemented"))
}

func (UnimplementedShardServiceHandler) StartReplica(context.Context, *connect_go.Request[raftpb.StartReplicaRequest]) (*connect_go.Response[raftpb.StartReplicaResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.ShardService.StartReplica is not implemented"))
}

func (UnimplementedShardServiceHandler) StartReplicaObserver(context.Context, *connect_go.Request[raftpb.StartReplicaObserverRequest]) (*connect_go.Response[raftpb.StartReplicaObserverResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.ShardService.StartReplicaObserver is not implemented"))
}

func (UnimplementedShardServiceHandler) StopReplica(context.Context, *connect_go.Request[raftpb.StopReplicaRequest]) (*connect_go.Response[raftpb.StopReplicaResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.ShardService.StopReplica is not implemented"))
}

// HostServiceClient is a client for the raftpb.HostService service.
type HostServiceClient interface {
	Compact(context.Context, *connect_go.Request[raftpb.CompactRequest]) (*connect_go.Response[raftpb.CompactResponse], error)
	GetHostConfig(context.Context, *connect_go.Request[raftpb.GetHostConfigRequest]) (*connect_go.Response[raftpb.GetHostConfigResponse], error)
	// rpc LeaderTransfer(LeaderTransferRequest) returns (LeaderTransferResponse);
	Snapshot(context.Context, *connect_go.Request[raftpb.SnapshotRequest]) (*connect_go.Response[raftpb.SnapshotResponse], error)
	Stop(context.Context, *connect_go.Request[raftpb.StopRequest]) (*connect_go.Response[raftpb.StopResponse], error)
}

// NewHostServiceClient constructs a client for the raftpb.HostService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHostServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) HostServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &hostServiceClient{
		compact: connect_go.NewClient[raftpb.CompactRequest, raftpb.CompactResponse](
			httpClient,
			baseURL+HostServiceCompactProcedure,
			opts...,
		),
		getHostConfig: connect_go.NewClient[raftpb.GetHostConfigRequest, raftpb.GetHostConfigResponse](
			httpClient,
			baseURL+HostServiceGetHostConfigProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		snapshot: connect_go.NewClient[raftpb.SnapshotRequest, raftpb.SnapshotResponse](
			httpClient,
			baseURL+HostServiceSnapshotProcedure,
			opts...,
		),
		stop: connect_go.NewClient[raftpb.StopRequest, raftpb.StopResponse](
			httpClient,
			baseURL+HostServiceStopProcedure,
			opts...,
		),
	}
}

// hostServiceClient implements HostServiceClient.
type hostServiceClient struct {
	compact       *connect_go.Client[raftpb.CompactRequest, raftpb.CompactResponse]
	getHostConfig *connect_go.Client[raftpb.GetHostConfigRequest, raftpb.GetHostConfigResponse]
	snapshot      *connect_go.Client[raftpb.SnapshotRequest, raftpb.SnapshotResponse]
	stop          *connect_go.Client[raftpb.StopRequest, raftpb.StopResponse]
}

// Compact calls raftpb.HostService.Compact.
func (c *hostServiceClient) Compact(ctx context.Context, req *connect_go.Request[raftpb.CompactRequest]) (*connect_go.Response[raftpb.CompactResponse], error) {
	return c.compact.CallUnary(ctx, req)
}

// GetHostConfig calls raftpb.HostService.GetHostConfig.
func (c *hostServiceClient) GetHostConfig(ctx context.Context, req *connect_go.Request[raftpb.GetHostConfigRequest]) (*connect_go.Response[raftpb.GetHostConfigResponse], error) {
	return c.getHostConfig.CallUnary(ctx, req)
}

// Snapshot calls raftpb.HostService.Snapshot.
func (c *hostServiceClient) Snapshot(ctx context.Context, req *connect_go.Request[raftpb.SnapshotRequest]) (*connect_go.Response[raftpb.SnapshotResponse], error) {
	return c.snapshot.CallUnary(ctx, req)
}

// Stop calls raftpb.HostService.Stop.
func (c *hostServiceClient) Stop(ctx context.Context, req *connect_go.Request[raftpb.StopRequest]) (*connect_go.Response[raftpb.StopResponse], error) {
	return c.stop.CallUnary(ctx, req)
}

// HostServiceHandler is an implementation of the raftpb.HostService service.
type HostServiceHandler interface {
	Compact(context.Context, *connect_go.Request[raftpb.CompactRequest]) (*connect_go.Response[raftpb.CompactResponse], error)
	GetHostConfig(context.Context, *connect_go.Request[raftpb.GetHostConfigRequest]) (*connect_go.Response[raftpb.GetHostConfigResponse], error)
	// rpc LeaderTransfer(LeaderTransferRequest) returns (LeaderTransferResponse);
	Snapshot(context.Context, *connect_go.Request[raftpb.SnapshotRequest]) (*connect_go.Response[raftpb.SnapshotResponse], error)
	Stop(context.Context, *connect_go.Request[raftpb.StopRequest]) (*connect_go.Response[raftpb.StopResponse], error)
}

// NewHostServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHostServiceHandler(svc HostServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	hostServiceCompactHandler := connect_go.NewUnaryHandler(
		HostServiceCompactProcedure,
		svc.Compact,
		opts...,
	)
	hostServiceGetHostConfigHandler := connect_go.NewUnaryHandler(
		HostServiceGetHostConfigProcedure,
		svc.GetHostConfig,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	hostServiceSnapshotHandler := connect_go.NewUnaryHandler(
		HostServiceSnapshotProcedure,
		svc.Snapshot,
		opts...,
	)
	hostServiceStopHandler := connect_go.NewUnaryHandler(
		HostServiceStopProcedure,
		svc.Stop,
		opts...,
	)
	return "/raftpb.HostService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case HostServiceCompactProcedure:
			hostServiceCompactHandler.ServeHTTP(w, r)
		case HostServiceGetHostConfigProcedure:
			hostServiceGetHostConfigHandler.ServeHTTP(w, r)
		case HostServiceSnapshotProcedure:
			hostServiceSnapshotHandler.ServeHTTP(w, r)
		case HostServiceStopProcedure:
			hostServiceStopHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedHostServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedHostServiceHandler struct{}

func (UnimplementedHostServiceHandler) Compact(context.Context, *connect_go.Request[raftpb.CompactRequest]) (*connect_go.Response[raftpb.CompactResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.HostService.Compact is not implemented"))
}

func (UnimplementedHostServiceHandler) GetHostConfig(context.Context, *connect_go.Request[raftpb.GetHostConfigRequest]) (*connect_go.Response[raftpb.GetHostConfigResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.HostService.GetHostConfig is not implemented"))
}

func (UnimplementedHostServiceHandler) Snapshot(context.Context, *connect_go.Request[raftpb.SnapshotRequest]) (*connect_go.Response[raftpb.SnapshotResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.HostService.Snapshot is not implemented"))
}

func (UnimplementedHostServiceHandler) Stop(context.Context, *connect_go.Request[raftpb.StopRequest]) (*connect_go.Response[raftpb.StopResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("raftpb.HostService.Stop is not implemented"))
}

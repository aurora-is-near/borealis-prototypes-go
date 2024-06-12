// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: blocksapi/services.proto

package pb_blocksapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BlocksProviderClient is the client API for BlocksProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlocksProviderClient interface {
	// Returns information about available block streams
	ListBlockStreams(ctx context.Context, in *ListBlockStreamsRequest, opts ...grpc.CallOption) (*ListBlockStreamsResponse, error)
	// Reads individual message
	GetBlockMessage(ctx context.Context, in *GetBlockMessageRequest, opts ...grpc.CallOption) (*GetBlockMessageResponse, error)
	// Streams blocks continiously
	ReceiveBlocks(ctx context.Context, in *ReceiveBlocksRequest, opts ...grpc.CallOption) (BlocksProvider_ReceiveBlocksClient, error)
}

type blocksProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewBlocksProviderClient(cc grpc.ClientConnInterface) BlocksProviderClient {
	return &blocksProviderClient{cc}
}

func (c *blocksProviderClient) ListBlockStreams(ctx context.Context, in *ListBlockStreamsRequest, opts ...grpc.CallOption) (*ListBlockStreamsResponse, error) {
	out := new(ListBlockStreamsResponse)
	err := c.cc.Invoke(ctx, "/borealis.blocksapi.BlocksProvider/ListBlockStreams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksProviderClient) GetBlockMessage(ctx context.Context, in *GetBlockMessageRequest, opts ...grpc.CallOption) (*GetBlockMessageResponse, error) {
	out := new(GetBlockMessageResponse)
	err := c.cc.Invoke(ctx, "/borealis.blocksapi.BlocksProvider/GetBlockMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksProviderClient) ReceiveBlocks(ctx context.Context, in *ReceiveBlocksRequest, opts ...grpc.CallOption) (BlocksProvider_ReceiveBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlocksProvider_ServiceDesc.Streams[0], "/borealis.blocksapi.BlocksProvider/ReceiveBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &blocksProviderReceiveBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlocksProvider_ReceiveBlocksClient interface {
	Recv() (*ReceiveBlocksResponse, error)
	grpc.ClientStream
}

type blocksProviderReceiveBlocksClient struct {
	grpc.ClientStream
}

func (x *blocksProviderReceiveBlocksClient) Recv() (*ReceiveBlocksResponse, error) {
	m := new(ReceiveBlocksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlocksProviderServer is the server API for BlocksProvider service.
// All implementations must embed UnimplementedBlocksProviderServer
// for forward compatibility
type BlocksProviderServer interface {
	// Returns information about available block streams
	ListBlockStreams(context.Context, *ListBlockStreamsRequest) (*ListBlockStreamsResponse, error)
	// Reads individual message
	GetBlockMessage(context.Context, *GetBlockMessageRequest) (*GetBlockMessageResponse, error)
	// Streams blocks continiously
	ReceiveBlocks(*ReceiveBlocksRequest, BlocksProvider_ReceiveBlocksServer) error
	mustEmbedUnimplementedBlocksProviderServer()
}

// UnimplementedBlocksProviderServer must be embedded to have forward compatible implementations.
type UnimplementedBlocksProviderServer struct {
}

func (UnimplementedBlocksProviderServer) ListBlockStreams(context.Context, *ListBlockStreamsRequest) (*ListBlockStreamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBlockStreams not implemented")
}
func (UnimplementedBlocksProviderServer) GetBlockMessage(context.Context, *GetBlockMessageRequest) (*GetBlockMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockMessage not implemented")
}
func (UnimplementedBlocksProviderServer) ReceiveBlocks(*ReceiveBlocksRequest, BlocksProvider_ReceiveBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveBlocks not implemented")
}
func (UnimplementedBlocksProviderServer) mustEmbedUnimplementedBlocksProviderServer() {}

// UnsafeBlocksProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlocksProviderServer will
// result in compilation errors.
type UnsafeBlocksProviderServer interface {
	mustEmbedUnimplementedBlocksProviderServer()
}

func RegisterBlocksProviderServer(s grpc.ServiceRegistrar, srv BlocksProviderServer) {
	s.RegisterService(&BlocksProvider_ServiceDesc, srv)
}

func _BlocksProvider_ListBlockStreams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBlockStreamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksProviderServer).ListBlockStreams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/borealis.blocksapi.BlocksProvider/ListBlockStreams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksProviderServer).ListBlockStreams(ctx, req.(*ListBlockStreamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksProvider_GetBlockMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksProviderServer).GetBlockMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/borealis.blocksapi.BlocksProvider/GetBlockMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksProviderServer).GetBlockMessage(ctx, req.(*GetBlockMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksProvider_ReceiveBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveBlocksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlocksProviderServer).ReceiveBlocks(m, &blocksProviderReceiveBlocksServer{stream})
}

type BlocksProvider_ReceiveBlocksServer interface {
	Send(*ReceiveBlocksResponse) error
	grpc.ServerStream
}

type blocksProviderReceiveBlocksServer struct {
	grpc.ServerStream
}

func (x *blocksProviderReceiveBlocksServer) Send(m *ReceiveBlocksResponse) error {
	return x.ServerStream.SendMsg(m)
}

// BlocksProvider_ServiceDesc is the grpc.ServiceDesc for BlocksProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlocksProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "borealis.blocksapi.BlocksProvider",
	HandlerType: (*BlocksProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBlockStreams",
			Handler:    _BlocksProvider_ListBlockStreams_Handler,
		},
		{
			MethodName: "GetBlockMessage",
			Handler:    _BlocksProvider_GetBlockMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveBlocks",
			Handler:       _BlocksProvider_ReceiveBlocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blocksapi/services.proto",
}

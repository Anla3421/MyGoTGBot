// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BotGetOrderClient is the client API for BotGetOrder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotGetOrderClient interface {
	BotOrderQuery(ctx context.Context, in *BotOrderQueryRequest, opts ...grpc.CallOption) (*BotOrderQueryResponse, error)
	BotOrderCreate(ctx context.Context, in *BotOrderCreateRequest, opts ...grpc.CallOption) (*BotOrderCreateResponse, error)
	BotOrderUpdate(ctx context.Context, in *BotOrderUpdateRequest, opts ...grpc.CallOption) (*BotOrderUpdateResponse, error)
}

type botGetOrderClient struct {
	cc grpc.ClientConnInterface
}

func NewBotGetOrderClient(cc grpc.ClientConnInterface) BotGetOrderClient {
	return &botGetOrderClient{cc}
}

func (c *botGetOrderClient) BotOrderQuery(ctx context.Context, in *BotOrderQueryRequest, opts ...grpc.CallOption) (*BotOrderQueryResponse, error) {
	out := new(BotOrderQueryResponse)
	err := c.cc.Invoke(ctx, "/protobuf.BotGetOrder/BotOrderQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botGetOrderClient) BotOrderCreate(ctx context.Context, in *BotOrderCreateRequest, opts ...grpc.CallOption) (*BotOrderCreateResponse, error) {
	out := new(BotOrderCreateResponse)
	err := c.cc.Invoke(ctx, "/protobuf.BotGetOrder/BotOrderCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botGetOrderClient) BotOrderUpdate(ctx context.Context, in *BotOrderUpdateRequest, opts ...grpc.CallOption) (*BotOrderUpdateResponse, error) {
	out := new(BotOrderUpdateResponse)
	err := c.cc.Invoke(ctx, "/protobuf.BotGetOrder/BotOrderUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotGetOrderServer is the server API for BotGetOrder service.
// All implementations must embed UnimplementedBotGetOrderServer
// for forward compatibility
type BotGetOrderServer interface {
	BotOrderQuery(context.Context, *BotOrderQueryRequest) (*BotOrderQueryResponse, error)
	BotOrderCreate(context.Context, *BotOrderCreateRequest) (*BotOrderCreateResponse, error)
	BotOrderUpdate(context.Context, *BotOrderUpdateRequest) (*BotOrderUpdateResponse, error)
	mustEmbedUnimplementedBotGetOrderServer()
}

// UnimplementedBotGetOrderServer must be embedded to have forward compatible implementations.
type UnimplementedBotGetOrderServer struct {
}

func (UnimplementedBotGetOrderServer) BotOrderQuery(context.Context, *BotOrderQueryRequest) (*BotOrderQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BotOrderQuery not implemented")
}
func (UnimplementedBotGetOrderServer) BotOrderCreate(context.Context, *BotOrderCreateRequest) (*BotOrderCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BotOrderCreate not implemented")
}
func (UnimplementedBotGetOrderServer) BotOrderUpdate(context.Context, *BotOrderUpdateRequest) (*BotOrderUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BotOrderUpdate not implemented")
}
func (UnimplementedBotGetOrderServer) mustEmbedUnimplementedBotGetOrderServer() {}

// UnsafeBotGetOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotGetOrderServer will
// result in compilation errors.
type UnsafeBotGetOrderServer interface {
	mustEmbedUnimplementedBotGetOrderServer()
}

func RegisterBotGetOrderServer(s grpc.ServiceRegistrar, srv BotGetOrderServer) {
	s.RegisterService(&BotGetOrder_ServiceDesc, srv)
}

func _BotGetOrder_BotOrderQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotOrderQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotGetOrderServer).BotOrderQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BotGetOrder/BotOrderQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotGetOrderServer).BotOrderQuery(ctx, req.(*BotOrderQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotGetOrder_BotOrderCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotOrderCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotGetOrderServer).BotOrderCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BotGetOrder/BotOrderCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotGetOrderServer).BotOrderCreate(ctx, req.(*BotOrderCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotGetOrder_BotOrderUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotOrderUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotGetOrderServer).BotOrderUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.BotGetOrder/BotOrderUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotGetOrderServer).BotOrderUpdate(ctx, req.(*BotOrderUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BotGetOrder_ServiceDesc is the grpc.ServiceDesc for BotGetOrder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BotGetOrder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.BotGetOrder",
	HandlerType: (*BotGetOrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BotOrderQuery",
			Handler:    _BotGetOrder_BotOrderQuery_Handler,
		},
		{
			MethodName: "BotOrderCreate",
			Handler:    _BotGetOrder_BotOrderCreate_Handler,
		},
		{
			MethodName: "BotOrderUpdate",
			Handler:    _BotGetOrder_BotOrderUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/botGetOrder.proto",
}

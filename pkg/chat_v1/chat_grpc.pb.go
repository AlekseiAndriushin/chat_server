// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: chat.proto

package chat_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// Chat_V1Client is the client API for Chat_V1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Chat_V1Client interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendMessage(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type chat_V1Client struct {
	cc grpc.ClientConnInterface
}

func NewChat_V1Client(cc grpc.ClientConnInterface) Chat_V1Client {
	return &chat_V1Client{cc}
}

func (c *chat_V1Client) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/chat_v1.Chat_V1/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chat_V1Client) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chat_v1.Chat_V1/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chat_V1Client) SendMessage(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chat_v1.Chat_V1/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Chat_V1Server is the server API for Chat_V1 service.
// All implementations must embed UnimplementedChat_V1Server
// for forward compatibility
type Chat_V1Server interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	SendMessage(context.Context, *SendRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedChat_V1Server()
}

// UnimplementedChat_V1Server must be embedded to have forward compatible implementations.
type UnimplementedChat_V1Server struct {
}

func (UnimplementedChat_V1Server) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedChat_V1Server) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedChat_V1Server) SendMessage(context.Context, *SendRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChat_V1Server) mustEmbedUnimplementedChat_V1Server() {}

// UnsafeChat_V1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Chat_V1Server will
// result in compilation errors.
type UnsafeChat_V1Server interface {
	mustEmbedUnimplementedChat_V1Server()
}

func RegisterChat_V1Server(s grpc.ServiceRegistrar, srv Chat_V1Server) {
	s.RegisterService(&Chat_V1_ServiceDesc, srv)
}

func _Chat_V1_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Chat_V1Server).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.Chat_V1/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Chat_V1Server).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_V1_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Chat_V1Server).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.Chat_V1/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Chat_V1Server).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_V1_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Chat_V1Server).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.Chat_V1/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Chat_V1Server).SendMessage(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_V1_ServiceDesc is the grpc.ServiceDesc for Chat_V1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_V1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat_v1.Chat_V1",
	HandlerType: (*Chat_V1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Chat_V1_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Chat_V1_Delete_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Chat_V1_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
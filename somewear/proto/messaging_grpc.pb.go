// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: messaging.proto

package __

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

const (
	Messaging_SendMessage_FullMethodName    = "/messaging.Messaging/SendMessage"
	Messaging_StreamMessages_FullMethodName = "/messaging.Messaging/StreamMessages"
)

// MessagingClient is the client API for Messaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingClient interface {
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	StreamMessages(ctx context.Context, opts ...grpc.CallOption) (Messaging_StreamMessagesClient, error)
}

type messagingClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingClient(cc grpc.ClientConnInterface) MessagingClient {
	return &messagingClient{cc}
}

func (c *messagingClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, Messaging_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) StreamMessages(ctx context.Context, opts ...grpc.CallOption) (Messaging_StreamMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messaging_ServiceDesc.Streams[0], Messaging_StreamMessages_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messagingStreamMessagesClient{stream}
	return x, nil
}

type Messaging_StreamMessagesClient interface {
	Send(*StreamMessagesRequest) error
	Recv() (*StreamMessagesResponse, error)
	grpc.ClientStream
}

type messagingStreamMessagesClient struct {
	grpc.ClientStream
}

func (x *messagingStreamMessagesClient) Send(m *StreamMessagesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messagingStreamMessagesClient) Recv() (*StreamMessagesResponse, error) {
	m := new(StreamMessagesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessagingServer is the server API for Messaging service.
// All implementations must embed UnimplementedMessagingServer
// for forward compatibility
type MessagingServer interface {
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	StreamMessages(Messaging_StreamMessagesServer) error
	mustEmbedUnimplementedMessagingServer()
}

// UnimplementedMessagingServer must be embedded to have forward compatible implementations.
type UnimplementedMessagingServer struct {
}

func (UnimplementedMessagingServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessagingServer) StreamMessages(Messaging_StreamMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMessages not implemented")
}
func (UnimplementedMessagingServer) mustEmbedUnimplementedMessagingServer() {}

// UnsafeMessagingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServer will
// result in compilation errors.
type UnsafeMessagingServer interface {
	mustEmbedUnimplementedMessagingServer()
}

func RegisterMessagingServer(s grpc.ServiceRegistrar, srv MessagingServer) {
	s.RegisterService(&Messaging_ServiceDesc, srv)
}

func _Messaging_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_StreamMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessagingServer).StreamMessages(&messagingStreamMessagesServer{stream})
}

type Messaging_StreamMessagesServer interface {
	Send(*StreamMessagesResponse) error
	Recv() (*StreamMessagesRequest, error)
	grpc.ServerStream
}

type messagingStreamMessagesServer struct {
	grpc.ServerStream
}

func (x *messagingStreamMessagesServer) Send(m *StreamMessagesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messagingStreamMessagesServer) Recv() (*StreamMessagesRequest, error) {
	m := new(StreamMessagesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Messaging_ServiceDesc is the grpc.ServiceDesc for Messaging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messaging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messaging.Messaging",
	HandlerType: (*MessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Messaging_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMessages",
			Handler:       _Messaging_StreamMessages_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "messaging.proto",
}

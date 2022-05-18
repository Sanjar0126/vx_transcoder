// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package publisher_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PublisherServiceClient is the client API for PublisherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublisherServiceClient interface {
	SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	PublishNotification(ctx context.Context, in *NotificationBody, opts ...grpc.CallOption) (*empty.Empty, error)
	HealthCheck(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateMultipartUpload(ctx context.Context, in *MultipartUpload, opts ...grpc.CallOption) (*MultipartUpload, error)
}

type publisherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublisherServiceClient(cc grpc.ClientConnInterface) PublisherServiceClient {
	return &publisherServiceClient{cc}
}

func (c *publisherServiceClient) SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tariff.PublisherService/SendSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherServiceClient) PublishNotification(ctx context.Context, in *NotificationBody, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tariff.PublisherService/PublishNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherServiceClient) HealthCheck(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tariff.PublisherService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherServiceClient) CreateMultipartUpload(ctx context.Context, in *MultipartUpload, opts ...grpc.CallOption) (*MultipartUpload, error) {
	out := new(MultipartUpload)
	err := c.cc.Invoke(ctx, "/tariff.PublisherService/CreateMultipartUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublisherServiceServer is the server API for PublisherService service.
// All implementations must embed UnimplementedPublisherServiceServer
// for forward compatibility
type PublisherServiceServer interface {
	SendSms(context.Context, *SendSmsRequest) (*empty.Empty, error)
	PublishNotification(context.Context, *NotificationBody) (*empty.Empty, error)
	HealthCheck(context.Context, *empty.Empty) (*empty.Empty, error)
	CreateMultipartUpload(context.Context, *MultipartUpload) (*MultipartUpload, error)
	mustEmbedUnimplementedPublisherServiceServer()
}

// UnimplementedPublisherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublisherServiceServer struct {
}

func (UnimplementedPublisherServiceServer) SendSms(context.Context, *SendSmsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}
func (UnimplementedPublisherServiceServer) PublishNotification(context.Context, *NotificationBody) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishNotification not implemented")
}
func (UnimplementedPublisherServiceServer) HealthCheck(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedPublisherServiceServer) CreateMultipartUpload(context.Context, *MultipartUpload) (*MultipartUpload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMultipartUpload not implemented")
}
func (UnimplementedPublisherServiceServer) mustEmbedUnimplementedPublisherServiceServer() {}

// UnsafePublisherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublisherServiceServer will
// result in compilation errors.
type UnsafePublisherServiceServer interface {
	mustEmbedUnimplementedPublisherServiceServer()
}

func RegisterPublisherServiceServer(s grpc.ServiceRegistrar, srv PublisherServiceServer) {
	s.RegisterService(&PublisherService_ServiceDesc, srv)
}

func _PublisherService_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServiceServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tariff.PublisherService/SendSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServiceServer).SendSms(ctx, req.(*SendSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublisherService_PublishNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServiceServer).PublishNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tariff.PublisherService/PublishNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServiceServer).PublishNotification(ctx, req.(*NotificationBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublisherService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tariff.PublisherService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServiceServer).HealthCheck(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublisherService_CreateMultipartUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultipartUpload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServiceServer).CreateMultipartUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tariff.PublisherService/CreateMultipartUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServiceServer).CreateMultipartUpload(ctx, req.(*MultipartUpload))
	}
	return interceptor(ctx, in, info, handler)
}

// PublisherService_ServiceDesc is the grpc.ServiceDesc for PublisherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublisherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tariff.PublisherService",
	HandlerType: (*PublisherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSms",
			Handler:    _PublisherService_SendSms_Handler,
		},
		{
			MethodName: "PublishNotification",
			Handler:    _PublisherService_PublishNotification_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _PublisherService_HealthCheck_Handler,
		},
		{
			MethodName: "CreateMultipartUpload",
			Handler:    _PublisherService_CreateMultipartUpload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publisher_service.proto",
}

// WebSocketServiceClient is the client API for WebSocketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebSocketServiceClient interface {
	SendToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type webSocketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebSocketServiceClient(cc grpc.ClientConnInterface) WebSocketServiceClient {
	return &webSocketServiceClient{cc}
}

func (c *webSocketServiceClient) SendToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tariff.WebSocketService/SendToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebSocketServiceServer is the server API for WebSocketService service.
// All implementations must embed UnimplementedWebSocketServiceServer
// for forward compatibility
type WebSocketServiceServer interface {
	SendToken(context.Context, *TokenRequest) (*empty.Empty, error)
	mustEmbedUnimplementedWebSocketServiceServer()
}

// UnimplementedWebSocketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWebSocketServiceServer struct {
}

func (UnimplementedWebSocketServiceServer) SendToken(context.Context, *TokenRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToken not implemented")
}
func (UnimplementedWebSocketServiceServer) mustEmbedUnimplementedWebSocketServiceServer() {}

// UnsafeWebSocketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebSocketServiceServer will
// result in compilation errors.
type UnsafeWebSocketServiceServer interface {
	mustEmbedUnimplementedWebSocketServiceServer()
}

func RegisterWebSocketServiceServer(s grpc.ServiceRegistrar, srv WebSocketServiceServer) {
	s.RegisterService(&WebSocketService_ServiceDesc, srv)
}

func _WebSocketService_SendToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebSocketServiceServer).SendToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tariff.WebSocketService/SendToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebSocketServiceServer).SendToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WebSocketService_ServiceDesc is the grpc.ServiceDesc for WebSocketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebSocketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tariff.WebSocketService",
	HandlerType: (*WebSocketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendToken",
			Handler:    _WebSocketService_SendToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publisher_service.proto",
}

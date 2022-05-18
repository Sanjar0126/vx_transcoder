// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package content_service

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

// FeedbackServiceClient is the client API for FeedbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedbackServiceClient interface {
	Get(ctx context.Context, in *FeedbackGetRequest, opts ...grpc.CallOption) (*FeedbackGetResponse, error)
	GetAll(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*FeedBackGetAllResponse, error)
	Update(ctx context.Context, in *FeedbackUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Create(ctx context.Context, in *FeedbackCreateRequest, opts ...grpc.CallOption) (*FeedBackCreateResponse, error)
	Delete(ctx context.Context, in *FeedbackDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type feedbackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedbackServiceClient(cc grpc.ClientConnInterface) FeedbackServiceClient {
	return &feedbackServiceClient{cc}
}

func (c *feedbackServiceClient) Get(ctx context.Context, in *FeedbackGetRequest, opts ...grpc.CallOption) (*FeedbackGetResponse, error) {
	out := new(FeedbackGetResponse)
	err := c.cc.Invoke(ctx, "/content.FeedbackService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackServiceClient) GetAll(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*FeedBackGetAllResponse, error) {
	out := new(FeedBackGetAllResponse)
	err := c.cc.Invoke(ctx, "/content.FeedbackService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackServiceClient) Update(ctx context.Context, in *FeedbackUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.FeedbackService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackServiceClient) Create(ctx context.Context, in *FeedbackCreateRequest, opts ...grpc.CallOption) (*FeedBackCreateResponse, error) {
	out := new(FeedBackCreateResponse)
	err := c.cc.Invoke(ctx, "/content.FeedbackService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackServiceClient) Delete(ctx context.Context, in *FeedbackDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.FeedbackService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedbackServiceServer is the server API for FeedbackService service.
// All implementations must embed UnimplementedFeedbackServiceServer
// for forward compatibility
type FeedbackServiceServer interface {
	Get(context.Context, *FeedbackGetRequest) (*FeedbackGetResponse, error)
	GetAll(context.Context, *FilterOptions) (*FeedBackGetAllResponse, error)
	Update(context.Context, *FeedbackUpdateRequest) (*empty.Empty, error)
	Create(context.Context, *FeedbackCreateRequest) (*FeedBackCreateResponse, error)
	Delete(context.Context, *FeedbackDeleteRequest) (*empty.Empty, error)
	mustEmbedUnimplementedFeedbackServiceServer()
}

// UnimplementedFeedbackServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeedbackServiceServer struct {
}

func (UnimplementedFeedbackServiceServer) Get(context.Context, *FeedbackGetRequest) (*FeedbackGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFeedbackServiceServer) GetAll(context.Context, *FilterOptions) (*FeedBackGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedFeedbackServiceServer) Update(context.Context, *FeedbackUpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFeedbackServiceServer) Create(context.Context, *FeedbackCreateRequest) (*FeedBackCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFeedbackServiceServer) Delete(context.Context, *FeedbackDeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFeedbackServiceServer) mustEmbedUnimplementedFeedbackServiceServer() {}

// UnsafeFeedbackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedbackServiceServer will
// result in compilation errors.
type UnsafeFeedbackServiceServer interface {
	mustEmbedUnimplementedFeedbackServiceServer()
}

func RegisterFeedbackServiceServer(s grpc.ServiceRegistrar, srv FeedbackServiceServer) {
	s.RegisterService(&FeedbackService_ServiceDesc, srv)
}

func _FeedbackService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedbackGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FeedbackService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).Get(ctx, req.(*FeedbackGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedbackService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FeedbackService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).GetAll(ctx, req.(*FilterOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedbackService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedbackUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FeedbackService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).Update(ctx, req.(*FeedbackUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedbackService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedbackCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FeedbackService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).Create(ctx, req.(*FeedbackCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedbackService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedbackDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FeedbackService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).Delete(ctx, req.(*FeedbackDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedbackService_ServiceDesc is the grpc.ServiceDesc for FeedbackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedbackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.FeedbackService",
	HandlerType: (*FeedbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _FeedbackService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _FeedbackService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FeedbackService_Update_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FeedbackService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FeedbackService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feedback_service.proto",
}

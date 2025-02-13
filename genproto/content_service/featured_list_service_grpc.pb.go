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

// FeaturedListServiceClient is the client API for FeaturedListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeaturedListServiceClient interface {
	Create(ctx context.Context, in *FeaturedListCreateRequest, opts ...grpc.CallOption) (*FeaturedListCreateResponse, error)
	Get(ctx context.Context, in *FeaturedListGetRequest, opts ...grpc.CallOption) (*FeaturedListGetResponse, error)
	GetAll(ctx context.Context, in *FeaturedListFilterOptions, opts ...grpc.CallOption) (*FeaturedListGetAllResponse, error)
	Update(ctx context.Context, in *FeaturedListUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *FeaturedListDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type featuredListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeaturedListServiceClient(cc grpc.ClientConnInterface) FeaturedListServiceClient {
	return &featuredListServiceClient{cc}
}

func (c *featuredListServiceClient) Create(ctx context.Context, in *FeaturedListCreateRequest, opts ...grpc.CallOption) (*FeaturedListCreateResponse, error) {
	out := new(FeaturedListCreateResponse)
	err := c.cc.Invoke(ctx, "/content.FeaturedListService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featuredListServiceClient) Get(ctx context.Context, in *FeaturedListGetRequest, opts ...grpc.CallOption) (*FeaturedListGetResponse, error) {
	out := new(FeaturedListGetResponse)
	err := c.cc.Invoke(ctx, "/content.FeaturedListService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featuredListServiceClient) GetAll(ctx context.Context, in *FeaturedListFilterOptions, opts ...grpc.CallOption) (*FeaturedListGetAllResponse, error) {
	out := new(FeaturedListGetAllResponse)
	err := c.cc.Invoke(ctx, "/content.FeaturedListService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featuredListServiceClient) Update(ctx context.Context, in *FeaturedListUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.FeaturedListService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featuredListServiceClient) Delete(ctx context.Context, in *FeaturedListDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.FeaturedListService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeaturedListServiceServer is the server API for FeaturedListService service.
// All implementations must embed UnimplementedFeaturedListServiceServer
// for forward compatibility
type FeaturedListServiceServer interface {
	Create(context.Context, *FeaturedListCreateRequest) (*FeaturedListCreateResponse, error)
	Get(context.Context, *FeaturedListGetRequest) (*FeaturedListGetResponse, error)
	GetAll(context.Context, *FeaturedListFilterOptions) (*FeaturedListGetAllResponse, error)
	Update(context.Context, *FeaturedListUpdateRequest) (*empty.Empty, error)
	Delete(context.Context, *FeaturedListDeleteRequest) (*empty.Empty, error)
	mustEmbedUnimplementedFeaturedListServiceServer()
}

// UnimplementedFeaturedListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeaturedListServiceServer struct {
}

func (UnimplementedFeaturedListServiceServer) Create(context.Context, *FeaturedListCreateRequest) (*FeaturedListCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFeaturedListServiceServer) Get(context.Context, *FeaturedListGetRequest) (*FeaturedListGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFeaturedListServiceServer) GetAll(context.Context, *FeaturedListFilterOptions) (*FeaturedListGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedFeaturedListServiceServer) Update(context.Context, *FeaturedListUpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFeaturedListServiceServer) Delete(context.Context, *FeaturedListDeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFeaturedListServiceServer) mustEmbedUnimplementedFeaturedListServiceServer() {}

// UnsafeFeaturedListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeaturedListServiceServer will
// result in compilation errors.
type UnsafeFeaturedListServiceServer interface {
	mustEmbedUnimplementedFeaturedListServiceServer()
}

func RegisterFeaturedListServiceServer(s grpc.ServiceRegistrar, srv FeaturedListServiceServer) {
	s.RegisterService(&FeaturedListService_ServiceDesc, srv)
}

func _FeaturedListService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeaturedListCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeaturedListServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FeaturedListService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeaturedListServiceServer).Create(ctx, req.(*FeaturedListCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeaturedListService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeaturedListGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeaturedListServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FeaturedListService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeaturedListServiceServer).Get(ctx, req.(*FeaturedListGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeaturedListService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeaturedListFilterOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeaturedListServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FeaturedListService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeaturedListServiceServer).GetAll(ctx, req.(*FeaturedListFilterOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeaturedListService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeaturedListUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeaturedListServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FeaturedListService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeaturedListServiceServer).Update(ctx, req.(*FeaturedListUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeaturedListService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeaturedListDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeaturedListServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FeaturedListService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeaturedListServiceServer).Delete(ctx, req.(*FeaturedListDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeaturedListService_ServiceDesc is the grpc.ServiceDesc for FeaturedListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeaturedListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.FeaturedListService",
	HandlerType: (*FeaturedListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FeaturedListService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _FeaturedListService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _FeaturedListService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FeaturedListService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FeaturedListService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "featured_list_service.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: staff_service.proto

package content_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("staff_service.proto", fileDescriptor_26160db039211e88) }

var fileDescriptor_26160db039211e88 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x87, 0x42, 0x85, 0xb8, 0xab, 0x08, 0x63, 0x66, 0xbb, 0xda, 0x03, 0xa4, 0xa0, 0xb7,
	0x82, 0x4e, 0x9d, 0xbd, 0x14, 0x1c, 0xde, 0x78, 0x23, 0x5b, 0xfd, 0x5b, 0x0a, 0x69, 0x4e, 0x6c,
	0xce, 0x04, 0x9f, 0xd4, 0xd7, 0x91, 0x25, 0xad, 0x4c, 0x5a, 0xd8, 0x65, 0xbe, 0xff, 0x9c, 0x8f,
	0xf3, 0x13, 0x71, 0xe1, 0x79, 0x53, 0x14, 0xef, 0x1e, 0xcd, 0x57, 0x95, 0x43, 0xbb, 0x86, 0x98,
	0xe4, 0x59, 0x4e, 0x96, 0x61, 0x59, 0x9d, 0x87, 0x34, 0x52, 0x35, 0xce, 0xa9, 0xae, 0xc9, 0xb6,
	0xaf, 0x59, 0x49, 0x54, 0x1a, 0xa4, 0xe1, 0xb5, 0xdd, 0x15, 0x29, 0x6a, 0xc7, 0xdf, 0x31, 0xbc,
	0xfa, 0x39, 0x11, 0xe3, 0xf5, 0x7e, 0x75, 0x1d, 0xbd, 0xf2, 0x46, 0x9c, 0x66, 0x60, 0x39, 0xd5,
	0xad, 0x59, 0x87, 0x34, 0x03, 0xbf, 0xe0, 0x73, 0x07, 0xcf, 0xea, 0x72, 0x20, 0xf1, 0x8e, 0xac,
	0xc7, 0x62, 0x24, 0xef, 0x44, 0x92, 0x81, 0x97, 0xc6, 0xc8, 0xc9, 0xdf, 0xd8, 0x53, 0x65, 0x18,
	0xcd, 0xb3, 0xe3, 0x8a, 0xac, 0x57, 0xf3, 0xde, 0xfa, 0xd2, 0x98, 0x03, 0xc3, 0xad, 0x48, 0x5e,
	0xdd, 0xc7, 0x86, 0x21, 0x67, 0xff, 0x27, 0x23, 0xed, 0xae, 0x98, 0xe8, 0xd8, 0x4a, 0x77, 0xad,
	0xf4, 0x6a, 0xdf, 0x6a, 0x31, 0x92, 0x2b, 0x91, 0x3c, 0x34, 0x18, 0x10, 0x44, 0xda, 0x09, 0xe6,
	0xc3, 0xe1, 0xe1, 0x1d, 0x8f, 0x30, 0xe8, 0x6b, 0x22, 0x3d, 0x7a, 0xc7, 0xbd, 0x7a, 0x9b, 0x96,
	0xb0, 0x01, 0xa7, 0xad, 0xa0, 0xfb, 0xbc, 0x6d, 0x12, 0xf0, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x18, 0x8c, 0xa0, 0x74, 0xd4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StaffServiceClient is the client API for StaffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StaffServiceClient interface {
	Get(ctx context.Context, in *StaffGetRequest, opts ...grpc.CallOption) (*StaffGetResponse, error)
	GetAll(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*StaffGetAllResponse, error)
	Update(ctx context.Context, in *StaffUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Create(ctx context.Context, in *StaffCreateRequest, opts ...grpc.CallOption) (*StaffCreateResponse, error)
	Delete(ctx context.Context, in *StaffDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type staffServiceClient struct {
	cc *grpc.ClientConn
}

func NewStaffServiceClient(cc *grpc.ClientConn) StaffServiceClient {
	return &staffServiceClient{cc}
}

func (c *staffServiceClient) Get(ctx context.Context, in *StaffGetRequest, opts ...grpc.CallOption) (*StaffGetResponse, error) {
	out := new(StaffGetResponse)
	err := c.cc.Invoke(ctx, "/content.StaffService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) GetAll(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*StaffGetAllResponse, error) {
	out := new(StaffGetAllResponse)
	err := c.cc.Invoke(ctx, "/content.StaffService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Update(ctx context.Context, in *StaffUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.StaffService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Create(ctx context.Context, in *StaffCreateRequest, opts ...grpc.CallOption) (*StaffCreateResponse, error) {
	out := new(StaffCreateResponse)
	err := c.cc.Invoke(ctx, "/content.StaffService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Delete(ctx context.Context, in *StaffDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/content.StaffService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServiceServer is the server API for StaffService service.
type StaffServiceServer interface {
	Get(context.Context, *StaffGetRequest) (*StaffGetResponse, error)
	GetAll(context.Context, *FilterOptions) (*StaffGetAllResponse, error)
	Update(context.Context, *StaffUpdateRequest) (*empty.Empty, error)
	Create(context.Context, *StaffCreateRequest) (*StaffCreateResponse, error)
	Delete(context.Context, *StaffDeleteRequest) (*empty.Empty, error)
}

// UnimplementedStaffServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStaffServiceServer struct {
}

func (*UnimplementedStaffServiceServer) Get(ctx context.Context, req *StaffGetRequest) (*StaffGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedStaffServiceServer) GetAll(ctx context.Context, req *FilterOptions) (*StaffGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedStaffServiceServer) Update(ctx context.Context, req *StaffUpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedStaffServiceServer) Create(ctx context.Context, req *StaffCreateRequest) (*StaffCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedStaffServiceServer) Delete(ctx context.Context, req *StaffDeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterStaffServiceServer(s *grpc.Server, srv StaffServiceServer) {
	s.RegisterService(&_StaffService_serviceDesc, srv)
}

func _StaffService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.StaffService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Get(ctx, req.(*StaffGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.StaffService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).GetAll(ctx, req.(*FilterOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.StaffService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Update(ctx, req.(*StaffUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.StaffService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Create(ctx, req.(*StaffCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.StaffService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Delete(ctx, req.(*StaffDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StaffService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "content.StaffService",
	HandlerType: (*StaffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _StaffService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _StaffService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StaffService_Update_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _StaffService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StaffService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staff_service.proto",
}

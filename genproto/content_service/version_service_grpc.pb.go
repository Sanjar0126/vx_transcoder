// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package content_service

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

// VersionServiceClient is the client API for VersionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VersionServiceClient interface {
	Get(ctx context.Context, in *VersionGetRequest, opts ...grpc.CallOption) (*Version, error)
}

type versionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVersionServiceClient(cc grpc.ClientConnInterface) VersionServiceClient {
	return &versionServiceClient{cc}
}

func (c *versionServiceClient) Get(ctx context.Context, in *VersionGetRequest, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/content.VersionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionServiceServer is the server API for VersionService service.
// All implementations must embed UnimplementedVersionServiceServer
// for forward compatibility
type VersionServiceServer interface {
	Get(context.Context, *VersionGetRequest) (*Version, error)
	mustEmbedUnimplementedVersionServiceServer()
}

// UnimplementedVersionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVersionServiceServer struct {
}

func (UnimplementedVersionServiceServer) Get(context.Context, *VersionGetRequest) (*Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedVersionServiceServer) mustEmbedUnimplementedVersionServiceServer() {}

// UnsafeVersionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VersionServiceServer will
// result in compilation errors.
type UnsafeVersionServiceServer interface {
	mustEmbedUnimplementedVersionServiceServer()
}

func RegisterVersionServiceServer(s grpc.ServiceRegistrar, srv VersionServiceServer) {
	s.RegisterService(&VersionService_ServiceDesc, srv)
}

func _VersionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.VersionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).Get(ctx, req.(*VersionGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VersionService_ServiceDesc is the grpc.ServiceDesc for VersionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VersionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.VersionService",
	HandlerType: (*VersionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _VersionService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "version_service.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: movie_service.proto

package content_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_movie_service_proto protoreflect.FileDescriptor

var file_movie_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x5f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8d, 0x07, 0x0a, 0x0c, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x5a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x47, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12,
	0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x42, 0x79, 0x53,
	0x6c, 0x75, 0x67, 0x73, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x32, 0xcc, 0x04, 0x0a, 0x12, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x67, 0x6f, 0x67,
	0x6f, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x47, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x51, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1e, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x1f,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x45, 0x70,
	0x69, 0x73, 0x6f, 0x64, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x45,
	0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x00, 0x32, 0x87, 0x02, 0x0a, 0x0f, 0x52, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x12, 0x48, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x64, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_movie_service_proto_goTypes = []interface{}{
	(*MovieGetRequest)(nil),               // 0: content.MovieGetRequest
	(*MovieFilterOptions)(nil),            // 1: content.MovieFilterOptions
	(*MovieCreateRequest)(nil),            // 2: content.MovieCreateRequest
	(*MovieDeleteRequest)(nil),            // 3: content.MovieDeleteRequest
	(*MovieUpdateRequest)(nil),            // 4: content.MovieUpdateRequest
	(*UpdatePricingRequest)(nil),          // 5: content.UpdatePricingRequest
	(*RelatedMoviesRequest)(nil),          // 6: content.RelatedMoviesRequest
	(*PurchasedMoviesRequest)(nil),        // 7: content.PurchasedMoviesRequest
	(*ExistMovieFeaturedList)(nil),        // 8: content.ExistMovieFeaturedList
	(*empty.Empty)(nil),                   // 9: google.protobuf.Empty
	(*VendorGetStreamRequest)(nil),        // 10: content.VendorGetStreamRequest
	(*VendorSubscriptionRequest)(nil),     // 11: content.VendorSubscriptionRequest
	(*MovieSeasonGetRequest)(nil),         // 12: content.MovieSeasonGetRequest
	(*MovieEpisodeGetRequest)(nil),        // 13: content.MovieEpisodeGetRequest
	(*RecentlyWatchedCreateRequest)(nil),  // 14: content.RecentlyWatchedCreateRequest
	(*FilterOptions)(nil),                 // 15: content.FilterOptions
	(*MovieGetResponse)(nil),              // 16: content.MovieGetResponse
	(*MovieGetAllResponse)(nil),           // 17: content.MovieGetAllResponse
	(*MovieCreateResponse)(nil),           // 18: content.MovieCreateResponse
	(*MovieGetAllByGenresResponse)(nil),   // 19: content.MovieGetAllByGenresResponse
	(*RelatedMoviesResponse)(nil),         // 20: content.RelatedMoviesResponse
	(*VendorGetStreamResponse)(nil),       // 21: content.VendorGetStreamResponse
	(*VendorSubscriptionResponse)(nil),    // 22: content.VendorSubscriptionResponse
	(*MovieSeasonGetResponse)(nil),        // 23: content.MovieSeasonGetResponse
	(*MovieEpisodes)(nil),                 // 24: content.MovieEpisodes
	(*CreateResponse)(nil),                // 25: content.CreateResponse
	(*RecentlyWatchedMoviesResponse)(nil), // 26: content.RecentlyWatchedMoviesResponse
}
var file_movie_service_proto_depIdxs = []int32{
	0,  // 0: content.MovieService.Get:input_type -> content.MovieGetRequest
	1,  // 1: content.MovieService.GetAll:input_type -> content.MovieFilterOptions
	2,  // 2: content.MovieService.Create:input_type -> content.MovieCreateRequest
	3,  // 3: content.MovieService.Delete:input_type -> content.MovieDeleteRequest
	4,  // 4: content.MovieService.Update:input_type -> content.MovieUpdateRequest
	5,  // 5: content.MovieService.UpdatePricing:input_type -> content.UpdatePricingRequest
	1,  // 6: content.MovieService.GetAllMovieByCategory:input_type -> content.MovieFilterOptions
	6,  // 7: content.MovieService.GetRelatedMovies:input_type -> content.RelatedMoviesRequest
	7,  // 8: content.MovieService.GetMoviesBySlugs:input_type -> content.PurchasedMoviesRequest
	8,  // 9: content.MovieService.UpdateMovieExistFeaturedList:input_type -> content.ExistMovieFeaturedList
	1,  // 10: content.MovieService.Search:input_type -> content.MovieFilterOptions
	9,  // 11: content.MovieService.HealthCheck:input_type -> google.protobuf.Empty
	0,  // 12: content.MovieServiceMegogo.Get:input_type -> content.MovieGetRequest
	1,  // 13: content.MovieServiceMegogo.GetAllMovieByCategory:input_type -> content.MovieFilterOptions
	6,  // 14: content.MovieServiceMegogo.GetRelatedMovies:input_type -> content.RelatedMoviesRequest
	10, // 15: content.MovieServiceMegogo.GetStreams:input_type -> content.VendorGetStreamRequest
	11, // 16: content.MovieServiceMegogo.Subscription:input_type -> content.VendorSubscriptionRequest
	12, // 17: content.MovieServiceMegogo.GetSeason:input_type -> content.MovieSeasonGetRequest
	13, // 18: content.MovieServiceMegogo.GetEpisode:input_type -> content.MovieEpisodeGetRequest
	14, // 19: content.RecentlyWatched.Create:input_type -> content.RecentlyWatchedCreateRequest
	15, // 20: content.RecentlyWatched.GetUsersRecentlyWatched:input_type -> content.FilterOptions
	14, // 21: content.RecentlyWatched.MarkCompleted:input_type -> content.RecentlyWatchedCreateRequest
	16, // 22: content.MovieService.Get:output_type -> content.MovieGetResponse
	17, // 23: content.MovieService.GetAll:output_type -> content.MovieGetAllResponse
	18, // 24: content.MovieService.Create:output_type -> content.MovieCreateResponse
	9,  // 25: content.MovieService.Delete:output_type -> google.protobuf.Empty
	9,  // 26: content.MovieService.Update:output_type -> google.protobuf.Empty
	9,  // 27: content.MovieService.UpdatePricing:output_type -> google.protobuf.Empty
	19, // 28: content.MovieService.GetAllMovieByCategory:output_type -> content.MovieGetAllByGenresResponse
	20, // 29: content.MovieService.GetRelatedMovies:output_type -> content.RelatedMoviesResponse
	20, // 30: content.MovieService.GetMoviesBySlugs:output_type -> content.RelatedMoviesResponse
	9,  // 31: content.MovieService.UpdateMovieExistFeaturedList:output_type -> google.protobuf.Empty
	17, // 32: content.MovieService.Search:output_type -> content.MovieGetAllResponse
	9,  // 33: content.MovieService.HealthCheck:output_type -> google.protobuf.Empty
	16, // 34: content.MovieServiceMegogo.Get:output_type -> content.MovieGetResponse
	19, // 35: content.MovieServiceMegogo.GetAllMovieByCategory:output_type -> content.MovieGetAllByGenresResponse
	20, // 36: content.MovieServiceMegogo.GetRelatedMovies:output_type -> content.RelatedMoviesResponse
	21, // 37: content.MovieServiceMegogo.GetStreams:output_type -> content.VendorGetStreamResponse
	22, // 38: content.MovieServiceMegogo.Subscription:output_type -> content.VendorSubscriptionResponse
	23, // 39: content.MovieServiceMegogo.GetSeason:output_type -> content.MovieSeasonGetResponse
	24, // 40: content.MovieServiceMegogo.GetEpisode:output_type -> content.MovieEpisodes
	25, // 41: content.RecentlyWatched.Create:output_type -> content.CreateResponse
	26, // 42: content.RecentlyWatched.GetUsersRecentlyWatched:output_type -> content.RecentlyWatchedMoviesResponse
	25, // 43: content.RecentlyWatched.MarkCompleted:output_type -> content.CreateResponse
	22, // [22:44] is the sub-list for method output_type
	0,  // [0:22] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_movie_service_proto_init() }
func file_movie_service_proto_init() {
	if File_movie_service_proto != nil {
		return
	}
	file_common_proto_init()
	file_movie_proto_init()
	file_movie_season_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_movie_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_movie_service_proto_goTypes,
		DependencyIndexes: file_movie_service_proto_depIdxs,
	}.Build()
	File_movie_service_proto = out.File
	file_movie_service_proto_rawDesc = nil
	file_movie_service_proto_goTypes = nil
	file_movie_service_proto_depIdxs = nil
}

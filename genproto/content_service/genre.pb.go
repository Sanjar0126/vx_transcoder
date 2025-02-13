// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: genre.proto

package content_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenreGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name        string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Slug        string                     `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug"`
	Lang        string                     `protobuf:"bytes,4,opt,name=lang,proto3" json:"lang"`
	Active      bool                       `protobuf:"varint,5,opt,name=active,proto3" json:"active"`
	CreatedAt   string                     `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string                     `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Attributes  []*AttributeCommonResponse `protobuf:"bytes,8,rep,name=attributes,proto3" json:"attributes"`
	Description string                     `protobuf:"bytes,9,opt,name=description,proto3" json:"description"`
}

func (x *GenreGetResponse) Reset() {
	*x = GenreGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genre_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenreGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreGetResponse) ProtoMessage() {}

func (x *GenreGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_genre_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenreGetResponse.ProtoReflect.Descriptor instead.
func (*GenreGetResponse) Descriptor() ([]byte, []int) {
	return file_genre_proto_rawDescGZIP(), []int{0}
}

func (x *GenreGetResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GenreGetResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenreGetResponse) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GenreGetResponse) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *GenreGetResponse) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *GenreGetResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GenreGetResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GenreGetResponse) GetAttributes() []*AttributeCommonResponse {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *GenreGetResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GenreGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug"`
	Lang string `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang"`
}

func (x *GenreGetRequest) Reset() {
	*x = GenreGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genre_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenreGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreGetRequest) ProtoMessage() {}

func (x *GenreGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_genre_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenreGetRequest.ProtoReflect.Descriptor instead.
func (*GenreGetRequest) Descriptor() ([]byte, []int) {
	return file_genre_proto_rawDescGZIP(), []int{1}
}

func (x *GenreGetRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GenreGetRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

type GenreDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug"`
}

func (x *GenreDeleteRequest) Reset() {
	*x = GenreDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genre_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenreDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreDeleteRequest) ProtoMessage() {}

func (x *GenreDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_genre_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenreDeleteRequest.ProtoReflect.Descriptor instead.
func (*GenreDeleteRequest) Descriptor() ([]byte, []int) {
	return file_genre_proto_rawDescGZIP(), []int{2}
}

func (x *GenreDeleteRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type GenreCommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
}

func (x *GenreCommonResponse) Reset() {
	*x = GenreCommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genre_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenreCommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreCommonResponse) ProtoMessage() {}

func (x *GenreCommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_genre_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenreCommonResponse.ProtoReflect.Descriptor instead.
func (*GenreCommonResponse) Descriptor() ([]byte, []int) {
	return file_genre_proto_rawDescGZIP(), []int{3}
}

func (x *GenreCommonResponse) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GenreCommonResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GenreCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Attributes  []string `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
}

func (x *GenreCreateRequest) Reset() {
	*x = GenreCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genre_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenreCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreCreateRequest) ProtoMessage() {}

func (x *GenreCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_genre_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenreCreateRequest.ProtoReflect.Descriptor instead.
func (*GenreCreateRequest) Descriptor() ([]byte, []int) {
	return file_genre_proto_rawDescGZIP(), []int{4}
}

func (x *GenreCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenreCreateRequest) GetAttributes() []string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *GenreCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GenreCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug"`
}

func (x *GenreCreateResponse) Reset() {
	*x = GenreCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genre_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenreCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreCreateResponse) ProtoMessage() {}

func (x *GenreCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_genre_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenreCreateResponse.ProtoReflect.Descriptor instead.
func (*GenreCreateResponse) Descriptor() ([]byte, []int) {
	return file_genre_proto_rawDescGZIP(), []int{5}
}

func (x *GenreCreateResponse) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type GenreUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug        string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug"`
	Lang        string   `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Active      bool     `protobuf:"varint,4,opt,name=active,proto3" json:"active"`
	Attributes  []string `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes"`
	Description string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description"`
}

func (x *GenreUpdateRequest) Reset() {
	*x = GenreUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genre_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenreUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreUpdateRequest) ProtoMessage() {}

func (x *GenreUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_genre_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenreUpdateRequest.ProtoReflect.Descriptor instead.
func (*GenreUpdateRequest) Descriptor() ([]byte, []int) {
	return file_genre_proto_rawDescGZIP(), []int{6}
}

func (x *GenreUpdateRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GenreUpdateRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *GenreUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenreUpdateRequest) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *GenreUpdateRequest) GetAttributes() []string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *GenreUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GenreGetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*GenreGetResponse `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects"`
	Count   int32               `protobuf:"varint,2,opt,name=count,proto3" json:"count"`
}

func (x *GenreGetAllResponse) Reset() {
	*x = GenreGetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genre_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenreGetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreGetAllResponse) ProtoMessage() {}

func (x *GenreGetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_genre_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenreGetAllResponse.ProtoReflect.Descriptor instead.
func (*GenreGetAllResponse) Descriptor() ([]byte, []int) {
	return file_genre_proto_rawDescGZIP(), []int{7}
}

func (x *GenreGetAllResponse) GetObjects() []*GenreGetResponse {
	if x != nil {
		return x.Objects
	}
	return nil
}

func (x *GenreGetAllResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_genre_proto protoreflect.FileDescriptor

var file_genre_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x02, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x6c, 0x75, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x40,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x39, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x28, 0x0a,
	0x12, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x22, 0x3d, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x72, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6a, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x22, 0xaa, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x60, 0x0a, 0x13, 0x47, 0x65,
	0x6e, 0x72, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x6e,
	0x72, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x1a, 0x5a, 0x18,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_genre_proto_rawDescOnce sync.Once
	file_genre_proto_rawDescData = file_genre_proto_rawDesc
)

func file_genre_proto_rawDescGZIP() []byte {
	file_genre_proto_rawDescOnce.Do(func() {
		file_genre_proto_rawDescData = protoimpl.X.CompressGZIP(file_genre_proto_rawDescData)
	})
	return file_genre_proto_rawDescData
}

var file_genre_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_genre_proto_goTypes = []interface{}{
	(*GenreGetResponse)(nil),        // 0: content.GenreGetResponse
	(*GenreGetRequest)(nil),         // 1: content.GenreGetRequest
	(*GenreDeleteRequest)(nil),      // 2: content.GenreDeleteRequest
	(*GenreCommonResponse)(nil),     // 3: content.GenreCommonResponse
	(*GenreCreateRequest)(nil),      // 4: content.GenreCreateRequest
	(*GenreCreateResponse)(nil),     // 5: content.GenreCreateResponse
	(*GenreUpdateRequest)(nil),      // 6: content.GenreUpdateRequest
	(*GenreGetAllResponse)(nil),     // 7: content.GenreGetAllResponse
	(*AttributeCommonResponse)(nil), // 8: content.AttributeCommonResponse
}
var file_genre_proto_depIdxs = []int32{
	8, // 0: content.GenreGetResponse.attributes:type_name -> content.AttributeCommonResponse
	0, // 1: content.GenreGetAllResponse.objects:type_name -> content.GenreGetResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_genre_proto_init() }
func file_genre_proto_init() {
	if File_genre_proto != nil {
		return
	}
	file_attribute_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_genre_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenreGetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_genre_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenreGetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_genre_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenreDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_genre_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenreCommonResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_genre_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenreCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_genre_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenreCreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_genre_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenreUpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_genre_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenreGetAllResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_genre_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_genre_proto_goTypes,
		DependencyIndexes: file_genre_proto_depIdxs,
		MessageInfos:      file_genre_proto_msgTypes,
	}.Build()
	File_genre_proto = out.File
	file_genre_proto_rawDesc = nil
	file_genre_proto_goTypes = nil
	file_genre_proto_depIdxs = nil
}

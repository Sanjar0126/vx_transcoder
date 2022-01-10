// Code generated by protoc-gen-go. DO NOT EDIT.
// source: genre.proto

package content_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GenreGetResponse struct {
	Id                   string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slug                 string                     `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Lang                 string                     `protobuf:"bytes,4,opt,name=lang,proto3" json:"lang,omitempty"`
	Active               bool                       `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
	CreatedAt            string                     `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string                     `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Attributes           []*AttributeCommonResponse `protobuf:"bytes,8,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GenreGetResponse) Reset()         { *m = GenreGetResponse{} }
func (m *GenreGetResponse) String() string { return proto.CompactTextString(m) }
func (*GenreGetResponse) ProtoMessage()    {}
func (*GenreGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b8f62296ba87e, []int{0}
}

func (m *GenreGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenreGetResponse.Unmarshal(m, b)
}
func (m *GenreGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenreGetResponse.Marshal(b, m, deterministic)
}
func (m *GenreGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenreGetResponse.Merge(m, src)
}
func (m *GenreGetResponse) XXX_Size() int {
	return xxx_messageInfo_GenreGetResponse.Size(m)
}
func (m *GenreGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenreGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenreGetResponse proto.InternalMessageInfo

func (m *GenreGetResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GenreGetResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GenreGetResponse) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *GenreGetResponse) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *GenreGetResponse) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *GenreGetResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GenreGetResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *GenreGetResponse) GetAttributes() []*AttributeCommonResponse {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type GenreGetRequest struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Lang                 string   `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenreGetRequest) Reset()         { *m = GenreGetRequest{} }
func (m *GenreGetRequest) String() string { return proto.CompactTextString(m) }
func (*GenreGetRequest) ProtoMessage()    {}
func (*GenreGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b8f62296ba87e, []int{1}
}

func (m *GenreGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenreGetRequest.Unmarshal(m, b)
}
func (m *GenreGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenreGetRequest.Marshal(b, m, deterministic)
}
func (m *GenreGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenreGetRequest.Merge(m, src)
}
func (m *GenreGetRequest) XXX_Size() int {
	return xxx_messageInfo_GenreGetRequest.Size(m)
}
func (m *GenreGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenreGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenreGetRequest proto.InternalMessageInfo

func (m *GenreGetRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *GenreGetRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

type GenreDeleteRequest struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenreDeleteRequest) Reset()         { *m = GenreDeleteRequest{} }
func (m *GenreDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*GenreDeleteRequest) ProtoMessage()    {}
func (*GenreDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b8f62296ba87e, []int{2}
}

func (m *GenreDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenreDeleteRequest.Unmarshal(m, b)
}
func (m *GenreDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenreDeleteRequest.Marshal(b, m, deterministic)
}
func (m *GenreDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenreDeleteRequest.Merge(m, src)
}
func (m *GenreDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_GenreDeleteRequest.Size(m)
}
func (m *GenreDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenreDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenreDeleteRequest proto.InternalMessageInfo

func (m *GenreDeleteRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type GenreCommonResponse struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenreCommonResponse) Reset()         { *m = GenreCommonResponse{} }
func (m *GenreCommonResponse) String() string { return proto.CompactTextString(m) }
func (*GenreCommonResponse) ProtoMessage()    {}
func (*GenreCommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b8f62296ba87e, []int{3}
}

func (m *GenreCommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenreCommonResponse.Unmarshal(m, b)
}
func (m *GenreCommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenreCommonResponse.Marshal(b, m, deterministic)
}
func (m *GenreCommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenreCommonResponse.Merge(m, src)
}
func (m *GenreCommonResponse) XXX_Size() int {
	return xxx_messageInfo_GenreCommonResponse.Size(m)
}
func (m *GenreCommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenreCommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenreCommonResponse proto.InternalMessageInfo

func (m *GenreCommonResponse) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *GenreCommonResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GenreCreateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Attributes           []string `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenreCreateRequest) Reset()         { *m = GenreCreateRequest{} }
func (m *GenreCreateRequest) String() string { return proto.CompactTextString(m) }
func (*GenreCreateRequest) ProtoMessage()    {}
func (*GenreCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b8f62296ba87e, []int{4}
}

func (m *GenreCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenreCreateRequest.Unmarshal(m, b)
}
func (m *GenreCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenreCreateRequest.Marshal(b, m, deterministic)
}
func (m *GenreCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenreCreateRequest.Merge(m, src)
}
func (m *GenreCreateRequest) XXX_Size() int {
	return xxx_messageInfo_GenreCreateRequest.Size(m)
}
func (m *GenreCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenreCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenreCreateRequest proto.InternalMessageInfo

func (m *GenreCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GenreCreateRequest) GetAttributes() []string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type GenreCreateResponse struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenreCreateResponse) Reset()         { *m = GenreCreateResponse{} }
func (m *GenreCreateResponse) String() string { return proto.CompactTextString(m) }
func (*GenreCreateResponse) ProtoMessage()    {}
func (*GenreCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b8f62296ba87e, []int{5}
}

func (m *GenreCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenreCreateResponse.Unmarshal(m, b)
}
func (m *GenreCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenreCreateResponse.Marshal(b, m, deterministic)
}
func (m *GenreCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenreCreateResponse.Merge(m, src)
}
func (m *GenreCreateResponse) XXX_Size() int {
	return xxx_messageInfo_GenreCreateResponse.Size(m)
}
func (m *GenreCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenreCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenreCreateResponse proto.InternalMessageInfo

func (m *GenreCreateResponse) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type GenreUpdateRequest struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Lang                 string   `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Active               bool     `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	Attributes           []string `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenreUpdateRequest) Reset()         { *m = GenreUpdateRequest{} }
func (m *GenreUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*GenreUpdateRequest) ProtoMessage()    {}
func (*GenreUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b8f62296ba87e, []int{6}
}

func (m *GenreUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenreUpdateRequest.Unmarshal(m, b)
}
func (m *GenreUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenreUpdateRequest.Marshal(b, m, deterministic)
}
func (m *GenreUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenreUpdateRequest.Merge(m, src)
}
func (m *GenreUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_GenreUpdateRequest.Size(m)
}
func (m *GenreUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenreUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenreUpdateRequest proto.InternalMessageInfo

func (m *GenreUpdateRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *GenreUpdateRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *GenreUpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GenreUpdateRequest) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *GenreUpdateRequest) GetAttributes() []string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type GenreGetAllResponse struct {
	Objects              []*GenreGetResponse `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	Count                int32               `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GenreGetAllResponse) Reset()         { *m = GenreGetAllResponse{} }
func (m *GenreGetAllResponse) String() string { return proto.CompactTextString(m) }
func (*GenreGetAllResponse) ProtoMessage()    {}
func (*GenreGetAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b8f62296ba87e, []int{7}
}

func (m *GenreGetAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenreGetAllResponse.Unmarshal(m, b)
}
func (m *GenreGetAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenreGetAllResponse.Marshal(b, m, deterministic)
}
func (m *GenreGetAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenreGetAllResponse.Merge(m, src)
}
func (m *GenreGetAllResponse) XXX_Size() int {
	return xxx_messageInfo_GenreGetAllResponse.Size(m)
}
func (m *GenreGetAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenreGetAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenreGetAllResponse proto.InternalMessageInfo

func (m *GenreGetAllResponse) GetObjects() []*GenreGetResponse {
	if m != nil {
		return m.Objects
	}
	return nil
}

func (m *GenreGetAllResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*GenreGetResponse)(nil), "content.GenreGetResponse")
	proto.RegisterType((*GenreGetRequest)(nil), "content.GenreGetRequest")
	proto.RegisterType((*GenreDeleteRequest)(nil), "content.GenreDeleteRequest")
	proto.RegisterType((*GenreCommonResponse)(nil), "content.GenreCommonResponse")
	proto.RegisterType((*GenreCreateRequest)(nil), "content.GenreCreateRequest")
	proto.RegisterType((*GenreCreateResponse)(nil), "content.GenreCreateResponse")
	proto.RegisterType((*GenreUpdateRequest)(nil), "content.GenreUpdateRequest")
	proto.RegisterType((*GenreGetAllResponse)(nil), "content.GenreGetAllResponse")
}

func init() { proto.RegisterFile("genre.proto", fileDescriptor_498b8f62296ba87e) }

var fileDescriptor_498b8f62296ba87e = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x4b, 0xeb, 0x40,
	0x10, 0x66, 0xf3, 0xa3, 0x3f, 0xb6, 0xf0, 0xfa, 0xd8, 0xf7, 0x78, 0xec, 0x2b, 0x28, 0x21, 0xa7,
	0x78, 0xa9, 0x60, 0x4f, 0x1e, 0x04, 0x63, 0x85, 0x7a, 0x0e, 0x78, 0xf1, 0x52, 0xd3, 0x64, 0x08,
	0x91, 0x74, 0xb7, 0x26, 0x93, 0xfe, 0x0d, 0xfe, 0xc5, 0x9e, 0x25, 0x9b, 0x4d, 0x9a, 0x54, 0x2b,
	0x78, 0x9b, 0x7c, 0x33, 0xb3, 0xf3, 0x7d, 0xdf, 0x4c, 0xe8, 0x24, 0x01, 0x91, 0xc3, 0x7c, 0x97,
	0x4b, 0x94, 0x6c, 0x18, 0x49, 0x81, 0x20, 0x70, 0x36, 0x0d, 0x11, 0xf3, 0x74, 0x53, 0xa2, 0xce,
	0xb8, 0xef, 0x84, 0xfe, 0x5e, 0x55, 0x95, 0x2b, 0xc0, 0x00, 0x8a, 0x9d, 0x14, 0x05, 0xb0, 0x5f,
	0xd4, 0x48, 0x63, 0x4e, 0x1c, 0xe2, 0x8d, 0x03, 0x23, 0x8d, 0x19, 0xa3, 0x96, 0x08, 0xb7, 0xc0,
	0x0d, 0x85, 0xa8, 0xb8, 0xc2, 0x8a, 0xac, 0x4c, 0xb8, 0x59, 0x63, 0x55, 0x5c, 0x61, 0x59, 0x28,
	0x12, 0x6e, 0xd5, 0x58, 0x15, 0xb3, 0x7f, 0x74, 0x10, 0x46, 0x98, 0xee, 0x81, 0xdb, 0x0e, 0xf1,
	0x46, 0x81, 0xfe, 0x62, 0x67, 0x94, 0x46, 0x39, 0x84, 0x08, 0xf1, 0x3a, 0x44, 0x3e, 0x50, 0x1d,
	0x63, 0x8d, 0xf8, 0x58, 0xa5, 0xcb, 0x5d, 0xdc, 0xa4, 0x87, 0x75, 0x5a, 0x23, 0x3e, 0xb2, 0x5b,
	0x4a, 0x5b, 0x25, 0x05, 0x1f, 0x39, 0xa6, 0x37, 0xb9, 0x72, 0xe6, 0x5a, 0xe5, 0xdc, 0x6f, 0x52,
	0x4b, 0xb9, 0xdd, 0x4a, 0xd1, 0xe8, 0x0a, 0x3a, 0x3d, 0xee, 0x35, 0x9d, 0x1e, 0x74, 0xbf, 0x96,
	0x50, 0x60, 0x2b, 0x89, 0x7c, 0x21, 0xc9, 0x38, 0x48, 0x72, 0x3d, 0xca, 0x54, 0xeb, 0x3d, 0x64,
	0x80, 0xf0, 0x4d, 0xb7, 0x7b, 0x43, 0xff, 0xa8, 0xca, 0x3e, 0x8f, 0x53, 0x83, 0x8e, 0x3d, 0x76,
	0x1f, 0xf4, 0xa0, 0xa5, 0xb2, 0xa5, 0x33, 0x48, 0x55, 0x92, 0xce, 0x36, 0xce, 0x7b, 0x7e, 0x58,
	0x8e, 0xe9, 0x8d, 0x7b, 0x6a, 0x2f, 0x1a, 0x22, 0xfa, 0xa5, 0xd3, 0x44, 0xdc, 0x37, 0xa2, 0xa7,
	0x3e, 0x2a, 0xb7, 0x7f, 0x68, 0x4e, 0xcb, 0xce, 0xec, 0xb0, 0x3b, 0xdc, 0x80, 0xd5, 0xbb, 0x81,
	0x3e, 0x6b, 0xfb, 0x13, 0xeb, 0x67, 0xcd, 0x7a, 0x05, 0xe8, 0x67, 0x59, 0xcb, 0x7a, 0x41, 0x87,
	0x72, 0xf3, 0x02, 0x11, 0x16, 0x9c, 0xa8, 0xcd, 0xff, 0x6f, 0x37, 0x7f, 0x7c, 0xca, 0x41, 0x53,
	0xc9, 0xfe, 0x52, 0x3b, 0x92, 0xa5, 0x40, 0x45, 0xd6, 0x0e, 0xea, 0x8f, 0xbb, 0xd9, 0x13, 0x4f,
	0x40, 0xa8, 0x5f, 0xe1, 0x52, 0xbf, 0xb1, 0x2e, 0x20, 0xdf, 0xa7, 0x11, 0x6c, 0x06, 0x0a, 0x5e,
	0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x7d, 0xb0, 0xf1, 0x4a, 0x03, 0x00, 0x00,
}

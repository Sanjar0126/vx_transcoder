// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: upload_video.proto

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

type UploadVideoCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieSlug     string `protobuf:"bytes,1,opt,name=movie_slug,json=movieSlug,proto3" json:"movie_slug"`
	SeasonNumber  int32  `protobuf:"varint,2,opt,name=season_number,json=seasonNumber,proto3" json:"season_number"`
	EpisodeNumber int32  `protobuf:"varint,3,opt,name=episode_number,json=episodeNumber,proto3" json:"episode_number"`
	Id            string `protobuf:"bytes,4,opt,name=id,proto3" json:"id"`
	Type          string `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	Stage         string `protobuf:"bytes,6,opt,name=stage,proto3" json:"stage"`
}

func (x *UploadVideoCreateRequest) Reset() {
	*x = UploadVideoCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoCreateRequest) ProtoMessage() {}

func (x *UploadVideoCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoCreateRequest.ProtoReflect.Descriptor instead.
func (*UploadVideoCreateRequest) Descriptor() ([]byte, []int) {
	return file_upload_video_proto_rawDescGZIP(), []int{0}
}

func (x *UploadVideoCreateRequest) GetMovieSlug() string {
	if x != nil {
		return x.MovieSlug
	}
	return ""
}

func (x *UploadVideoCreateRequest) GetSeasonNumber() int32 {
	if x != nil {
		return x.SeasonNumber
	}
	return 0
}

func (x *UploadVideoCreateRequest) GetEpisodeNumber() int32 {
	if x != nil {
		return x.EpisodeNumber
	}
	return 0
}

func (x *UploadVideoCreateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UploadVideoCreateRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UploadVideoCreateRequest) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

type UploadVideoGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *UploadVideoGetRequest) Reset() {
	*x = UploadVideoGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoGetRequest) ProtoMessage() {}

func (x *UploadVideoGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoGetRequest.ProtoReflect.Descriptor instead.
func (*UploadVideoGetRequest) Descriptor() ([]byte, []int) {
	return file_upload_video_proto_rawDescGZIP(), []int{1}
}

func (x *UploadVideoGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UploadVideoGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path"`
}

func (x *UploadVideoGetResponse) Reset() {
	*x = UploadVideoGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoGetResponse) ProtoMessage() {}

func (x *UploadVideoGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoGetResponse.ProtoReflect.Descriptor instead.
func (*UploadVideoGetResponse) Descriptor() ([]byte, []int) {
	return file_upload_video_proto_rawDescGZIP(), []int{2}
}

func (x *UploadVideoGetResponse) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type UploadedVideoGetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage string `protobuf:"bytes,1,opt,name=stage,proto3" json:"stage"`
}

func (x *UploadedVideoGetAllRequest) Reset() {
	*x = UploadedVideoGetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadedVideoGetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadedVideoGetAllRequest) ProtoMessage() {}

func (x *UploadedVideoGetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadedVideoGetAllRequest.ProtoReflect.Descriptor instead.
func (*UploadedVideoGetAllRequest) Descriptor() ([]byte, []int) {
	return file_upload_video_proto_rawDescGZIP(), []int{3}
}

func (x *UploadedVideoGetAllRequest) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

type UploadedVideoGetFullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	MovieSlug     string `protobuf:"bytes,2,opt,name=movie_slug,json=movieSlug,proto3" json:"movie_slug"`
	SeasonNumber  int32  `protobuf:"varint,3,opt,name=season_number,json=seasonNumber,proto3" json:"season_number"`
	EpisodeNumber int32  `protobuf:"varint,4,opt,name=episode_number,json=episodeNumber,proto3" json:"episode_number"`
	Type          string `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	Stage         string `protobuf:"bytes,6,opt,name=stage,proto3" json:"stage"`
}

func (x *UploadedVideoGetFullResponse) Reset() {
	*x = UploadedVideoGetFullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadedVideoGetFullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadedVideoGetFullResponse) ProtoMessage() {}

func (x *UploadedVideoGetFullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadedVideoGetFullResponse.ProtoReflect.Descriptor instead.
func (*UploadedVideoGetFullResponse) Descriptor() ([]byte, []int) {
	return file_upload_video_proto_rawDescGZIP(), []int{4}
}

func (x *UploadedVideoGetFullResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UploadedVideoGetFullResponse) GetMovieSlug() string {
	if x != nil {
		return x.MovieSlug
	}
	return ""
}

func (x *UploadedVideoGetFullResponse) GetSeasonNumber() int32 {
	if x != nil {
		return x.SeasonNumber
	}
	return 0
}

func (x *UploadedVideoGetFullResponse) GetEpisodeNumber() int32 {
	if x != nil {
		return x.EpisodeNumber
	}
	return 0
}

func (x *UploadedVideoGetFullResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UploadedVideoGetFullResponse) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

type UploadedVideoGetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*UploadedVideoGetFullResponse `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects"`
}

func (x *UploadedVideoGetAllResponse) Reset() {
	*x = UploadedVideoGetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadedVideoGetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadedVideoGetAllResponse) ProtoMessage() {}

func (x *UploadedVideoGetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadedVideoGetAllResponse.ProtoReflect.Descriptor instead.
func (*UploadedVideoGetAllResponse) Descriptor() ([]byte, []int) {
	return file_upload_video_proto_rawDescGZIP(), []int{5}
}

func (x *UploadedVideoGetAllResponse) GetObjects() []*UploadedVideoGetFullResponse {
	if x != nil {
		return x.Objects
	}
	return nil
}

type UpdateUploadedVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Stage string `protobuf:"bytes,2,opt,name=stage,proto3" json:"stage"`
}

func (x *UpdateUploadedVideoRequest) Reset() {
	*x = UpdateUploadedVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_video_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUploadedVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUploadedVideoRequest) ProtoMessage() {}

func (x *UpdateUploadedVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_video_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUploadedVideoRequest.ProtoReflect.Descriptor instead.
func (*UpdateUploadedVideoRequest) Descriptor() ([]byte, []int) {
	return file_upload_video_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUploadedVideoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUploadedVideoRequest) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

type DeleteUploadedVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteUploadedVideoRequest) Reset() {
	*x = DeleteUploadedVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_video_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUploadedVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUploadedVideoRequest) ProtoMessage() {}

func (x *DeleteUploadedVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_video_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUploadedVideoRequest.ProtoReflect.Descriptor instead.
func (*DeleteUploadedVideoRequest) Descriptor() ([]byte, []int) {
	return file_upload_video_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUploadedVideoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_upload_video_proto protoreflect.FileDescriptor

var file_upload_video_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xbf, 0x01,
	0x0a, 0x18, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25,
	0x0a, 0x0e, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x22,
	0x27, 0x0a, 0x15, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x16, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x32, 0x0a, 0x1a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x73, 0x6c,
	0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x53,
	0x6c, 0x75, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x70, 0x69, 0x73,
	0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x22, 0x5e, 0x0a, 0x1b, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x42, 0x0a, 0x1a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a,
	0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x1a, 0x5a, 0x18, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_upload_video_proto_rawDescOnce sync.Once
	file_upload_video_proto_rawDescData = file_upload_video_proto_rawDesc
)

func file_upload_video_proto_rawDescGZIP() []byte {
	file_upload_video_proto_rawDescOnce.Do(func() {
		file_upload_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_upload_video_proto_rawDescData)
	})
	return file_upload_video_proto_rawDescData
}

var file_upload_video_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_upload_video_proto_goTypes = []interface{}{
	(*UploadVideoCreateRequest)(nil),     // 0: content.UploadVideoCreateRequest
	(*UploadVideoGetRequest)(nil),        // 1: content.UploadVideoGetRequest
	(*UploadVideoGetResponse)(nil),       // 2: content.UploadVideoGetResponse
	(*UploadedVideoGetAllRequest)(nil),   // 3: content.UploadedVideoGetAllRequest
	(*UploadedVideoGetFullResponse)(nil), // 4: content.UploadedVideoGetFullResponse
	(*UploadedVideoGetAllResponse)(nil),  // 5: content.UploadedVideoGetAllResponse
	(*UpdateUploadedVideoRequest)(nil),   // 6: content.UpdateUploadedVideoRequest
	(*DeleteUploadedVideoRequest)(nil),   // 7: content.DeleteUploadedVideoRequest
}
var file_upload_video_proto_depIdxs = []int32{
	4, // 0: content.UploadedVideoGetAllResponse.objects:type_name -> content.UploadedVideoGetFullResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_upload_video_proto_init() }
func file_upload_video_proto_init() {
	if File_upload_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upload_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoCreateRequest); i {
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
		file_upload_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoGetRequest); i {
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
		file_upload_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoGetResponse); i {
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
		file_upload_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadedVideoGetAllRequest); i {
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
		file_upload_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadedVideoGetFullResponse); i {
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
		file_upload_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadedVideoGetAllResponse); i {
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
		file_upload_video_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUploadedVideoRequest); i {
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
		file_upload_video_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUploadedVideoRequest); i {
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
			RawDescriptor: file_upload_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_upload_video_proto_goTypes,
		DependencyIndexes: file_upload_video_proto_depIdxs,
		MessageInfos:      file_upload_video_proto_msgTypes,
	}.Build()
	File_upload_video_proto = out.File
	file_upload_video_proto_rawDesc = nil
	file_upload_video_proto_goTypes = nil
	file_upload_video_proto_depIdxs = nil
}

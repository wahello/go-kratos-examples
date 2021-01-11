// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: api/blog/v1/tag.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateTagRequest) Reset() {
	*x = CreateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagRequest) ProtoMessage() {}

func (x *CreateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagRequest.ProtoReflect.Descriptor instead.
func (*CreateTagRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_tag_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTagRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CreateTagRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateTagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTagReply) Reset() {
	*x = CreateTagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagReply) ProtoMessage() {}

func (x *CreateTagReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagReply.ProtoReflect.Descriptor instead.
func (*CreateTagReply) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_tag_proto_rawDescGZIP(), []int{1}
}

type UpdateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTagRequest) Reset() {
	*x = UpdateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTagRequest) ProtoMessage() {}

func (x *UpdateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTagRequest.ProtoReflect.Descriptor instead.
func (*UpdateTagRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_tag_proto_rawDescGZIP(), []int{2}
}

type UpdateTagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTagReply) Reset() {
	*x = UpdateTagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTagReply) ProtoMessage() {}

func (x *UpdateTagReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTagReply.ProtoReflect.Descriptor instead.
func (*UpdateTagReply) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_tag_proto_rawDescGZIP(), []int{3}
}

type DeleteTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTagRequest) Reset() {
	*x = DeleteTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagRequest) ProtoMessage() {}

func (x *DeleteTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagRequest.ProtoReflect.Descriptor instead.
func (*DeleteTagRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_tag_proto_rawDescGZIP(), []int{4}
}

type DeleteTagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTagReply) Reset() {
	*x = DeleteTagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagReply) ProtoMessage() {}

func (x *DeleteTagReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagReply.ProtoReflect.Descriptor instead.
func (*DeleteTagReply) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_tag_proto_rawDescGZIP(), []int{5}
}

type GetTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTagRequest) Reset() {
	*x = GetTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_tag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagRequest) ProtoMessage() {}

func (x *GetTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_tag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagRequest.ProtoReflect.Descriptor instead.
func (*GetTagRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_tag_proto_rawDescGZIP(), []int{6}
}

type GetTagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTagReply) Reset() {
	*x = GetTagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_tag_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagReply) ProtoMessage() {}

func (x *GetTagReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_tag_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagReply.ProtoReflect.Descriptor instead.
func (*GetTagReply) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_tag_proto_rawDescGZIP(), []int{7}
}

type ListTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTagRequest) Reset() {
	*x = ListTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_tag_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagRequest) ProtoMessage() {}

func (x *ListTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_tag_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagRequest.ProtoReflect.Descriptor instead.
func (*ListTagRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_tag_proto_rawDescGZIP(), []int{8}
}

type ListTagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTagReply) Reset() {
	*x = ListTagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_tag_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagReply) ProtoMessage() {}

func (x *ListTagReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_tag_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagReply.ProtoReflect.Descriptor instead.
func (*ListTagReply) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_tag_proto_rawDescGZIP(), []int{9}
}

var File_api_blog_v1_tag_proto protoreflect.FileDescriptor

var file_api_blog_v1_tag_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x22, 0x3a, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x0e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x32, 0xe3, 0x02, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x47, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x47, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x47, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x3e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x12, 0x1b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x24, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x13, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_blog_v1_tag_proto_rawDescOnce sync.Once
	file_api_blog_v1_tag_proto_rawDescData = file_api_blog_v1_tag_proto_rawDesc
)

func file_api_blog_v1_tag_proto_rawDescGZIP() []byte {
	file_api_blog_v1_tag_proto_rawDescOnce.Do(func() {
		file_api_blog_v1_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_blog_v1_tag_proto_rawDescData)
	})
	return file_api_blog_v1_tag_proto_rawDescData
}

var file_api_blog_v1_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_blog_v1_tag_proto_goTypes = []interface{}{
	(*CreateTagRequest)(nil), // 0: api.blog.v1.CreateTagRequest
	(*CreateTagReply)(nil),   // 1: api.blog.v1.CreateTagReply
	(*UpdateTagRequest)(nil), // 2: api.blog.v1.UpdateTagRequest
	(*UpdateTagReply)(nil),   // 3: api.blog.v1.UpdateTagReply
	(*DeleteTagRequest)(nil), // 4: api.blog.v1.DeleteTagRequest
	(*DeleteTagReply)(nil),   // 5: api.blog.v1.DeleteTagReply
	(*GetTagRequest)(nil),    // 6: api.blog.v1.GetTagRequest
	(*GetTagReply)(nil),      // 7: api.blog.v1.GetTagReply
	(*ListTagRequest)(nil),   // 8: api.blog.v1.ListTagRequest
	(*ListTagReply)(nil),     // 9: api.blog.v1.ListTagReply
}
var file_api_blog_v1_tag_proto_depIdxs = []int32{
	0, // 0: api.blog.v1.Tag.CreateTag:input_type -> api.blog.v1.CreateTagRequest
	2, // 1: api.blog.v1.Tag.UpdateTag:input_type -> api.blog.v1.UpdateTagRequest
	4, // 2: api.blog.v1.Tag.DeleteTag:input_type -> api.blog.v1.DeleteTagRequest
	6, // 3: api.blog.v1.Tag.GetTag:input_type -> api.blog.v1.GetTagRequest
	8, // 4: api.blog.v1.Tag.ListTag:input_type -> api.blog.v1.ListTagRequest
	1, // 5: api.blog.v1.Tag.CreateTag:output_type -> api.blog.v1.CreateTagReply
	3, // 6: api.blog.v1.Tag.UpdateTag:output_type -> api.blog.v1.UpdateTagReply
	5, // 7: api.blog.v1.Tag.DeleteTag:output_type -> api.blog.v1.DeleteTagReply
	7, // 8: api.blog.v1.Tag.GetTag:output_type -> api.blog.v1.GetTagReply
	9, // 9: api.blog.v1.Tag.ListTag:output_type -> api.blog.v1.ListTagReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_blog_v1_tag_proto_init() }
func file_api_blog_v1_tag_proto_init() {
	if File_api_blog_v1_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_blog_v1_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagRequest); i {
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
		file_api_blog_v1_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagReply); i {
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
		file_api_blog_v1_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTagRequest); i {
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
		file_api_blog_v1_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTagReply); i {
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
		file_api_blog_v1_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagRequest); i {
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
		file_api_blog_v1_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagReply); i {
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
		file_api_blog_v1_tag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagRequest); i {
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
		file_api_blog_v1_tag_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagReply); i {
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
		file_api_blog_v1_tag_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagRequest); i {
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
		file_api_blog_v1_tag_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagReply); i {
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
			RawDescriptor: file_api_blog_v1_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_blog_v1_tag_proto_goTypes,
		DependencyIndexes: file_api_blog_v1_tag_proto_depIdxs,
		MessageInfos:      file_api_blog_v1_tag_proto_msgTypes,
	}.Build()
	File_api_blog_v1_tag_proto = out.File
	file_api_blog_v1_tag_proto_rawDesc = nil
	file_api_blog_v1_tag_proto_goTypes = nil
	file_api_blog_v1_tag_proto_depIdxs = nil
}

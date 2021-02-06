// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: api/blog/v1/comment.proto

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

type CreateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCommentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCommentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateCommentReply) Reset() {
	*x = CreateCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentReply) ProtoMessage() {}

func (x *CreateCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentReply.ProtoReflect.Descriptor instead.
func (*CreateCommentReply) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_comment_proto_rawDescGZIP(), []int{1}
}

type UpdateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCommentRequest) Reset() {
	*x = UpdateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentRequest) ProtoMessage() {}

func (x *UpdateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentRequest.ProtoReflect.Descriptor instead.
func (*UpdateCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_comment_proto_rawDescGZIP(), []int{2}
}

type UpdateCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCommentReply) Reset() {
	*x = UpdateCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentReply) ProtoMessage() {}

func (x *UpdateCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentReply.ProtoReflect.Descriptor instead.
func (*UpdateCommentReply) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_comment_proto_rawDescGZIP(), []int{3}
}

type DeleteCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCommentRequest) Reset() {
	*x = DeleteCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentRequest) ProtoMessage() {}

func (x *DeleteCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_comment_proto_rawDescGZIP(), []int{4}
}

type DeleteCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCommentReply) Reset() {
	*x = DeleteCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentReply) ProtoMessage() {}

func (x *DeleteCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentReply.ProtoReflect.Descriptor instead.
func (*DeleteCommentReply) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_comment_proto_rawDescGZIP(), []int{5}
}

type GetCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCommentRequest) Reset() {
	*x = GetCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentRequest) ProtoMessage() {}

func (x *GetCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentRequest.ProtoReflect.Descriptor instead.
func (*GetCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_comment_proto_rawDescGZIP(), []int{6}
}

type GetCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCommentReply) Reset() {
	*x = GetCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentReply) ProtoMessage() {}

func (x *GetCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentReply.ProtoReflect.Descriptor instead.
func (*GetCommentReply) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_comment_proto_rawDescGZIP(), []int{7}
}

type ListCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCommentRequest) Reset() {
	*x = ListCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_comment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentRequest) ProtoMessage() {}

func (x *ListCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_comment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentRequest.ProtoReflect.Descriptor instead.
func (*ListCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_comment_proto_rawDescGZIP(), []int{8}
}

type ListCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCommentReply) Reset() {
	*x = ListCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_v1_comment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentReply) ProtoMessage() {}

func (x *ListCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_v1_comment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentReply.ProtoReflect.Descriptor instead.
func (*ListCommentReply) Descriptor() ([]byte, []int) {
	return file_api_blog_v1_comment_proto_rawDescGZIP(), []int{9}
}

var File_api_blog_v1_comment_proto protoreflect.FileDescriptor

var file_api_blog_v1_comment_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x22, 0x44, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x14,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x32, 0xa3, 0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x53,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x53, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x53, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4a, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4d, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x24, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x13, 0x62, 0x6c, 0x6f, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_blog_v1_comment_proto_rawDescOnce sync.Once
	file_api_blog_v1_comment_proto_rawDescData = file_api_blog_v1_comment_proto_rawDesc
)

func file_api_blog_v1_comment_proto_rawDescGZIP() []byte {
	file_api_blog_v1_comment_proto_rawDescOnce.Do(func() {
		file_api_blog_v1_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_blog_v1_comment_proto_rawDescData)
	})
	return file_api_blog_v1_comment_proto_rawDescData
}

var file_api_blog_v1_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_blog_v1_comment_proto_goTypes = []interface{}{
	(*CreateCommentRequest)(nil), // 0: api.blog.v1.CreateCommentRequest
	(*CreateCommentReply)(nil),   // 1: api.blog.v1.CreateCommentReply
	(*UpdateCommentRequest)(nil), // 2: api.blog.v1.UpdateCommentRequest
	(*UpdateCommentReply)(nil),   // 3: api.blog.v1.UpdateCommentReply
	(*DeleteCommentRequest)(nil), // 4: api.blog.v1.DeleteCommentRequest
	(*DeleteCommentReply)(nil),   // 5: api.blog.v1.DeleteCommentReply
	(*GetCommentRequest)(nil),    // 6: api.blog.v1.GetCommentRequest
	(*GetCommentReply)(nil),      // 7: api.blog.v1.GetCommentReply
	(*ListCommentRequest)(nil),   // 8: api.blog.v1.ListCommentRequest
	(*ListCommentReply)(nil),     // 9: api.blog.v1.ListCommentReply
}
var file_api_blog_v1_comment_proto_depIdxs = []int32{
	0, // 0: api.blog.v1.Comment.CreateComment:input_type -> api.blog.v1.CreateCommentRequest
	2, // 1: api.blog.v1.Comment.UpdateComment:input_type -> api.blog.v1.UpdateCommentRequest
	4, // 2: api.blog.v1.Comment.DeleteComment:input_type -> api.blog.v1.DeleteCommentRequest
	6, // 3: api.blog.v1.Comment.GetComment:input_type -> api.blog.v1.GetCommentRequest
	8, // 4: api.blog.v1.Comment.ListComment:input_type -> api.blog.v1.ListCommentRequest
	1, // 5: api.blog.v1.Comment.CreateComment:output_type -> api.blog.v1.CreateCommentReply
	3, // 6: api.blog.v1.Comment.UpdateComment:output_type -> api.blog.v1.UpdateCommentReply
	5, // 7: api.blog.v1.Comment.DeleteComment:output_type -> api.blog.v1.DeleteCommentReply
	7, // 8: api.blog.v1.Comment.GetComment:output_type -> api.blog.v1.GetCommentReply
	9, // 9: api.blog.v1.Comment.ListComment:output_type -> api.blog.v1.ListCommentReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_blog_v1_comment_proto_init() }
func file_api_blog_v1_comment_proto_init() {
	if File_api_blog_v1_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_blog_v1_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentRequest); i {
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
		file_api_blog_v1_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentReply); i {
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
		file_api_blog_v1_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentRequest); i {
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
		file_api_blog_v1_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentReply); i {
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
		file_api_blog_v1_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentRequest); i {
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
		file_api_blog_v1_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentReply); i {
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
		file_api_blog_v1_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentRequest); i {
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
		file_api_blog_v1_comment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentReply); i {
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
		file_api_blog_v1_comment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentRequest); i {
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
		file_api_blog_v1_comment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentReply); i {
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
			RawDescriptor: file_api_blog_v1_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_blog_v1_comment_proto_goTypes,
		DependencyIndexes: file_api_blog_v1_comment_proto_depIdxs,
		MessageInfos:      file_api_blog_v1_comment_proto_msgTypes,
	}.Build()
	File_api_blog_v1_comment_proto = out.File
	file_api_blog_v1_comment_proto_rawDesc = nil
	file_api_blog_v1_comment_proto_goTypes = nil
	file_api_blog_v1_comment_proto_depIdxs = nil
}

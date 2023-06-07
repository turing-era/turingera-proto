// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: knowledge.proto

package knowledge

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

type KnowledgeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KId     uint32 `protobuf:"varint,1,opt,name=k_id,json=kId,proto3" json:"k_id,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RobotId string `protobuf:"bytes,3,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`
	Text    string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *KnowledgeInfo) Reset() {
	*x = KnowledgeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_knowledge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KnowledgeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KnowledgeInfo) ProtoMessage() {}

func (x *KnowledgeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_knowledge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KnowledgeInfo.ProtoReflect.Descriptor instead.
func (*KnowledgeInfo) Descriptor() ([]byte, []int) {
	return file_knowledge_proto_rawDescGZIP(), []int{0}
}

func (x *KnowledgeInfo) GetKId() uint32 {
	if x != nil {
		return x.KId
	}
	return 0
}

func (x *KnowledgeInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *KnowledgeInfo) GetRobotId() string {
	if x != nil {
		return x.RobotId
	}
	return ""
}

func (x *KnowledgeInfo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetKnowledgeListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetKnowledgeListReq) Reset() {
	*x = GetKnowledgeListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_knowledge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKnowledgeListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKnowledgeListReq) ProtoMessage() {}

func (x *GetKnowledgeListReq) ProtoReflect() protoreflect.Message {
	mi := &file_knowledge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKnowledgeListReq.ProtoReflect.Descriptor instead.
func (*GetKnowledgeListReq) Descriptor() ([]byte, []int) {
	return file_knowledge_proto_rawDescGZIP(), []int{1}
}

type GetKnowledgeListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*KnowledgeInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *GetKnowledgeListRsp) Reset() {
	*x = GetKnowledgeListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_knowledge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKnowledgeListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKnowledgeListRsp) ProtoMessage() {}

func (x *GetKnowledgeListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_knowledge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKnowledgeListRsp.ProtoReflect.Descriptor instead.
func (*GetKnowledgeListRsp) Descriptor() ([]byte, []int) {
	return file_knowledge_proto_rawDescGZIP(), []int{2}
}

func (x *GetKnowledgeListRsp) GetInfos() []*KnowledgeInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type SaveKnowledgeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *KnowledgeInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SaveKnowledgeReq) Reset() {
	*x = SaveKnowledgeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_knowledge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveKnowledgeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveKnowledgeReq) ProtoMessage() {}

func (x *SaveKnowledgeReq) ProtoReflect() protoreflect.Message {
	mi := &file_knowledge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveKnowledgeReq.ProtoReflect.Descriptor instead.
func (*SaveKnowledgeReq) Descriptor() ([]byte, []int) {
	return file_knowledge_proto_rawDescGZIP(), []int{3}
}

func (x *SaveKnowledgeReq) GetInfo() *KnowledgeInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type SaveKnowledgeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *KnowledgeInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SaveKnowledgeRsp) Reset() {
	*x = SaveKnowledgeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_knowledge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveKnowledgeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveKnowledgeRsp) ProtoMessage() {}

func (x *SaveKnowledgeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_knowledge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveKnowledgeRsp.ProtoReflect.Descriptor instead.
func (*SaveKnowledgeRsp) Descriptor() ([]byte, []int) {
	return file_knowledge_proto_rawDescGZIP(), []int{4}
}

func (x *SaveKnowledgeRsp) GetInfo() *KnowledgeInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type RemKnowledgeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RobotId string `protobuf:"bytes,1,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`
	KId     uint32 `protobuf:"varint,2,opt,name=k_id,json=kId,proto3" json:"k_id,omitempty"`
}

func (x *RemKnowledgeReq) Reset() {
	*x = RemKnowledgeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_knowledge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemKnowledgeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemKnowledgeReq) ProtoMessage() {}

func (x *RemKnowledgeReq) ProtoReflect() protoreflect.Message {
	mi := &file_knowledge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemKnowledgeReq.ProtoReflect.Descriptor instead.
func (*RemKnowledgeReq) Descriptor() ([]byte, []int) {
	return file_knowledge_proto_rawDescGZIP(), []int{5}
}

func (x *RemKnowledgeReq) GetRobotId() string {
	if x != nil {
		return x.RobotId
	}
	return ""
}

func (x *RemKnowledgeReq) GetKId() uint32 {
	if x != nil {
		return x.KId
	}
	return 0
}

type RemKnowledgeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemKnowledgeRsp) Reset() {
	*x = RemKnowledgeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_knowledge_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemKnowledgeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemKnowledgeRsp) ProtoMessage() {}

func (x *RemKnowledgeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_knowledge_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemKnowledgeRsp.ProtoReflect.Descriptor instead.
func (*RemKnowledgeRsp) Descriptor() ([]byte, []int) {
	return file_knowledge_proto_rawDescGZIP(), []int{6}
}

type LoadKnowledgeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RobotId string `protobuf:"bytes,1,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`
}

func (x *LoadKnowledgeReq) Reset() {
	*x = LoadKnowledgeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_knowledge_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadKnowledgeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadKnowledgeReq) ProtoMessage() {}

func (x *LoadKnowledgeReq) ProtoReflect() protoreflect.Message {
	mi := &file_knowledge_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadKnowledgeReq.ProtoReflect.Descriptor instead.
func (*LoadKnowledgeReq) Descriptor() ([]byte, []int) {
	return file_knowledge_proto_rawDescGZIP(), []int{7}
}

func (x *LoadKnowledgeReq) GetRobotId() string {
	if x != nil {
		return x.RobotId
	}
	return ""
}

type LoadKnowledgeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoadKnowledgeRsp) Reset() {
	*x = LoadKnowledgeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_knowledge_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadKnowledgeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadKnowledgeRsp) ProtoMessage() {}

func (x *LoadKnowledgeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_knowledge_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadKnowledgeRsp.ProtoReflect.Descriptor instead.
func (*LoadKnowledgeRsp) Descriptor() ([]byte, []int) {
	return file_knowledge_proto_rawDescGZIP(), []int{8}
}

var File_knowledge_proto protoreflect.FileDescriptor

var file_knowledge_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x22, 0x6a, 0x0a,
	0x0d, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11,
	0x0a, 0x04, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x22, 0x56, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x3f, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x51, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65,
	0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x3d, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x75, 0x72,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x51, 0x0a, 0x10, 0x53,
	0x61, 0x76, 0x65, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52, 0x73, 0x70, 0x12,
	0x3d, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x3f,
	0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x11, 0x0a, 0x04,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x49, 0x64, 0x22,
	0x11, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52,
	0x73, 0x70, 0x22, 0x2d, 0x0a, 0x10, 0x4c, 0x6f, 0x61, 0x64, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49,
	0x64, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x6f, 0x61, 0x64, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x52, 0x73, 0x70, 0x32, 0xc8, 0x03, 0x0a, 0x0c, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x53, 0x76, 0x72, 0x12, 0x74, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4b, 0x6e, 0x6f,
	0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x2e, 0x74, 0x75, 0x72,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e, 0x74, 0x75,
	0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x6e, 0x6f, 0x77,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x6b, 0x0a, 0x0d,
	0x53, 0x61, 0x76, 0x65, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x12, 0x2c, 0x2e,
	0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2c, 0x2e, 0x74, 0x75,
	0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4b, 0x6e, 0x6f,
	0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52, 0x73, 0x70, 0x12, 0x68, 0x0a, 0x0c, 0x52, 0x65, 0x6d,
	0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x12, 0x2b, 0x2e, 0x74, 0x75, 0x72, 0x69,
	0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x52, 0x73, 0x70, 0x12, 0x6b, 0x0a, 0x0d, 0x4c, 0x6f, 0x61, 0x64, 0x4b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x12, 0x2c, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x2c, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e,
	0x4c, 0x6f, 0x61, 0x64, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52, 0x73, 0x70,
	0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x75, 0x72, 0x69, 0x6e, 0x67, 0x2d, 0x65, 0x72, 0x61, 0x2f, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_knowledge_proto_rawDescOnce sync.Once
	file_knowledge_proto_rawDescData = file_knowledge_proto_rawDesc
)

func file_knowledge_proto_rawDescGZIP() []byte {
	file_knowledge_proto_rawDescOnce.Do(func() {
		file_knowledge_proto_rawDescData = protoimpl.X.CompressGZIP(file_knowledge_proto_rawDescData)
	})
	return file_knowledge_proto_rawDescData
}

var file_knowledge_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_knowledge_proto_goTypes = []interface{}{
	(*KnowledgeInfo)(nil),       // 0: turingera.server.knowledge.KnowledgeInfo
	(*GetKnowledgeListReq)(nil), // 1: turingera.server.knowledge.GetKnowledgeListReq
	(*GetKnowledgeListRsp)(nil), // 2: turingera.server.knowledge.GetKnowledgeListRsp
	(*SaveKnowledgeReq)(nil),    // 3: turingera.server.knowledge.SaveKnowledgeReq
	(*SaveKnowledgeRsp)(nil),    // 4: turingera.server.knowledge.SaveKnowledgeRsp
	(*RemKnowledgeReq)(nil),     // 5: turingera.server.knowledge.RemKnowledgeReq
	(*RemKnowledgeRsp)(nil),     // 6: turingera.server.knowledge.RemKnowledgeRsp
	(*LoadKnowledgeReq)(nil),    // 7: turingera.server.knowledge.LoadKnowledgeReq
	(*LoadKnowledgeRsp)(nil),    // 8: turingera.server.knowledge.LoadKnowledgeRsp
}
var file_knowledge_proto_depIdxs = []int32{
	0, // 0: turingera.server.knowledge.GetKnowledgeListRsp.infos:type_name -> turingera.server.knowledge.KnowledgeInfo
	0, // 1: turingera.server.knowledge.SaveKnowledgeReq.info:type_name -> turingera.server.knowledge.KnowledgeInfo
	0, // 2: turingera.server.knowledge.SaveKnowledgeRsp.info:type_name -> turingera.server.knowledge.KnowledgeInfo
	1, // 3: turingera.server.knowledge.KnowledgeSvr.GetKnowledgeList:input_type -> turingera.server.knowledge.GetKnowledgeListReq
	3, // 4: turingera.server.knowledge.KnowledgeSvr.SaveKnowledge:input_type -> turingera.server.knowledge.SaveKnowledgeReq
	5, // 5: turingera.server.knowledge.KnowledgeSvr.RemKnowledge:input_type -> turingera.server.knowledge.RemKnowledgeReq
	7, // 6: turingera.server.knowledge.KnowledgeSvr.LoadKnowledge:input_type -> turingera.server.knowledge.LoadKnowledgeReq
	2, // 7: turingera.server.knowledge.KnowledgeSvr.GetKnowledgeList:output_type -> turingera.server.knowledge.GetKnowledgeListRsp
	4, // 8: turingera.server.knowledge.KnowledgeSvr.SaveKnowledge:output_type -> turingera.server.knowledge.SaveKnowledgeRsp
	6, // 9: turingera.server.knowledge.KnowledgeSvr.RemKnowledge:output_type -> turingera.server.knowledge.RemKnowledgeRsp
	8, // 10: turingera.server.knowledge.KnowledgeSvr.LoadKnowledge:output_type -> turingera.server.knowledge.LoadKnowledgeRsp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_knowledge_proto_init() }
func file_knowledge_proto_init() {
	if File_knowledge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_knowledge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KnowledgeInfo); i {
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
		file_knowledge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKnowledgeListReq); i {
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
		file_knowledge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKnowledgeListRsp); i {
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
		file_knowledge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveKnowledgeReq); i {
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
		file_knowledge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveKnowledgeRsp); i {
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
		file_knowledge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemKnowledgeReq); i {
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
		file_knowledge_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemKnowledgeRsp); i {
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
		file_knowledge_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadKnowledgeReq); i {
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
		file_knowledge_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadKnowledgeRsp); i {
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
			RawDescriptor: file_knowledge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_knowledge_proto_goTypes,
		DependencyIndexes: file_knowledge_proto_depIdxs,
		MessageInfos:      file_knowledge_proto_msgTypes,
	}.Build()
	File_knowledge_proto = out.File
	file_knowledge_proto_rawDesc = nil
	file_knowledge_proto_goTypes = nil
	file_knowledge_proto_depIdxs = nil
}
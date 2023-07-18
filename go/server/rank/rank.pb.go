// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: rank.proto

package rank

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

type RankType int32

const (
	RankType_RANK_TYPE_UNKNOWN RankType = 0 // 未知
	RankType_RANK_TYPE_SCORE   RankType = 1 // 积分
	RankType_RANK_TYPE_INVITE  RankType = 2 // 邀请
)

// Enum value maps for RankType.
var (
	RankType_name = map[int32]string{
		0: "RANK_TYPE_UNKNOWN",
		1: "RANK_TYPE_SCORE",
		2: "RANK_TYPE_INVITE",
	}
	RankType_value = map[string]int32{
		"RANK_TYPE_UNKNOWN": 0,
		"RANK_TYPE_SCORE":   1,
		"RANK_TYPE_INVITE":  2,
	}
)

func (x RankType) Enum() *RankType {
	p := new(RankType)
	*p = x
	return p
}

func (x RankType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RankType) Descriptor() protoreflect.EnumDescriptor {
	return file_rank_proto_enumTypes[0].Descriptor()
}

func (RankType) Type() protoreflect.EnumType {
	return &file_rank_proto_enumTypes[0]
}

func (x RankType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RankType.Descriptor instead.
func (RankType) EnumDescriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{0}
}

type RankID int32

const (
	RankID_RANK_ID_UNKNOWN RankID = 0   // 未知
	RankID_RANK_ID_WEEK    RankID = 100 // 周榜
	RankID_RANK_ID_MONTH   RankID = 200 // 月榜
)

// Enum value maps for RankID.
var (
	RankID_name = map[int32]string{
		0:   "RANK_ID_UNKNOWN",
		100: "RANK_ID_WEEK",
		200: "RANK_ID_MONTH",
	}
	RankID_value = map[string]int32{
		"RANK_ID_UNKNOWN": 0,
		"RANK_ID_WEEK":    100,
		"RANK_ID_MONTH":   200,
	}
)

func (x RankID) Enum() *RankID {
	p := new(RankID)
	*p = x
	return p
}

func (x RankID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RankID) Descriptor() protoreflect.EnumDescriptor {
	return file_rank_proto_enumTypes[1].Descriptor()
}

func (RankID) Type() protoreflect.EnumType {
	return &file_rank_proto_enumTypes[1]
}

func (x RankID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RankID.Descriptor instead.
func (RankID) EnumDescriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{1}
}

type RankInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StartTime int64  `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64  `protobuf:"varint,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	RankType  uint32 `protobuf:"varint,4,opt,name=rank_type,json=rankType,proto3" json:"rank_type,omitempty"` // RankType
	RankId    uint32 `protobuf:"varint,5,opt,name=rank_id,json=rankId,proto3" json:"rank_id,omitempty"`       // RankID
}

func (x *RankInfo) Reset() {
	*x = RankInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankInfo) ProtoMessage() {}

func (x *RankInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankInfo.ProtoReflect.Descriptor instead.
func (*RankInfo) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{0}
}

func (x *RankInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RankInfo) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *RankInfo) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *RankInfo) GetRankType() uint32 {
	if x != nil {
		return x.RankType
	}
	return 0
}

func (x *RankInfo) GetRankId() uint32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

type RankItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Score  uint32 `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *RankItem) Reset() {
	*x = RankItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankItem) ProtoMessage() {}

func (x *RankItem) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankItem.ProtoReflect.Descriptor instead.
func (*RankItem) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{1}
}

func (x *RankItem) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RankItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RankItem) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *RankItem) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type GetRankListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankType uint32 `protobuf:"varint,1,opt,name=rank_type,json=rankType,proto3" json:"rank_type,omitempty"` // RankType
	RankId   uint32 `protobuf:"varint,2,opt,name=rank_id,json=rankId,proto3" json:"rank_id,omitempty"`       // RankID
	Cookie   []byte `protobuf:"bytes,10,opt,name=cookie,proto3" json:"cookie,omitempty"`
}

func (x *GetRankListReq) Reset() {
	*x = GetRankListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankListReq) ProtoMessage() {}

func (x *GetRankListReq) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankListReq.ProtoReflect.Descriptor instead.
func (*GetRankListReq) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{2}
}

func (x *GetRankListReq) GetRankType() uint32 {
	if x != nil {
		return x.RankType
	}
	return 0
}

func (x *GetRankListReq) GetRankId() uint32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *GetRankListReq) GetCookie() []byte {
	if x != nil {
		return x.Cookie
	}
	return nil
}

type GetRankListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankInfo *RankInfo   `protobuf:"bytes,1,opt,name=rank_info,json=rankInfo,proto3" json:"rank_info,omitempty"`
	Items    []*RankItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Cookie   []byte      `protobuf:"bytes,10,opt,name=cookie,proto3" json:"cookie,omitempty"`
	IsEnd    bool        `protobuf:"varint,11,opt,name=is_end,json=isEnd,proto3" json:"is_end,omitempty"`
}

func (x *GetRankListRsp) Reset() {
	*x = GetRankListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankListRsp) ProtoMessage() {}

func (x *GetRankListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankListRsp.ProtoReflect.Descriptor instead.
func (*GetRankListRsp) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{3}
}

func (x *GetRankListRsp) GetRankInfo() *RankInfo {
	if x != nil {
		return x.RankInfo
	}
	return nil
}

func (x *GetRankListRsp) GetItems() []*RankItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GetRankListRsp) GetCookie() []byte {
	if x != nil {
		return x.Cookie
	}
	return nil
}

func (x *GetRankListRsp) GetIsEnd() bool {
	if x != nil {
		return x.IsEnd
	}
	return false
}

type UpdateScoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Score    uint32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	RankType uint32 `protobuf:"varint,3,opt,name=rank_type,json=rankType,proto3" json:"rank_type,omitempty"` // RankType
}

func (x *UpdateScoreReq) Reset() {
	*x = UpdateScoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScoreReq) ProtoMessage() {}

func (x *UpdateScoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScoreReq.ProtoReflect.Descriptor instead.
func (*UpdateScoreReq) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateScoreReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateScoreReq) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *UpdateScoreReq) GetRankType() uint32 {
	if x != nil {
		return x.RankType
	}
	return 0
}

type UpdateScoreRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateScoreRsp) Reset() {
	*x = UpdateScoreRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScoreRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScoreRsp) ProtoMessage() {}

func (x *UpdateScoreRsp) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScoreRsp.ProtoReflect.Descriptor instead.
func (*UpdateScoreRsp) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{5}
}

var File_rank_proto protoreflect.FileDescriptor

var file_rank_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x75,
	0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x22, 0x8e, 0x01, 0x0a, 0x08, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x72, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61,
	0x6e, 0x6b, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x08, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x5e, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x72, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61,
	0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x6e,
	0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x3c,
	0x0a, 0x09, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x72, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x75,
	0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x69,
	0x73, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x45,
	0x6e, 0x64, 0x22, 0x5c, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x73, 0x70, 0x2a, 0x4c, 0x0a, 0x08, 0x52, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15,
	0x0a, 0x11, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x41,
	0x4e, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x10, 0x02,
	0x2a, 0x43, 0x0a, 0x06, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x44, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x41,
	0x4e, 0x4b, 0x5f, 0x49, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x10, 0x0a, 0x0c, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x49, 0x44, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x10,
	0x64, 0x12, 0x12, 0x0a, 0x0d, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x49, 0x44, 0x5f, 0x4d, 0x4f, 0x4e,
	0x54, 0x48, 0x10, 0xc8, 0x01, 0x32, 0xc3, 0x01, 0x0a, 0x07, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x76,
	0x72, 0x12, 0x5b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x25, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x5b,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x25, 0x2e,
	0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x73, 0x70, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67,
	0x2d, 0x65, 0x72, 0x61, 0x2f, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72,
	0x61, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rank_proto_rawDescOnce sync.Once
	file_rank_proto_rawDescData = file_rank_proto_rawDesc
)

func file_rank_proto_rawDescGZIP() []byte {
	file_rank_proto_rawDescOnce.Do(func() {
		file_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_rank_proto_rawDescData)
	})
	return file_rank_proto_rawDescData
}

var file_rank_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rank_proto_goTypes = []interface{}{
	(RankType)(0),          // 0: turingera.server.rank.RankType
	(RankID)(0),            // 1: turingera.server.rank.RankID
	(*RankInfo)(nil),       // 2: turingera.server.rank.RankInfo
	(*RankItem)(nil),       // 3: turingera.server.rank.RankItem
	(*GetRankListReq)(nil), // 4: turingera.server.rank.GetRankListReq
	(*GetRankListRsp)(nil), // 5: turingera.server.rank.GetRankListRsp
	(*UpdateScoreReq)(nil), // 6: turingera.server.rank.UpdateScoreReq
	(*UpdateScoreRsp)(nil), // 7: turingera.server.rank.UpdateScoreRsp
}
var file_rank_proto_depIdxs = []int32{
	2, // 0: turingera.server.rank.GetRankListRsp.rank_info:type_name -> turingera.server.rank.RankInfo
	3, // 1: turingera.server.rank.GetRankListRsp.items:type_name -> turingera.server.rank.RankItem
	4, // 2: turingera.server.rank.RankSvr.GetRankList:input_type -> turingera.server.rank.GetRankListReq
	6, // 3: turingera.server.rank.RankSvr.UpdateScore:input_type -> turingera.server.rank.UpdateScoreReq
	5, // 4: turingera.server.rank.RankSvr.GetRankList:output_type -> turingera.server.rank.GetRankListRsp
	7, // 5: turingera.server.rank.RankSvr.UpdateScore:output_type -> turingera.server.rank.UpdateScoreRsp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rank_proto_init() }
func file_rank_proto_init() {
	if File_rank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankInfo); i {
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
		file_rank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankItem); i {
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
		file_rank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankListReq); i {
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
		file_rank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankListRsp); i {
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
		file_rank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScoreReq); i {
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
		file_rank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScoreRsp); i {
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
			RawDescriptor: file_rank_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rank_proto_goTypes,
		DependencyIndexes: file_rank_proto_depIdxs,
		EnumInfos:         file_rank_proto_enumTypes,
		MessageInfos:      file_rank_proto_msgTypes,
	}.Build()
	File_rank_proto = out.File
	file_rank_proto_rawDesc = nil
	file_rank_proto_goTypes = nil
	file_rank_proto_depIdxs = nil
}

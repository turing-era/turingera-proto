// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: chat.proto

package chat

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

type UserType int32

const (
	UserType_USER_TYPE_HUMAN UserType = 0
	UserType_USER_TYPE_ROBOT UserType = 1
)

// Enum value maps for UserType.
var (
	UserType_name = map[int32]string{
		0: "USER_TYPE_HUMAN",
		1: "USER_TYPE_ROBOT",
	}
	UserType_value = map[string]int32{
		"USER_TYPE_HUMAN": 0,
		"USER_TYPE_ROBOT": 1,
	}
)

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_chat_proto_enumTypes[0].Descriptor()
}

func (UserType) Type() protoreflect.EnumType {
	return &file_chat_proto_enumTypes[0]
}

func (x UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserType.Descriptor instead.
func (UserType) EnumDescriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User      string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Text      string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Message) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId          string     `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	ChatTime        uint32     `protobuf:"varint,2,opt,name=chat_time,json=chatTime,proto3" json:"chat_time,omitempty"`
	UserId          string     `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt       int64      `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	NumParticipants uint32     `protobuf:"varint,5,opt,name=num_participants,json=numParticipants,proto3" json:"num_participants,omitempty"`
	PartnerId       string     `protobuf:"bytes,6,opt,name=partner_id,json=partnerId,proto3" json:"partner_id,omitempty"`
	PartnerName     string     `protobuf:"bytes,7,opt,name=partner_name,json=partnerName,proto3" json:"partner_name,omitempty"`
	PartnerAvatar   string     `protobuf:"bytes,8,opt,name=partner_avatar,json=partnerAvatar,proto3" json:"partner_avatar,omitempty"`
	IsMyTurn        bool       `protobuf:"varint,9,opt,name=is_my_turn,json=isMyTurn,proto3" json:"is_my_turn,omitempty"`
	TurnTime        uint32     `protobuf:"varint,10,opt,name=turn_time,json=turnTime,proto3" json:"turn_time,omitempty"`
	IsActive        bool       `protobuf:"varint,11,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Assign          bool       `protobuf:"varint,12,opt,name=assign,proto3" json:"assign,omitempty"`
	Messages        []*Message `protobuf:"bytes,20,rep,name=messages,proto3" json:"messages,omitempty"`
	Guessed         bool       `protobuf:"varint,50,opt,name=guessed,proto3" json:"guessed,omitempty"`
	Correct         bool       `protobuf:"varint,51,opt,name=correct,proto3" json:"correct,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Info) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *Info) GetChatTime() uint32 {
	if x != nil {
		return x.ChatTime
	}
	return 0
}

func (x *Info) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Info) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Info) GetNumParticipants() uint32 {
	if x != nil {
		return x.NumParticipants
	}
	return 0
}

func (x *Info) GetPartnerId() string {
	if x != nil {
		return x.PartnerId
	}
	return ""
}

func (x *Info) GetPartnerName() string {
	if x != nil {
		return x.PartnerName
	}
	return ""
}

func (x *Info) GetPartnerAvatar() string {
	if x != nil {
		return x.PartnerAvatar
	}
	return ""
}

func (x *Info) GetIsMyTurn() bool {
	if x != nil {
		return x.IsMyTurn
	}
	return false
}

func (x *Info) GetTurnTime() uint32 {
	if x != nil {
		return x.TurnTime
	}
	return 0
}

func (x *Info) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Info) GetAssign() bool {
	if x != nil {
		return x.Assign
	}
	return false
}

func (x *Info) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *Info) GetGuessed() bool {
	if x != nil {
		return x.Guessed
	}
	return false
}

func (x *Info) GetCorrect() bool {
	if x != nil {
		return x.Correct
	}
	return false
}

type NewChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RobotId string `protobuf:"bytes,1,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"` // 对话机器人ID，为空表示匹配
}

func (x *NewChatReq) Reset() {
	*x = NewChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewChatReq) ProtoMessage() {}

func (x *NewChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewChatReq.ProtoReflect.Descriptor instead.
func (*NewChatReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *NewChatReq) GetRobotId() string {
	if x != nil {
		return x.RobotId
	}
	return ""
}

type NewChatRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *NewChatRsp) Reset() {
	*x = NewChatRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewChatRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewChatRsp) ProtoMessage() {}

func (x *NewChatRsp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewChatRsp.ProtoReflect.Descriptor instead.
func (*NewChatRsp) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *NewChatRsp) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type WaitMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *WaitMessageReq) Reset() {
	*x = WaitMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaitMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitMessageReq) ProtoMessage() {}

func (x *WaitMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitMessageReq.ProtoReflect.Descriptor instead.
func (*WaitMessageReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *WaitMessageReq) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type WaitMessageRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *WaitMessageRsp) Reset() {
	*x = WaitMessageRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaitMessageRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitMessageRsp) ProtoMessage() {}

func (x *WaitMessageRsp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitMessageRsp.ProtoReflect.Descriptor instead.
func (*WaitMessageRsp) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *WaitMessageRsp) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type SendMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SendMessageReq) Reset() {
	*x = SendMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReq) ProtoMessage() {}

func (x *SendMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReq.ProtoReflect.Descriptor instead.
func (*SendMessageReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{6}
}

func (x *SendMessageReq) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *SendMessageReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SendMessageRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SendMessageRsp) Reset() {
	*x = SendMessageRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRsp) ProtoMessage() {}

func (x *SendMessageRsp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRsp.ProtoReflect.Descriptor instead.
func (*SendMessageRsp) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{7}
}

func (x *SendMessageRsp) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type GuessReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId      string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	PartnerType uint32 `protobuf:"varint,2,opt,name=partner_type,json=partnerType,proto3" json:"partner_type,omitempty"` // UserType
}

func (x *GuessReq) Reset() {
	*x = GuessReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuessReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuessReq) ProtoMessage() {}

func (x *GuessReq) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuessReq.ProtoReflect.Descriptor instead.
func (*GuessReq) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{8}
}

func (x *GuessReq) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *GuessReq) GetPartnerType() uint32 {
	if x != nil {
		return x.PartnerType
	}
	return 0
}

type GuessRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartnerType uint32 `protobuf:"varint,1,opt,name=partner_type,json=partnerType,proto3" json:"partner_type,omitempty"` // UserType
	Correct     bool   `protobuf:"varint,2,opt,name=correct,proto3" json:"correct,omitempty"`
}

func (x *GuessRsp) Reset() {
	*x = GuessRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuessRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuessRsp) ProtoMessage() {}

func (x *GuessRsp) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuessRsp.ProtoReflect.Descriptor instead.
func (*GuessRsp) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{9}
}

func (x *GuessRsp) GetPartnerType() uint32 {
	if x != nil {
		return x.PartnerType
	}
	return 0
}

func (x *GuessRsp) GetCorrect() bool {
	if x != nil {
		return x.Correct
	}
	return false
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x75,
	0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x76, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x22, 0x60, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x22, 0xe5, 0x03, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68,
	0x61, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x6e, 0x75, 0x6d,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1c, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x6d, 0x79, 0x5f, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x4d, 0x79, 0x54, 0x75, 0x72, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x75, 0x72, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x74, 0x75, 0x72, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x37, 0x0a,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x76, 0x72, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x22, 0x27, 0x0a, 0x0a, 0x4e, 0x65,
	0x77, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x43, 0x68, 0x61, 0x74, 0x52, 0x73,
	0x70, 0x12, 0x2c, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x76, 0x72, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0x29, 0x0a, 0x0e, 0x57, 0x61, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x57, 0x61,
	0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x75, 0x72,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x76, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x3d, 0x0a, 0x0e, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x75, 0x72, 0x69,
	0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x76, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x46, 0x0a, 0x08, 0x47, 0x75, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x47, 0x0a, 0x08, 0x47, 0x75, 0x65, 0x73, 0x73, 0x52, 0x73, 0x70, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x2a, 0x34, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x48, 0x55, 0x4d, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x4f, 0x42, 0x4f, 0x54, 0x10, 0x01,
	0x32, 0xc7, 0x02, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x53, 0x76, 0x72, 0x12, 0x49, 0x0a, 0x07,
	0x4e, 0x65, 0x77, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1e, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x61, 0x2e, 0x73, 0x76, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4e, 0x65, 0x77,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x61, 0x2e, 0x73, 0x76, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4e, 0x65, 0x77,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x73, 0x70, 0x12, 0x55, 0x0a, 0x0b, 0x57, 0x61, 0x69, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x61, 0x2e, 0x73, 0x76, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x57, 0x61, 0x69, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x74, 0x75, 0x72,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x76, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x57, 0x61, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x73, 0x70, 0x12, 0x55,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x2e,
	0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x76, 0x72, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x22, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x76,
	0x72, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x05, 0x47, 0x75, 0x65, 0x73, 0x73, 0x12, 0x1c,
	0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x76, 0x72, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x74,
	0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x76, 0x72, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x73, 0x52, 0x73, 0x70, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x2d,
	0x65, 0x72, 0x61, 0x2f, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x76, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_chat_proto_goTypes = []interface{}{
	(UserType)(0),          // 0: turingera.svr.chat.UserType
	(*Message)(nil),        // 1: turingera.svr.chat.Message
	(*Info)(nil),           // 2: turingera.svr.chat.Info
	(*NewChatReq)(nil),     // 3: turingera.svr.chat.NewChatReq
	(*NewChatRsp)(nil),     // 4: turingera.svr.chat.NewChatRsp
	(*WaitMessageReq)(nil), // 5: turingera.svr.chat.WaitMessageReq
	(*WaitMessageRsp)(nil), // 6: turingera.svr.chat.WaitMessageRsp
	(*SendMessageReq)(nil), // 7: turingera.svr.chat.SendMessageReq
	(*SendMessageRsp)(nil), // 8: turingera.svr.chat.SendMessageRsp
	(*GuessReq)(nil),       // 9: turingera.svr.chat.GuessReq
	(*GuessRsp)(nil),       // 10: turingera.svr.chat.GuessRsp
}
var file_chat_proto_depIdxs = []int32{
	1,  // 0: turingera.svr.chat.Info.messages:type_name -> turingera.svr.chat.Message
	2,  // 1: turingera.svr.chat.NewChatRsp.info:type_name -> turingera.svr.chat.Info
	2,  // 2: turingera.svr.chat.WaitMessageRsp.info:type_name -> turingera.svr.chat.Info
	2,  // 3: turingera.svr.chat.SendMessageRsp.info:type_name -> turingera.svr.chat.Info
	3,  // 4: turingera.svr.chat.ChatSvr.NewChat:input_type -> turingera.svr.chat.NewChatReq
	5,  // 5: turingera.svr.chat.ChatSvr.WaitMessage:input_type -> turingera.svr.chat.WaitMessageReq
	7,  // 6: turingera.svr.chat.ChatSvr.SendMessage:input_type -> turingera.svr.chat.SendMessageReq
	9,  // 7: turingera.svr.chat.ChatSvr.Guess:input_type -> turingera.svr.chat.GuessReq
	4,  // 8: turingera.svr.chat.ChatSvr.NewChat:output_type -> turingera.svr.chat.NewChatRsp
	6,  // 9: turingera.svr.chat.ChatSvr.WaitMessage:output_type -> turingera.svr.chat.WaitMessageRsp
	8,  // 10: turingera.svr.chat.ChatSvr.SendMessage:output_type -> turingera.svr.chat.SendMessageRsp
	10, // 11: turingera.svr.chat.ChatSvr.Guess:output_type -> turingera.svr.chat.GuessRsp
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewChatReq); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewChatRsp); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaitMessageReq); i {
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
		file_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaitMessageRsp); i {
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
		file_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageReq); i {
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
		file_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRsp); i {
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
		file_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuessReq); i {
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
		file_chat_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuessRsp); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		EnumInfos:         file_chat_proto_enumTypes,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

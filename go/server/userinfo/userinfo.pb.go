// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: userinfo.proto

package userinfo

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

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string       `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Nick         string       `protobuf:"bytes,2,opt,name=nick,proto3" json:"nick,omitempty"`
	Avatar       string       `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	LoginType    uint32       `protobuf:"varint,4,opt,name=login_type,json=loginType,proto3" json:"login_type,omitempty"`
	PlatformUid  string       `protobuf:"bytes,5,opt,name=platform_uid,json=platformUid,proto3" json:"platform_uid,omitempty"`
	Email        string       `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	RegisterTime int64        `protobuf:"varint,7,opt,name=register_time,json=registerTime,proto3" json:"register_time,omitempty"`
	InviteCode   string       `protobuf:"bytes,8,opt,name=invite_code,json=inviteCode,proto3" json:"invite_code,omitempty"`
	Dynamic      *UserDynamic `protobuf:"bytes,100,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserInfo) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

func (x *UserInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfo) GetLoginType() uint32 {
	if x != nil {
		return x.LoginType
	}
	return 0
}

func (x *UserInfo) GetPlatformUid() string {
	if x != nil {
		return x.PlatformUid
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetRegisterTime() int64 {
	if x != nil {
		return x.RegisterTime
	}
	return 0
}

func (x *UserInfo) GetInviteCode() string {
	if x != nil {
		return x.InviteCode
	}
	return ""
}

func (x *UserInfo) GetDynamic() *UserDynamic {
	if x != nil {
		return x.Dynamic
	}
	return nil
}

type UserDynamic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score     int32  `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`                          // 积分
	GameTotal uint32 `protobuf:"varint,2,opt,name=game_total,json=gameTotal,proto3" json:"game_total,omitempty"` // 总场数
	GameWin   uint32 `protobuf:"varint,3,opt,name=game_win,json=gameWin,proto3" json:"game_win,omitempty"`       // 胜利数
}

func (x *UserDynamic) Reset() {
	*x = UserDynamic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDynamic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDynamic) ProtoMessage() {}

func (x *UserDynamic) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDynamic.ProtoReflect.Descriptor instead.
func (*UserDynamic) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{1}
}

func (x *UserDynamic) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *UserDynamic) GetGameTotal() uint32 {
	if x != nil {
		return x.GameTotal
	}
	return 0
}

func (x *UserDynamic) GetGameWin() uint32 {
	if x != nil {
		return x.GameWin
	}
	return 0
}

type InitUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitInfo *UserInfo `protobuf:"bytes,1,opt,name=init_info,json=initInfo,proto3" json:"init_info,omitempty"`
}

func (x *InitUserReq) Reset() {
	*x = InitUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitUserReq) ProtoMessage() {}

func (x *InitUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitUserReq.ProtoReflect.Descriptor instead.
func (*InitUserReq) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{2}
}

func (x *InitUserReq) GetInitInfo() *UserInfo {
	if x != nil {
		return x.InitInfo
	}
	return nil
}

type InitUserRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	EarlyAccessCode string `protobuf:"bytes,2,opt,name=early_access_code,json=earlyAccessCode,proto3" json:"early_access_code,omitempty"`
}

func (x *InitUserRsp) Reset() {
	*x = InitUserRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitUserRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitUserRsp) ProtoMessage() {}

func (x *InitUserRsp) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitUserRsp.ProtoReflect.Descriptor instead.
func (*InitUserRsp) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{3}
}

func (x *InitUserRsp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *InitUserRsp) GetEarlyAccessCode() string {
	if x != nil {
		return x.EarlyAccessCode
	}
	return ""
}

type GetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{4}
}

type GetUserInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserInfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserInfoRsp) Reset() {
	*x = GetUserInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRsp) ProtoMessage() {}

func (x *GetUserInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRsp.ProtoReflect.Descriptor instead.
func (*GetUserInfoRsp) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserInfoRsp) GetUser() *UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

type BatchGetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *BatchGetUserInfoReq) Reset() {
	*x = BatchGetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetUserInfoReq) ProtoMessage() {}

func (x *BatchGetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetUserInfoReq.ProtoReflect.Descriptor instead.
func (*BatchGetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{6}
}

func (x *BatchGetUserInfoReq) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type BatchGetUserInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserInfo `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *BatchGetUserInfoRsp) Reset() {
	*x = BatchGetUserInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetUserInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetUserInfoRsp) ProtoMessage() {}

func (x *BatchGetUserInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetUserInfoRsp.ProtoReflect.Descriptor instead.
func (*BatchGetUserInfoRsp) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{7}
}

func (x *BatchGetUserInfoRsp) GetUsers() []*UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

type UpdateDynamicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string       `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Dynamic *UserDynamic `protobuf:"bytes,2,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
}

func (x *UpdateDynamicReq) Reset() {
	*x = UpdateDynamicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDynamicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDynamicReq) ProtoMessage() {}

func (x *UpdateDynamicReq) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDynamicReq.ProtoReflect.Descriptor instead.
func (*UpdateDynamicReq) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateDynamicReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateDynamicReq) GetDynamic() *UserDynamic {
	if x != nil {
		return x.Dynamic
	}
	return nil
}

type UpdateDynamicRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDynamicRsp) Reset() {
	*x = UpdateDynamicRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDynamicRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDynamicRsp) ProtoMessage() {}

func (x *UpdateDynamicRsp) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDynamicRsp.ProtoReflect.Descriptor instead.
func (*UpdateDynamicRsp) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{9}
}

type UpdateEarlyAccessReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	EarlyAccessCode string `protobuf:"bytes,2,opt,name=early_access_code,json=earlyAccessCode,proto3" json:"early_access_code,omitempty"`
}

func (x *UpdateEarlyAccessReq) Reset() {
	*x = UpdateEarlyAccessReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEarlyAccessReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEarlyAccessReq) ProtoMessage() {}

func (x *UpdateEarlyAccessReq) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEarlyAccessReq.ProtoReflect.Descriptor instead.
func (*UpdateEarlyAccessReq) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateEarlyAccessReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateEarlyAccessReq) GetEarlyAccessCode() string {
	if x != nil {
		return x.EarlyAccessCode
	}
	return ""
}

type UpdateEarlyAccessRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateEarlyAccessRsp) Reset() {
	*x = UpdateEarlyAccessRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEarlyAccessRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEarlyAccessRsp) ProtoMessage() {}

func (x *UpdateEarlyAccessRsp) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEarlyAccessRsp.ProtoReflect.Descriptor instead.
func (*UpdateEarlyAccessRsp) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{11}
}

var File_userinfo_proto protoreflect.FileDescriptor

var file_userinfo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xaf, 0x02, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x69, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x55, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74,
	0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x52, 0x07, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x22, 0x5d, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x22, 0x4f, 0x0a, 0x0b,
	0x49, 0x6e, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x40, 0x0a, 0x09, 0x69,
	0x6e, 0x69, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x69, 0x6e, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x52, 0x0a,
	0x0b, 0x49, 0x6e, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x22, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x30,
	0x0a, 0x13, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x22, 0x50, 0x0a, 0x13, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x6d, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x40, 0x0a, 0x07, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x52, 0x07, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x52, 0x73, 0x70, 0x22, 0x5b, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x61, 0x72, 0x6c, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x61, 0x72, 0x6c,
	0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x73, 0x70, 0x32, 0xa1, 0x04, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x5a, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x49, 0x6e, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x74, 0x75,
	0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x73, 0x70, 0x12, 0x63, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x29, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e,
	0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x72, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x2e, 0x74,
	0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x2e, 0x2e, 0x74,
	0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x69, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x12, 0x2b, 0x2e,
	0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x74, 0x75, 0x72,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x52, 0x73, 0x70, 0x12, 0x75, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x61, 0x72, 0x6c, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x2e, 0x74,
	0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x61, 0x72, 0x6c, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e,
	0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x61, 0x72, 0x6c, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x73, 0x70, 0x42, 0x3a,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75, 0x72,
	0x69, 0x6e, 0x67, 0x2d, 0x65, 0x72, 0x61, 0x2f, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_userinfo_proto_rawDescOnce sync.Once
	file_userinfo_proto_rawDescData = file_userinfo_proto_rawDesc
)

func file_userinfo_proto_rawDescGZIP() []byte {
	file_userinfo_proto_rawDescOnce.Do(func() {
		file_userinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_userinfo_proto_rawDescData)
	})
	return file_userinfo_proto_rawDescData
}

var file_userinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_userinfo_proto_goTypes = []interface{}{
	(*UserInfo)(nil),             // 0: turingera.server.userinfo.UserInfo
	(*UserDynamic)(nil),          // 1: turingera.server.userinfo.UserDynamic
	(*InitUserReq)(nil),          // 2: turingera.server.userinfo.InitUserReq
	(*InitUserRsp)(nil),          // 3: turingera.server.userinfo.InitUserRsp
	(*GetUserInfoReq)(nil),       // 4: turingera.server.userinfo.GetUserInfoReq
	(*GetUserInfoRsp)(nil),       // 5: turingera.server.userinfo.GetUserInfoRsp
	(*BatchGetUserInfoReq)(nil),  // 6: turingera.server.userinfo.BatchGetUserInfoReq
	(*BatchGetUserInfoRsp)(nil),  // 7: turingera.server.userinfo.BatchGetUserInfoRsp
	(*UpdateDynamicReq)(nil),     // 8: turingera.server.userinfo.UpdateDynamicReq
	(*UpdateDynamicRsp)(nil),     // 9: turingera.server.userinfo.UpdateDynamicRsp
	(*UpdateEarlyAccessReq)(nil), // 10: turingera.server.userinfo.UpdateEarlyAccessReq
	(*UpdateEarlyAccessRsp)(nil), // 11: turingera.server.userinfo.UpdateEarlyAccessRsp
}
var file_userinfo_proto_depIdxs = []int32{
	1,  // 0: turingera.server.userinfo.UserInfo.dynamic:type_name -> turingera.server.userinfo.UserDynamic
	0,  // 1: turingera.server.userinfo.InitUserReq.init_info:type_name -> turingera.server.userinfo.UserInfo
	0,  // 2: turingera.server.userinfo.GetUserInfoRsp.user:type_name -> turingera.server.userinfo.UserInfo
	0,  // 3: turingera.server.userinfo.BatchGetUserInfoRsp.users:type_name -> turingera.server.userinfo.UserInfo
	1,  // 4: turingera.server.userinfo.UpdateDynamicReq.dynamic:type_name -> turingera.server.userinfo.UserDynamic
	2,  // 5: turingera.server.userinfo.Userinfo.InitUser:input_type -> turingera.server.userinfo.InitUserReq
	4,  // 6: turingera.server.userinfo.Userinfo.GetUserInfo:input_type -> turingera.server.userinfo.GetUserInfoReq
	6,  // 7: turingera.server.userinfo.Userinfo.BatchGetUserInfo:input_type -> turingera.server.userinfo.BatchGetUserInfoReq
	8,  // 8: turingera.server.userinfo.Userinfo.UpdateDynamic:input_type -> turingera.server.userinfo.UpdateDynamicReq
	10, // 9: turingera.server.userinfo.Userinfo.UpdateEarlyAccess:input_type -> turingera.server.userinfo.UpdateEarlyAccessReq
	3,  // 10: turingera.server.userinfo.Userinfo.InitUser:output_type -> turingera.server.userinfo.InitUserRsp
	5,  // 11: turingera.server.userinfo.Userinfo.GetUserInfo:output_type -> turingera.server.userinfo.GetUserInfoRsp
	7,  // 12: turingera.server.userinfo.Userinfo.BatchGetUserInfo:output_type -> turingera.server.userinfo.BatchGetUserInfoRsp
	9,  // 13: turingera.server.userinfo.Userinfo.UpdateDynamic:output_type -> turingera.server.userinfo.UpdateDynamicRsp
	11, // 14: turingera.server.userinfo.Userinfo.UpdateEarlyAccess:output_type -> turingera.server.userinfo.UpdateEarlyAccessRsp
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_userinfo_proto_init() }
func file_userinfo_proto_init() {
	if File_userinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_userinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDynamic); i {
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
		file_userinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitUserReq); i {
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
		file_userinfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitUserRsp); i {
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
		file_userinfo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReq); i {
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
		file_userinfo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoRsp); i {
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
		file_userinfo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetUserInfoReq); i {
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
		file_userinfo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetUserInfoRsp); i {
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
		file_userinfo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDynamicReq); i {
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
		file_userinfo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDynamicRsp); i {
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
		file_userinfo_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEarlyAccessReq); i {
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
		file_userinfo_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEarlyAccessRsp); i {
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
			RawDescriptor: file_userinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userinfo_proto_goTypes,
		DependencyIndexes: file_userinfo_proto_depIdxs,
		MessageInfos:      file_userinfo_proto_msgTypes,
	}.Build()
	File_userinfo_proto = out.File
	file_userinfo_proto_rawDesc = nil
	file_userinfo_proto_goTypes = nil
	file_userinfo_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: auth.proto

package auth

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

type LoginType int32

const (
	LoginType_LOGIN_TYPE_UNKNOWN LoginType = 0
	LoginType_LOGIN_TYPE_GOOGLE  LoginType = 1 // google
	LoginType_LOGIN_TYPE_TWITTER LoginType = 2 // twitter
	LoginType_LOGIN_TYPE_WALLET  LoginType = 3 // 钱包
)

// Enum value maps for LoginType.
var (
	LoginType_name = map[int32]string{
		0: "LOGIN_TYPE_UNKNOWN",
		1: "LOGIN_TYPE_GOOGLE",
		2: "LOGIN_TYPE_TWITTER",
		3: "LOGIN_TYPE_WALLET",
	}
	LoginType_value = map[string]int32{
		"LOGIN_TYPE_UNKNOWN": 0,
		"LOGIN_TYPE_GOOGLE":  1,
		"LOGIN_TYPE_TWITTER": 2,
		"LOGIN_TYPE_WALLET":  3,
	}
)

func (x LoginType) Enum() *LoginType {
	p := new(LoginType)
	*p = x
	return p
}

func (x LoginType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginType) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[0].Descriptor()
}

func (LoginType) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[0]
}

func (x LoginType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginType.Descriptor instead.
func (LoginType) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

type LoginResFlag int32

const (
	LoginResFlag_LOGIN_SUCC                LoginResFlag = 0
	LoginResFlag_LOGIN_ACCESS_CODE_EMPTY   LoginResFlag = 1 // 需要填写邀请码
	LoginResFlag_LOGIN_ACCESS_CODE_USED    LoginResFlag = 2 // 邀请码已被使用
	LoginResFlag_LOGIN_ACCESS_CODE_INVALID LoginResFlag = 3 // 邀请码错误
)

// Enum value maps for LoginResFlag.
var (
	LoginResFlag_name = map[int32]string{
		0: "LOGIN_SUCC",
		1: "LOGIN_ACCESS_CODE_EMPTY",
		2: "LOGIN_ACCESS_CODE_USED",
		3: "LOGIN_ACCESS_CODE_INVALID",
	}
	LoginResFlag_value = map[string]int32{
		"LOGIN_SUCC":                0,
		"LOGIN_ACCESS_CODE_EMPTY":   1,
		"LOGIN_ACCESS_CODE_USED":    2,
		"LOGIN_ACCESS_CODE_INVALID": 3,
	}
)

func (x LoginResFlag) Enum() *LoginResFlag {
	p := new(LoginResFlag)
	*p = x
	return p
}

func (x LoginResFlag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginResFlag) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[1].Descriptor()
}

func (LoginResFlag) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[1]
}

func (x LoginResFlag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginResFlag.Descriptor instead.
func (LoginResFlag) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

type NonceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NonceReq) Reset() {
	*x = NonceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonceReq) ProtoMessage() {}

func (x *NonceReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonceReq.ProtoReflect.Descriptor instead.
func (*NonceReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

type NonceRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *NonceRsp) Reset() {
	*x = NonceRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonceRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonceRsp) ProtoMessage() {}

func (x *NonceRsp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonceRsp.ProtoReflect.Descriptor instead.
func (*NonceRsp) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *NonceRsp) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type EarlyAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EarlyAccessCode string `protobuf:"bytes,1,opt,name=early_access_code,json=earlyAccessCode,proto3" json:"early_access_code,omitempty"`
	UserId          string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *EarlyAccess) Reset() {
	*x = EarlyAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EarlyAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EarlyAccess) ProtoMessage() {}

func (x *EarlyAccess) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EarlyAccess.ProtoReflect.Descriptor instead.
func (*EarlyAccess) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *EarlyAccess) GetEarlyAccessCode() string {
	if x != nil {
		return x.EarlyAccessCode
	}
	return ""
}

func (x *EarlyAccess) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Secret      string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	// 钱包登录使用
	Message         string `protobuf:"bytes,10,opt,name=message,proto3" json:"message,omitempty"`
	Signature       string `protobuf:"bytes,11,opt,name=signature,proto3" json:"signature,omitempty"`
	LoginType       uint32 `protobuf:"varint,50,opt,name=login_type,json=loginType,proto3" json:"login_type,omitempty"`
	EarlyAccessCode string `protobuf:"bytes,51,opt,name=early_access_code,json=earlyAccessCode,proto3" json:"early_access_code,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *LoginReq) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginReq) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *LoginReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginReq) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *LoginReq) GetLoginType() uint32 {
	if x != nil {
		return x.LoginType
	}
	return 0
}

func (x *LoginReq) GetEarlyAccessCode() string {
	if x != nil {
		return x.EarlyAccessCode
	}
	return ""
}

type LoginRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Expire  uint32 `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"`
	ResFlag uint32 `protobuf:"varint,4,opt,name=res_flag,json=resFlag,proto3" json:"res_flag,omitempty"`
}

func (x *LoginRsp) Reset() {
	*x = LoginRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRsp) ProtoMessage() {}

func (x *LoginRsp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRsp.ProtoReflect.Descriptor instead.
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRsp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LoginRsp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginRsp) GetExpire() uint32 {
	if x != nil {
		return x.Expire
	}
	return 0
}

func (x *LoginRsp) GetResFlag() uint32 {
	if x != nil {
		return x.ResFlag
	}
	return 0
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x75,
	0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x22, 0x0a, 0x0a, 0x08, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x22,
	0x20, 0x0a, 0x08, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x22, 0x52, 0x0a, 0x0b, 0x45, 0x61, 0x72, 0x6c, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x2a, 0x0a, 0x11, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x61, 0x72,
	0x6c, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xc8, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x6c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x73, 0x46, 0x6c, 0x61, 0x67, 0x2a, 0x69,
	0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x4c,
	0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x4f,
	0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x57, 0x49, 0x54, 0x54, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x57, 0x41, 0x4c, 0x4c, 0x45, 0x54, 0x10, 0x03, 0x2a, 0x76, 0x0a, 0x0c, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x47,
	0x49, 0x4e, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x4f, 0x47,
	0x49, 0x4e, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45,
	0x4d, 0x50, 0x54, 0x59, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x41, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x03, 0x32, 0x9c, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x49, 0x0a, 0x05, 0x4e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4e, 0x6f, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4e, 0x6f, 0x6e,
	0x63, 0x65, 0x52, 0x73, 0x70, 0x12, 0x49, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1f,
	0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x1f, 0x2e, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x75, 0x72, 0x69, 0x6e, 0x67, 0x2d, 0x65, 0x72, 0x61, 0x2f, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_auth_proto_goTypes = []interface{}{
	(LoginType)(0),      // 0: turingera.server.auth.LoginType
	(LoginResFlag)(0),   // 1: turingera.server.auth.LoginResFlag
	(*NonceReq)(nil),    // 2: turingera.server.auth.NonceReq
	(*NonceRsp)(nil),    // 3: turingera.server.auth.NonceRsp
	(*EarlyAccess)(nil), // 4: turingera.server.auth.EarlyAccess
	(*LoginReq)(nil),    // 5: turingera.server.auth.LoginReq
	(*LoginRsp)(nil),    // 6: turingera.server.auth.LoginRsp
}
var file_auth_proto_depIdxs = []int32{
	2, // 0: turingera.server.auth.Auth.Nonce:input_type -> turingera.server.auth.NonceReq
	5, // 1: turingera.server.auth.Auth.Login:input_type -> turingera.server.auth.LoginReq
	3, // 2: turingera.server.auth.Auth.Nonce:output_type -> turingera.server.auth.NonceRsp
	6, // 3: turingera.server.auth.Auth.Login:output_type -> turingera.server.auth.LoginRsp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonceReq); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonceRsp); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EarlyAccess); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRsp); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		EnumInfos:         file_auth_proto_enumTypes,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}

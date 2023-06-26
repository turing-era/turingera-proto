// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: chat.proto

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatSvrClient is the client API for ChatSvr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatSvrClient interface {
	// 新建对话
	NewChat(ctx context.Context, in *NewChatReq, opts ...grpc.CallOption) (*NewChatRsp, error)
	// 等待消息
	WaitMessage(ctx context.Context, in *WaitMessageReq, opts ...grpc.CallOption) (*WaitMessageRsp, error)
	// 发送消息
	SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageRsp, error)
	// 做出判断
	Guess(ctx context.Context, in *GuessReq, opts ...grpc.CallOption) (*GuessRsp, error)
	// 等待判断
	WaitGuess(ctx context.Context, in *WaitGuessReq, opts ...grpc.CallOption) (*WaitGuessRsp, error)
	// 分享
	Share(ctx context.Context, in *ShareReq, opts ...grpc.CallOption) (*ShareRsp, error)
}

type chatSvrClient struct {
	cc grpc.ClientConnInterface
}

func NewChatSvrClient(cc grpc.ClientConnInterface) ChatSvrClient {
	return &chatSvrClient{cc}
}

func (c *chatSvrClient) NewChat(ctx context.Context, in *NewChatReq, opts ...grpc.CallOption) (*NewChatRsp, error) {
	out := new(NewChatRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.chat.ChatSvr/NewChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatSvrClient) WaitMessage(ctx context.Context, in *WaitMessageReq, opts ...grpc.CallOption) (*WaitMessageRsp, error) {
	out := new(WaitMessageRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.chat.ChatSvr/WaitMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatSvrClient) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageRsp, error) {
	out := new(SendMessageRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.chat.ChatSvr/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatSvrClient) Guess(ctx context.Context, in *GuessReq, opts ...grpc.CallOption) (*GuessRsp, error) {
	out := new(GuessRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.chat.ChatSvr/Guess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatSvrClient) WaitGuess(ctx context.Context, in *WaitGuessReq, opts ...grpc.CallOption) (*WaitGuessRsp, error) {
	out := new(WaitGuessRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.chat.ChatSvr/WaitGuess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatSvrClient) Share(ctx context.Context, in *ShareReq, opts ...grpc.CallOption) (*ShareRsp, error) {
	out := new(ShareRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.chat.ChatSvr/Share", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatSvrServer is the server API for ChatSvr service.
// All implementations must embed UnimplementedChatSvrServer
// for forward compatibility
type ChatSvrServer interface {
	// 新建对话
	NewChat(context.Context, *NewChatReq) (*NewChatRsp, error)
	// 等待消息
	WaitMessage(context.Context, *WaitMessageReq) (*WaitMessageRsp, error)
	// 发送消息
	SendMessage(context.Context, *SendMessageReq) (*SendMessageRsp, error)
	// 做出判断
	Guess(context.Context, *GuessReq) (*GuessRsp, error)
	// 等待判断
	WaitGuess(context.Context, *WaitGuessReq) (*WaitGuessRsp, error)
	// 分享
	Share(context.Context, *ShareReq) (*ShareRsp, error)
	mustEmbedUnimplementedChatSvrServer()
}

// UnimplementedChatSvrServer must be embedded to have forward compatible implementations.
type UnimplementedChatSvrServer struct {
}

func (UnimplementedChatSvrServer) NewChat(context.Context, *NewChatReq) (*NewChatRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewChat not implemented")
}
func (UnimplementedChatSvrServer) WaitMessage(context.Context, *WaitMessageReq) (*WaitMessageRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitMessage not implemented")
}
func (UnimplementedChatSvrServer) SendMessage(context.Context, *SendMessageReq) (*SendMessageRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatSvrServer) Guess(context.Context, *GuessReq) (*GuessRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Guess not implemented")
}
func (UnimplementedChatSvrServer) WaitGuess(context.Context, *WaitGuessReq) (*WaitGuessRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitGuess not implemented")
}
func (UnimplementedChatSvrServer) Share(context.Context, *ShareReq) (*ShareRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Share not implemented")
}
func (UnimplementedChatSvrServer) mustEmbedUnimplementedChatSvrServer() {}

// UnsafeChatSvrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatSvrServer will
// result in compilation errors.
type UnsafeChatSvrServer interface {
	mustEmbedUnimplementedChatSvrServer()
}

func RegisterChatSvrServer(s grpc.ServiceRegistrar, srv ChatSvrServer) {
	s.RegisterService(&ChatSvr_ServiceDesc, srv)
}

func _ChatSvr_NewChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatSvrServer).NewChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.chat.ChatSvr/NewChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatSvrServer).NewChat(ctx, req.(*NewChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatSvr_WaitMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatSvrServer).WaitMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.chat.ChatSvr/WaitMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatSvrServer).WaitMessage(ctx, req.(*WaitMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatSvr_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatSvrServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.chat.ChatSvr/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatSvrServer).SendMessage(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatSvr_Guess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuessReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatSvrServer).Guess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.chat.ChatSvr/Guess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatSvrServer).Guess(ctx, req.(*GuessReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatSvr_WaitGuess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitGuessReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatSvrServer).WaitGuess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.chat.ChatSvr/WaitGuess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatSvrServer).WaitGuess(ctx, req.(*WaitGuessReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatSvr_Share_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatSvrServer).Share(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.chat.ChatSvr/Share",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatSvrServer).Share(ctx, req.(*ShareReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatSvr_ServiceDesc is the grpc.ServiceDesc for ChatSvr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatSvr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "turingera.server.chat.ChatSvr",
	HandlerType: (*ChatSvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewChat",
			Handler:    _ChatSvr_NewChat_Handler,
		},
		{
			MethodName: "WaitMessage",
			Handler:    _ChatSvr_WaitMessage_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatSvr_SendMessage_Handler,
		},
		{
			MethodName: "Guess",
			Handler:    _ChatSvr_Guess_Handler,
		},
		{
			MethodName: "WaitGuess",
			Handler:    _ChatSvr_WaitGuess_Handler,
		},
		{
			MethodName: "Share",
			Handler:    _ChatSvr_Share_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}

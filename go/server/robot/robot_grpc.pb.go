// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: robot.proto

package robot

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

// RobotSvrClient is the client API for RobotSvr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RobotSvrClient interface {
	// 获取性别列表
	GetGenderList(ctx context.Context, in *GetGenderListReq, opts ...grpc.CallOption) (*GetGenderListRsp, error)
	// 获取身份列表
	GetIdentityList(ctx context.Context, in *GetIdentityListReq, opts ...grpc.CallOption) (*GetIdentityListRsp, error)
	// 获取标签列表
	GetTagList(ctx context.Context, in *GetTagListReq, opts ...grpc.CallOption) (*GetTagListRsp, error)
	// 获取AI
	GetRobot(ctx context.Context, in *GetRobotReq, opts ...grpc.CallOption) (*GetRobotRsp, error)
	// 获取AI详情
	GetRobotDetail(ctx context.Context, in *GetRobotDetailReq, opts ...grpc.CallOption) (*GetRobotDetailRsp, error)
	// 新建AI
	NewRobot(ctx context.Context, in *NewRobotReq, opts ...grpc.CallOption) (*NewRobotRsp, error)
	// 编辑AI
	EditRobot(ctx context.Context, in *EditRobotReq, opts ...grpc.CallOption) (*EditRobotRsp, error)
	// 删除AI
	DeleteRobot(ctx context.Context, in *DeleteRobotReq, opts ...grpc.CallOption) (*DeleteRobotRsp, error)
	// 获取我的AI列表
	GetRobots(ctx context.Context, in *GetRobotsReq, opts ...grpc.CallOption) (*GetRobotsRsp, error)
	// 推荐AI列表
	RecommendRobots(ctx context.Context, in *RecommendRobotsReq, opts ...grpc.CallOption) (*RecommendRobotsRsp, error)
	// 随机AI
	RandomRobotID(ctx context.Context, in *RandomRobotIDReq, opts ...grpc.CallOption) (*RandomRobotIDRsp, error)
}

type robotSvrClient struct {
	cc grpc.ClientConnInterface
}

func NewRobotSvrClient(cc grpc.ClientConnInterface) RobotSvrClient {
	return &robotSvrClient{cc}
}

func (c *robotSvrClient) GetGenderList(ctx context.Context, in *GetGenderListReq, opts ...grpc.CallOption) (*GetGenderListRsp, error) {
	out := new(GetGenderListRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.robot.RobotSvr/GetGenderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotSvrClient) GetIdentityList(ctx context.Context, in *GetIdentityListReq, opts ...grpc.CallOption) (*GetIdentityListRsp, error) {
	out := new(GetIdentityListRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.robot.RobotSvr/GetIdentityList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotSvrClient) GetTagList(ctx context.Context, in *GetTagListReq, opts ...grpc.CallOption) (*GetTagListRsp, error) {
	out := new(GetTagListRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.robot.RobotSvr/GetTagList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotSvrClient) GetRobot(ctx context.Context, in *GetRobotReq, opts ...grpc.CallOption) (*GetRobotRsp, error) {
	out := new(GetRobotRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.robot.RobotSvr/GetRobot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotSvrClient) GetRobotDetail(ctx context.Context, in *GetRobotDetailReq, opts ...grpc.CallOption) (*GetRobotDetailRsp, error) {
	out := new(GetRobotDetailRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.robot.RobotSvr/GetRobotDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotSvrClient) NewRobot(ctx context.Context, in *NewRobotReq, opts ...grpc.CallOption) (*NewRobotRsp, error) {
	out := new(NewRobotRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.robot.RobotSvr/NewRobot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotSvrClient) EditRobot(ctx context.Context, in *EditRobotReq, opts ...grpc.CallOption) (*EditRobotRsp, error) {
	out := new(EditRobotRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.robot.RobotSvr/EditRobot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotSvrClient) DeleteRobot(ctx context.Context, in *DeleteRobotReq, opts ...grpc.CallOption) (*DeleteRobotRsp, error) {
	out := new(DeleteRobotRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.robot.RobotSvr/DeleteRobot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotSvrClient) GetRobots(ctx context.Context, in *GetRobotsReq, opts ...grpc.CallOption) (*GetRobotsRsp, error) {
	out := new(GetRobotsRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.robot.RobotSvr/GetRobots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotSvrClient) RecommendRobots(ctx context.Context, in *RecommendRobotsReq, opts ...grpc.CallOption) (*RecommendRobotsRsp, error) {
	out := new(RecommendRobotsRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.robot.RobotSvr/RecommendRobots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotSvrClient) RandomRobotID(ctx context.Context, in *RandomRobotIDReq, opts ...grpc.CallOption) (*RandomRobotIDRsp, error) {
	out := new(RandomRobotIDRsp)
	err := c.cc.Invoke(ctx, "/turingera.server.robot.RobotSvr/RandomRobotID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotSvrServer is the server API for RobotSvr service.
// All implementations must embed UnimplementedRobotSvrServer
// for forward compatibility
type RobotSvrServer interface {
	// 获取性别列表
	GetGenderList(context.Context, *GetGenderListReq) (*GetGenderListRsp, error)
	// 获取身份列表
	GetIdentityList(context.Context, *GetIdentityListReq) (*GetIdentityListRsp, error)
	// 获取标签列表
	GetTagList(context.Context, *GetTagListReq) (*GetTagListRsp, error)
	// 获取AI
	GetRobot(context.Context, *GetRobotReq) (*GetRobotRsp, error)
	// 获取AI详情
	GetRobotDetail(context.Context, *GetRobotDetailReq) (*GetRobotDetailRsp, error)
	// 新建AI
	NewRobot(context.Context, *NewRobotReq) (*NewRobotRsp, error)
	// 编辑AI
	EditRobot(context.Context, *EditRobotReq) (*EditRobotRsp, error)
	// 删除AI
	DeleteRobot(context.Context, *DeleteRobotReq) (*DeleteRobotRsp, error)
	// 获取我的AI列表
	GetRobots(context.Context, *GetRobotsReq) (*GetRobotsRsp, error)
	// 推荐AI列表
	RecommendRobots(context.Context, *RecommendRobotsReq) (*RecommendRobotsRsp, error)
	// 随机AI
	RandomRobotID(context.Context, *RandomRobotIDReq) (*RandomRobotIDRsp, error)
	mustEmbedUnimplementedRobotSvrServer()
}

// UnimplementedRobotSvrServer must be embedded to have forward compatible implementations.
type UnimplementedRobotSvrServer struct {
}

func (UnimplementedRobotSvrServer) GetGenderList(context.Context, *GetGenderListReq) (*GetGenderListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenderList not implemented")
}
func (UnimplementedRobotSvrServer) GetIdentityList(context.Context, *GetIdentityListReq) (*GetIdentityListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentityList not implemented")
}
func (UnimplementedRobotSvrServer) GetTagList(context.Context, *GetTagListReq) (*GetTagListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagList not implemented")
}
func (UnimplementedRobotSvrServer) GetRobot(context.Context, *GetRobotReq) (*GetRobotRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobot not implemented")
}
func (UnimplementedRobotSvrServer) GetRobotDetail(context.Context, *GetRobotDetailReq) (*GetRobotDetailRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobotDetail not implemented")
}
func (UnimplementedRobotSvrServer) NewRobot(context.Context, *NewRobotReq) (*NewRobotRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewRobot not implemented")
}
func (UnimplementedRobotSvrServer) EditRobot(context.Context, *EditRobotReq) (*EditRobotRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditRobot not implemented")
}
func (UnimplementedRobotSvrServer) DeleteRobot(context.Context, *DeleteRobotReq) (*DeleteRobotRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRobot not implemented")
}
func (UnimplementedRobotSvrServer) GetRobots(context.Context, *GetRobotsReq) (*GetRobotsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobots not implemented")
}
func (UnimplementedRobotSvrServer) RecommendRobots(context.Context, *RecommendRobotsReq) (*RecommendRobotsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendRobots not implemented")
}
func (UnimplementedRobotSvrServer) RandomRobotID(context.Context, *RandomRobotIDReq) (*RandomRobotIDRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RandomRobotID not implemented")
}
func (UnimplementedRobotSvrServer) mustEmbedUnimplementedRobotSvrServer() {}

// UnsafeRobotSvrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RobotSvrServer will
// result in compilation errors.
type UnsafeRobotSvrServer interface {
	mustEmbedUnimplementedRobotSvrServer()
}

func RegisterRobotSvrServer(s grpc.ServiceRegistrar, srv RobotSvrServer) {
	s.RegisterService(&RobotSvr_ServiceDesc, srv)
}

func _RobotSvr_GetGenderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGenderListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotSvrServer).GetGenderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.robot.RobotSvr/GetGenderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotSvrServer).GetGenderList(ctx, req.(*GetGenderListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotSvr_GetIdentityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdentityListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotSvrServer).GetIdentityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.robot.RobotSvr/GetIdentityList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotSvrServer).GetIdentityList(ctx, req.(*GetIdentityListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotSvr_GetTagList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotSvrServer).GetTagList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.robot.RobotSvr/GetTagList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotSvrServer).GetTagList(ctx, req.(*GetTagListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotSvr_GetRobot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRobotReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotSvrServer).GetRobot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.robot.RobotSvr/GetRobot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotSvrServer).GetRobot(ctx, req.(*GetRobotReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotSvr_GetRobotDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRobotDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotSvrServer).GetRobotDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.robot.RobotSvr/GetRobotDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotSvrServer).GetRobotDetail(ctx, req.(*GetRobotDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotSvr_NewRobot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRobotReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotSvrServer).NewRobot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.robot.RobotSvr/NewRobot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotSvrServer).NewRobot(ctx, req.(*NewRobotReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotSvr_EditRobot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditRobotReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotSvrServer).EditRobot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.robot.RobotSvr/EditRobot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotSvrServer).EditRobot(ctx, req.(*EditRobotReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotSvr_DeleteRobot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRobotReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotSvrServer).DeleteRobot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.robot.RobotSvr/DeleteRobot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotSvrServer).DeleteRobot(ctx, req.(*DeleteRobotReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotSvr_GetRobots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRobotsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotSvrServer).GetRobots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.robot.RobotSvr/GetRobots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotSvrServer).GetRobots(ctx, req.(*GetRobotsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotSvr_RecommendRobots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendRobotsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotSvrServer).RecommendRobots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.robot.RobotSvr/RecommendRobots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotSvrServer).RecommendRobots(ctx, req.(*RecommendRobotsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotSvr_RandomRobotID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomRobotIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotSvrServer).RandomRobotID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turingera.server.robot.RobotSvr/RandomRobotID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotSvrServer).RandomRobotID(ctx, req.(*RandomRobotIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RobotSvr_ServiceDesc is the grpc.ServiceDesc for RobotSvr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RobotSvr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "turingera.server.robot.RobotSvr",
	HandlerType: (*RobotSvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGenderList",
			Handler:    _RobotSvr_GetGenderList_Handler,
		},
		{
			MethodName: "GetIdentityList",
			Handler:    _RobotSvr_GetIdentityList_Handler,
		},
		{
			MethodName: "GetTagList",
			Handler:    _RobotSvr_GetTagList_Handler,
		},
		{
			MethodName: "GetRobot",
			Handler:    _RobotSvr_GetRobot_Handler,
		},
		{
			MethodName: "GetRobotDetail",
			Handler:    _RobotSvr_GetRobotDetail_Handler,
		},
		{
			MethodName: "NewRobot",
			Handler:    _RobotSvr_NewRobot_Handler,
		},
		{
			MethodName: "EditRobot",
			Handler:    _RobotSvr_EditRobot_Handler,
		},
		{
			MethodName: "DeleteRobot",
			Handler:    _RobotSvr_DeleteRobot_Handler,
		},
		{
			MethodName: "GetRobots",
			Handler:    _RobotSvr_GetRobots_Handler,
		},
		{
			MethodName: "RecommendRobots",
			Handler:    _RobotSvr_RecommendRobots_Handler,
		},
		{
			MethodName: "RandomRobotID",
			Handler:    _RobotSvr_RandomRobotID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "robot.proto",
}

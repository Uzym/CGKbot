// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/api.proto

package api

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

// BotApiClient is the client API for BotApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotApiClient interface {
	SetNewUser(ctx context.Context, in *NewUserWrite, opts ...grpc.CallOption) (*NewUserRead, error)
	GetQuestion(ctx context.Context, in *QuestionWrite, opts ...grpc.CallOption) (*QuestionRead, error)
	PostAnswer(ctx context.Context, in *AnswerWrite, opts ...grpc.CallOption) (*AnswerRead, error)
	GetLeaderboard(ctx context.Context, in *LeaderboardWrite, opts ...grpc.CallOption) (*LeaderboardRead, error)
}

type botApiClient struct {
	cc grpc.ClientConnInterface
}

func NewBotApiClient(cc grpc.ClientConnInterface) BotApiClient {
	return &botApiClient{cc}
}

func (c *botApiClient) SetNewUser(ctx context.Context, in *NewUserWrite, opts ...grpc.CallOption) (*NewUserRead, error) {
	out := new(NewUserRead)
	err := c.cc.Invoke(ctx, "/api.BotApi/SetNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botApiClient) GetQuestion(ctx context.Context, in *QuestionWrite, opts ...grpc.CallOption) (*QuestionRead, error) {
	out := new(QuestionRead)
	err := c.cc.Invoke(ctx, "/api.BotApi/GetQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botApiClient) PostAnswer(ctx context.Context, in *AnswerWrite, opts ...grpc.CallOption) (*AnswerRead, error) {
	out := new(AnswerRead)
	err := c.cc.Invoke(ctx, "/api.BotApi/PostAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botApiClient) GetLeaderboard(ctx context.Context, in *LeaderboardWrite, opts ...grpc.CallOption) (*LeaderboardRead, error) {
	out := new(LeaderboardRead)
	err := c.cc.Invoke(ctx, "/api.BotApi/GetLeaderboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotApiServer is the server API for BotApi service.
// All implementations must embed UnimplementedBotApiServer
// for forward compatibility
type BotApiServer interface {
	SetNewUser(context.Context, *NewUserWrite) (*NewUserRead, error)
	GetQuestion(context.Context, *QuestionWrite) (*QuestionRead, error)
	PostAnswer(context.Context, *AnswerWrite) (*AnswerRead, error)
	GetLeaderboard(context.Context, *LeaderboardWrite) (*LeaderboardRead, error)
	mustEmbedUnimplementedBotApiServer()
}

// UnimplementedBotApiServer must be embedded to have forward compatible implementations.
type UnimplementedBotApiServer struct {
}

func (UnimplementedBotApiServer) SetNewUser(context.Context, *NewUserWrite) (*NewUserRead, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNewUser not implemented")
}
func (UnimplementedBotApiServer) GetQuestion(context.Context, *QuestionWrite) (*QuestionRead, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestion not implemented")
}
func (UnimplementedBotApiServer) PostAnswer(context.Context, *AnswerWrite) (*AnswerRead, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostAnswer not implemented")
}
func (UnimplementedBotApiServer) GetLeaderboard(context.Context, *LeaderboardWrite) (*LeaderboardRead, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderboard not implemented")
}
func (UnimplementedBotApiServer) mustEmbedUnimplementedBotApiServer() {}

// UnsafeBotApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotApiServer will
// result in compilation errors.
type UnsafeBotApiServer interface {
	mustEmbedUnimplementedBotApiServer()
}

func RegisterBotApiServer(s grpc.ServiceRegistrar, srv BotApiServer) {
	s.RegisterService(&BotApi_ServiceDesc, srv)
}

func _BotApi_SetNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUserWrite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotApiServer).SetNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotApi/SetNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotApiServer).SetNewUser(ctx, req.(*NewUserWrite))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotApi_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionWrite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotApiServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotApi/GetQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotApiServer).GetQuestion(ctx, req.(*QuestionWrite))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotApi_PostAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerWrite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotApiServer).PostAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotApi/PostAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotApiServer).PostAnswer(ctx, req.(*AnswerWrite))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotApi_GetLeaderboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardWrite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotApiServer).GetLeaderboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotApi/GetLeaderboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotApiServer).GetLeaderboard(ctx, req.(*LeaderboardWrite))
	}
	return interceptor(ctx, in, info, handler)
}

// BotApi_ServiceDesc is the grpc.ServiceDesc for BotApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BotApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.BotApi",
	HandlerType: (*BotApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetNewUser",
			Handler:    _BotApi_SetNewUser_Handler,
		},
		{
			MethodName: "GetQuestion",
			Handler:    _BotApi_GetQuestion_Handler,
		},
		{
			MethodName: "PostAnswer",
			Handler:    _BotApi_PostAnswer_Handler,
		},
		{
			MethodName: "GetLeaderboard",
			Handler:    _BotApi_GetLeaderboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}

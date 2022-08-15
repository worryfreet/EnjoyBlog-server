// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"EnjoyBlog/app/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserInfoReq  = user.UserInfoReq
	UserInfoResp = user.UserInfoResp
	UserListReq  = user.UserListReq
	UserListResp = user.UserListResp

	User interface {
		GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		GetUserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		GetAll(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (user.User_GetAllClient, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUser) GetUserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserList(ctx, in, opts...)
}

func (m *defaultUser) GetAll(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (user.User_GetAllClient, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetAll(ctx, in, opts...)
}
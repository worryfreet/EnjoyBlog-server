package svc

import (
	"EnjoyBlog/app/admin/api/internal/config"
	"EnjoyBlog/app/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRPC user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: user.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}

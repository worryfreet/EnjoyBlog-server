package svc

import (
	"EnjoyBlog/app/user/api/internal/config"
	"EnjoyBlog/app/user/api/internal/middleware"
	"EnjoyBlog/app/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	AdminAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		AdminAuth: middleware.NewAdminAuthMiddleware().Handle,
	}
}

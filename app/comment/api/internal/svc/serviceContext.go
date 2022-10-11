package svc

import (
	"EnjoyBlog/app/comment/api/internal/config"
	"EnjoyBlog/app/comment/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	model.GlobalComment = model.NewCommentModel(mysqlConn, c.CacheRedis)
	return &ServiceContext{
		Config:       c,
		CommentModel: model.GlobalComment,
	}
}

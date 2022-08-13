package svc

import (
	"EnjoyBlog/app/article/api/internal/config"
	"EnjoyBlog/app/article/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config               config.Config
	ArticleModel         model.ArticleModel
	ArticleGroupModel    model.ArticleGroupModel
	ArticleGroupRelModel model.ArticleGroupRelModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	model.GlobalArticleGroupRel = model.NewArticleGroupRelModel(mysqlConn, c.CacheRedis)
	return &ServiceContext{
		Config:               c,
		ArticleModel:         model.NewArticleModel(mysqlConn, c.CacheRedis),
		ArticleGroupModel:    model.NewArticleGroupModel(mysqlConn, c.CacheRedis),
		ArticleGroupRelModel: model.GlobalArticleGroupRel,
	}
}

package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var GlobalArticleGroup ArticleGroupModel = (*customArticleGroupModel)(nil)

type (
	// ArticleGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleGroupModel.
	ArticleGroupModel interface {
		articleGroupModel
	}

	customArticleGroupModel struct {
		*defaultArticleGroupModel
	}
)

// NewArticleGroupModel returns a model for the database table.
func NewArticleGroupModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleGroupModel {
	return &customArticleGroupModel{
		defaultArticleGroupModel: newArticleGroupModel(conn, c),
	}
}

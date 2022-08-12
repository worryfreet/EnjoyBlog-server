package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleGroupRelModel = (*customArticleGroupRelModel)(nil)

type (
	// ArticleGroupRelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleGroupRelModel.
	ArticleGroupRelModel interface {
		articleGroupRelModel
	}

	customArticleGroupRelModel struct {
		*defaultArticleGroupRelModel
	}
)

// NewArticleGroupRelModel returns a model for the database table.
func NewArticleGroupRelModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleGroupRelModel {
	return &customArticleGroupRelModel{
		defaultArticleGroupRelModel: newArticleGroupRelModel(conn, c),
	}
}

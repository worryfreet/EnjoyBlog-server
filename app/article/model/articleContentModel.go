package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleContentModel = (*customArticleContentModel)(nil)

type (
	// ArticleContentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleContentModel.
	ArticleContentModel interface {
		articleContentModel
	}

	customArticleContentModel struct {
		*defaultArticleContentModel
	}
)

// NewArticleContentModel returns a model for the database table.
func NewArticleContentModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleContentModel {
	return &customArticleContentModel{
		defaultArticleContentModel: newArticleContentModel(conn, c),
	}
}

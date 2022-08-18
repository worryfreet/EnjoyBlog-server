package model

import (
	"context"
	"fmt"
)

type articleContentOtherModel interface {
	FindListBySearch(ctx context.Context, searchKey string) ([]string, error)
}

func (c *customArticleContentModel) FindListBySearch(ctx context.Context, searchKey string) ([]string, error) {
	var articleIds []string
	query := fmt.Sprintf("select %s from %s where match(article_ctt_html) against(? in boolean mode)", "article_id", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &articleIds, query, "+"+searchKey)
	if err != nil {
		return nil, err
	}
	return articleIds, nil
}

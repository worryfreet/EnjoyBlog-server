package model

import (
	"context"
	"fmt"
)

type articleGroupRelOtherModel interface {
	FindArticleIdsByGroupId(ctx context.Context, groupId string) ([]string, error)
}

func (c *customArticleGroupRelModel) FindArticleIdsByGroupId(ctx context.Context, groupId string) ([]string, error) {
	resp := make([]string, 0)
	query := fmt.Sprintf(`select %s from %s where article_group_id = ?`, "article_id", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, groupId)
	return resp, err
}

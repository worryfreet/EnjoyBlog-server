package model

import (
	"EnjoyBlog/common/request"
	"context"
	"fmt"
	"strings"
)

type articleOtherModel interface {
	FindList(ctx context.Context, req *request.ArticleListReq) ([]*Article, error)
	FindMyList(ctx context.Context, req *request.MyArticleListReq) ([]*Article, error)
	SearchList(ctx context.Context, req *request.SearchArticleListReq) ([]*Article, error)
	DeleteWithUserId(ctx context.Context, userId, articleId string) error
}

func (c *customArticleModel) FindList(ctx context.Context, req *request.ArticleListReq) ([]*Article, error) {
	var (
		resp     = make([]*Article, 0, 10)
		orderKey = "id"
		order    = "asc"
		limit    = req.PageSize
		offset   = (req.Page - 1) * req.PageSize
	)
	if req.OrderKey != "" {
		orderKey = req.OrderKey
	}
	if req.Desc == 1 {
		order = "desc"
	}
	label := "%" + req.Tag + "%"
	query := fmt.Sprintf("select %s from %s where label like ? order by %s %s limit ? offset ?", articleRows, c.table, orderKey, order)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, label, limit, offset)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *customArticleModel) FindMyList(ctx context.Context, req *request.MyArticleListReq) ([]*Article, error) {
	var (
		resp     = make([]*Article, 0, 10)
		orderKey = "id"
		order    = "asc"
		limit    = req.PageSize
		offset   = (req.Page - 1) * req.PageSize
	)
	if req.OrderKey != "" {
		orderKey = req.OrderKey
	}
	if req.Desc == 1 {
		order = "desc"
	}
	// 分类查询我的文章列表
	if req.ArticleGroupId != "" {
		articleIds, err := GlobalArticleGroupRel.FindArticleIdsByGroupId(ctx, req.ArticleGroupId)
		idsQuery := strings.Join(articleIds, ",")
		query := fmt.Sprintf(`select %s from %s where article_id in (?) order by %s %s limit ? offset ?`, articleRows, c.table, orderKey, order)
		err = c.QueryRowsNoCacheCtx(ctx, &resp, query, idsQuery, limit, offset)
		return resp, err
	}
	// 根据标签查询我的文章列表
	if req.Tag != "" {
		label := "%" + req.Tag + "%"
		query := fmt.Sprintf(`select %s from %s where label like ? and user_id = ? order by %s %s limit ? offset ?`, articleRows, c.table, orderKey, order)
		err := c.QueryRowsNoCacheCtx(ctx, &resp, query, label, req.UserId, limit, offset)
		return resp, err
	}
	// 无条件查询我的文章列表
	query := fmt.Sprintf(`select %s from %s and user_id = ? order by %s %s limit ? offset ?`, articleRows, c.table, orderKey, order)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, req.UserId, limit, offset)
	return resp, err
}

func (c *customArticleModel) SearchList(ctx context.Context, req *request.SearchArticleListReq) ([]*Article, error) {
	return nil, nil
}

func (c *customArticleModel) DeleteWithUserId(ctx context.Context, userId, articleId string) error {
	return nil
}

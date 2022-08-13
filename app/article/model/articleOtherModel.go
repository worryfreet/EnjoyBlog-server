package model

import (
	"EnjoyBlog/common/request"
	"context"
	"fmt"
	"strings"
)

type articleOtherModel interface {
	FindList(ctx context.Context, req *request.ArticleListReq, Pub int) ([]*Article, error)
	FindListByLabel(ctx context.Context, req *request.ArticleListByLabelReq, Pub int) ([]*Article, error)
	FindListByGroupId(ctx context.Context, req *request.ArticleListByGroupIdReq, Pub int) ([]*Article, error)
	FindListBySearch(ctx context.Context, req *request.SearchArticleListReq, Pub int) ([]*Article, error)
	DeleteWithUserId(ctx context.Context, userId, articleId string) error
}

// FindList 可无条件 || 根据标签获取公开的文章
func (c *customArticleModel) FindList(ctx context.Context, req *request.ArticleListReq, pub int) ([]*Article, error) {
	var (
		resp     = make([]*Article, 0, 10)
		orderKey = "create_time"
		order    = "desc"
		limit    = req.PageSize
		offset   = (req.Page - 1) * req.PageSize
	)
	if req.OrderKey != "" && strings.Index(articleRows, req.OrderKey) != -1 {
		orderKey = req.OrderKey
	}
	if req.Desc == 0 {
		order = "asc"
	}
	// 需要和前端商量好, 关键字段数字代表的含义
	var query string
	if req.UserId != "-1" {
		query = fmt.Sprintf("select %s from %s where is_pub = ? and user_id = ? order by %s %s limit ? offset ?", articleRows, c.table, orderKey, order)
		err := c.QueryRowsNoCacheCtx(ctx, &resp, query, pub, req.UserId, limit, offset)
		return resp, err
	}
	query = fmt.Sprintf("select %s from %s where is_pub = ? order by %s %s limit ? offset ?", articleRows, c.table, orderKey, order)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, pub, limit, offset)
	return resp, err
}

func (c *customArticleModel) FindListByLabel(ctx context.Context, req *request.ArticleListByLabelReq, Pub int) ([]*Article, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customArticleModel) FindListByGroupId(ctx context.Context, req *request.ArticleListByGroupIdReq, Pub int) ([]*Article, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customArticleModel) FindListBySearch(ctx context.Context, req *request.SearchArticleListReq, Pub int) ([]*Article, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customArticleModel) DeleteWithUserId(ctx context.Context, userId, articleId string) error {
	return nil
}

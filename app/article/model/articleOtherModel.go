package model

import (
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/request"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

type articleOtherModel interface {
	FindList(ctx context.Context, req *request.ArticleListReq, tokenUserId string) ([]*Article, error)
	FindListByLabel(ctx context.Context, req *request.ArticleListByLabelReq, tokenUserId string) ([]*Article, error)
	FindListByGroupId(ctx context.Context, req *request.ArticleListByGroupIdReq, tokenUserId string) ([]*Article, error)
	FindListBySearch(ctx context.Context, req *request.SearchArticleListReq, tokenUserId string) ([]*Article, error)
	DeleteWithUserId(ctx context.Context, userId, articleId string) error
	TransAdd(ctx context.Context, article *Article, rel *ArticleGroupRel, ctt *ArticleContent) error
	TransUpdate(ctx context.Context, article *Article, ctt *ArticleContent) error
}

// FindList 可无条件 || 根据标签获取公开的文章
func (c *customArticleModel) FindList(ctx context.Context, req *request.ArticleListReq, tokenUserId string) ([]*Article, error) {
	var (
		resp     = make([]*Article, 0, 10)
		orderKey = "create_time"
		order    = "desc"
		limit    = req.PageSize
		offset   = (req.Page - 1) * req.PageSize
		query    string
	)
	// 检查sql参数
	if req.OrderKey != "" && strings.Index(articleRows, req.OrderKey) != -1 {
		orderKey = req.OrderKey
	}
	if req.Desc == 0 {
		order = "asc"
	}
	// 获取自己的文章列表
	if req.UserId == tokenUserId {
		query = fmt.Sprintf("select %s from %s where user_id = ? order by %s %s limit ? offset ?", articleRows, c.table, orderKey, order)
		err := c.QueryRowsNoCacheCtx(ctx, &resp, query, req.UserId, limit, offset)
		return resp, err
	}
	// 获取某一个人的开放文章
	if req.UserId != "-1" {
		query = fmt.Sprintf("select %s from %s where is_pub = 1 and user_id = ? order by %s %s limit ? offset ?", articleRows, c.table, orderKey, order)
		err := c.QueryRowsNoCacheCtx(ctx, &resp, query, req.UserId, limit, offset)
		return resp, err
	}
	// userId = -1 获取所有人的公共文章
	query = fmt.Sprintf("select %s from %s where is_pub = 1 order by %s %s limit ? offset ?", articleRows, c.table, orderKey, order)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, limit, offset)
	return resp, err
}

func (c *customArticleModel) FindListByLabel(ctx context.Context, req *request.ArticleListByLabelReq, tokenUserId string) ([]*Article, error) {
	var (
		resp     = make([]*Article, 0, 10)
		orderKey = "create_time"
		order    = "desc"
		limit    = req.PageSize
		offset   = (req.Page - 1) * req.PageSize
		query    string
		labelReq = "%" + req.Label + "%"
	)
	// 检查sql参数
	if req.OrderKey != "" && strings.Index(articleRows, req.OrderKey) != -1 {
		orderKey = req.OrderKey
	}
	if req.Desc == 0 {
		order = "asc"
	}
	// 根据label获取自己的文章列表
	if req.UserId == tokenUserId {
		query = fmt.Sprintf("select %s from %s where user_id = ? and label like ? order by %s %s limit ? offset ?", articleRows, c.table, orderKey, order)
		err := c.QueryRowsNoCacheCtx(ctx, &resp, query, req.UserId, labelReq, limit, offset)
		return resp, err
	}
	// 根据label获取某一个人的开放文章
	if req.UserId != "-1" {
		query = fmt.Sprintf("select %s from %s where is_pub = 1 and user_id = ? and label like ? order by %s %s limit ? offset ?", articleRows, c.table, orderKey, order)
		err := c.QueryRowsNoCacheCtx(ctx, &resp, query, req.UserId, labelReq, limit, offset)
		return resp, err
	}
	// userId = -1 根据label获取所有人的公共文章
	query = fmt.Sprintf("select %s from %s where is_pub = 1 and label like ? order by %s %s limit ? offset ?", articleRows, c.table, orderKey, order)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, labelReq, limit, offset)
	return resp, err
}

func (c *customArticleModel) FindListByGroupId(ctx context.Context, req *request.ArticleListByGroupIdReq, tokenUserId string) ([]*Article, error) {
	var (
		resp          = make([]*Article, 0, 10)
		orderKey      = "create_time"
		order         = "desc"
		limit         = req.PageSize
		offset        = (req.Page - 1) * req.PageSize
		query         string
		articleIdsReq string
	)
	// 检查sql参数
	if req.OrderKey != "" && strings.Index(articleRows, req.OrderKey) != -1 {
		orderKey = req.OrderKey
	}
	if req.Desc == 0 {
		order = "asc"
	}
	articleIds, err := GlobalArticleGroupRel.FindArticleIdsByGroupId(ctx, req.ArticleGroupId)
	if err != nil {
		if err == ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	articleIdsReq = strings.Join(articleIds, "','")
	articleIdsReq = "'" + articleIdsReq + "'"
	// 根据articleGroupId获取自己的文章列表
	if req.UserId == tokenUserId {
		query = fmt.Sprintf("select %s from %s where user_id = ? and article_id in (%s) order by %s %s limit ? offset ?", articleRows, c.table, articleIdsReq, orderKey, order)
		err = c.QueryRowsNoCacheCtx(ctx, &resp, query, req.UserId, limit, offset)
		return resp, err
	}
	// 根据articleGroupId获取某一个人的开放文章
	query = fmt.Sprintf("select %s from %s where is_pub = 1 and user_id = ? and article_id in (%s) order by %s %s limit ? offset ?", articleRows, c.table, articleIdsReq, orderKey, order)
	err = c.QueryRowsNoCacheCtx(ctx, &resp, query, req.UserId, limit, offset)
	return resp, err
}

func (c *customArticleModel) FindListBySearch(ctx context.Context, req *request.SearchArticleListReq, tokenUserId string) ([]*Article, error) {
	var (
		resp          = make([]*Article, 0, 10)
		orderKey      = "create_time"
		order         = "desc"
		limit         = req.PageSize
		offset        = (req.Page - 1) * req.PageSize
		query         string
		articleIdsReq string
	)
	// 检查sql参数
	if req.OrderKey != "" && strings.Index(articleRows, req.OrderKey) != -1 {
		orderKey = req.OrderKey
	}
	if req.Desc == 0 {
		order = "asc"
	}

	articleIds, err := GlobalArticleContent.FindListBySearch(ctx, req.SearchKey)
	articleIdsReq = strings.Join(articleIds, "','")
	articleIdsReq = "'" + articleIdsReq + "'"
	if err != nil {
		if err == ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	// 根据searchKey获取自己的文章列表
	if req.UserId == tokenUserId {
		query = fmt.Sprintf("select %s from %s where user_id = ? and article_id in (%s) order by %s %s limit ? offset ?", articleRows, c.table, articleIdsReq, orderKey, order)
		err = c.QueryRowsNoCacheCtx(ctx, &resp, query, req.UserId, limit, offset)
		return resp, err
	}
	// 根据searchKey获取所有符合条件的文章列表
	if req.UserId == "-1" {
		query = fmt.Sprintf("select %s from %s where article_id in (%s) order by %s %s limit ? offset ?", articleRows, c.table, articleIdsReq, orderKey, order)
		err = c.QueryRowsNoCacheCtx(ctx, &resp, query, limit, offset)
		return resp, err
	}
	// 根据articleGroupId获取某一个人的开放文章
	query = fmt.Sprintf("select %s from %s where is_pub = 1 and user_id = ? and article_id in (%s) order by %s %s limit ? offset ?", articleRows, c.table, articleIdsReq, orderKey, order)
	err = c.QueryRowsNoCacheCtx(ctx, &resp, query, req.UserId, limit, offset)
	return resp, err
}

func (c *customArticleModel) DeleteWithUserId(ctx context.Context, userId, articleId string) error {
	article, err := c.FindOneByArticleId(ctx, articleId)
	if err != nil {
		return err
	}
	if article.UserId != userId {
		return errorx.StatusErrUserNoAuth
	}
	relInfo, err := GlobalArticleGroupRel.FindOneByArticleId(ctx, articleId)
	if err != nil {
		return err
	}
	cttInfo, err := GlobalArticleContent.FindOneByArticleId(ctx, articleId)
	if err != nil {
		return err
	}
	err = c.TransactCtx(ctx, func(innerCtx context.Context, session sqlx.Session) error {
		if err = c.Delete(ctx, article.Id, session); err != nil {
			return err
		}
		if err = GlobalArticleGroupRel.Delete(ctx, relInfo.Id, session); err != nil {
			return err
		}
		if err = GlobalArticleContent.Delete(ctx, cttInfo.Id, session); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (c *customArticleModel) TransAdd(ctx context.Context, article *Article, rel *ArticleGroupRel, ctt *ArticleContent) error {
	return c.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		// 1. 插入article表
		if _, err := c.Insert(ctx, article, session); err != nil {
			return err
		}
		// 2. 插入article_group_rel表
		if _, err := GlobalArticleGroupRel.Insert(ctx, rel, session); err != nil {
			return err
		}
		// 3. 插入article_content表
		if _, err := GlobalArticleContent.Insert(ctx, ctt, session); err != nil {
			return err
		}
		return nil
	})
}

func (c *customArticleModel) TransUpdate(ctx context.Context, article *Article, ctt *ArticleContent) error {
	return c.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		// 1. 更新article表
		if err := c.Update(ctx, article, session); err != nil {
			return err
		}
		// 2. 更新article_content表
		if err := GlobalArticleContent.Update(ctx, ctt, session); err != nil {
			return err
		}
		return nil
	})
}

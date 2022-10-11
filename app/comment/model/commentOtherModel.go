package model

import (
	"EnjoyBlog/common/request"
	"context"
	"errors"
	"strconv"
	"strings"
)

const (
	TOTAL = "cmt:total"
	IS    = "cmt:is"
	TRUE  = "true"
	ZERO  = "0"
)

type commentOtherModel interface {
	AddComment(ctx context.Context, req *Comment, cmtInfo *request.CmtSupport) error
	FindListByArticleId(ctx context.Context, articleId string) ([]*Comment, error)
	SyncMultiCmtSupport(ctx context.Context) (int, error)
	RedisSupport(ctx context.Context, req *request.CmtSupport) error
}

func (c *customCommentModel) AddComment(ctx context.Context, req *Comment, cmtInfo *request.CmtSupport) error {
	// 1. 插入Mysql
	if _, err := c.Insert(ctx, req, nil); err != nil {
		return err
	}
	// 2. 数量存入Redis
	if err := c.Redis.SetCtx(ctx, Key(TOTAL, cmtInfo), ZERO); err != nil {
		return err
	}
	return nil
}

func (c *customCommentModel) FindListByArticleId(ctx context.Context, articleId string) ([]*Comment, error) {
	return nil, nil
}

func (c *customCommentModel) SyncMultiCmtSupport(ctx context.Context) (int, error) {
	// TODO
	// 从mysql 评论表里中查找所有信息 存到数组里
	// 依次从redis中寻找, 然后更改数组里的值
	// (如果不一样才修改) Update Mysql
	var keys []string
	for i, key := range keys {
		// 1. 获取点赞数量
		ks := strings.Split(key, ":")
		req := &request.CmtSupport{
			CmtId: ks[4],
			IP:    ks[3],
		}
		total, err := c.Redis.GetCtx(ctx, Key(TOTAL, req))
		if err != nil {
			return i, err
		}
		// 2. 从mysql查找评论信息
		cmtInfo, err := c.FindOneByCommentId(ctx, req.CmtId)
		if err != nil {
			return i, err
		}
		// 3. 更新mysql
		newTotal, err := strconv.Atoi(total)
		if err != nil {
			return i, err
		}
		cmtInfo.SupportTotal = int64(newTotal)
		err = c.Update(ctx, cmtInfo, nil)
		if err != nil {
			return i, err
		}
	}
	return len(keys), nil
}

func (c *customCommentModel) RedisSupport(ctx context.Context, req *request.CmtSupport) error {
	// 1. 检查是否重复点赞
	err := c.CheckRedisSupport(ctx, req)
	if err != nil {
		return err
	}
	// 2. 设置已点赞 true
	err = c.Redis.SetCtx(ctx, Key("is", req), TRUE)
	if err != nil {
		return err
	}
	// 3. 总量增加
	_, err = c.Redis.IncrCtx(ctx, Key("total", req))
	if err != nil {
		return err
	}
	return nil
}

func (c *customCommentModel) CheckRedisSupport(ctx context.Context, req *request.CmtSupport) error {
	isSpt, err := c.Redis.GetCtx(ctx, Key("is", req))
	if err != nil {
		return err
	}
	if isSpt == TRUE {
		return errors.New("已经点过赞啦")
	}
	return nil
}

// cmt:total:support:IP:articleId
// cmt:is:support:IP:articleId
// art:total:support:IP:commentId
// art:is:support:IP:commentId
func Key(head string, req *request.CmtSupport) string {
	ks := []string{head, "support", req.IP, req.CmtId}
	return strings.Join(ks, ":")
}

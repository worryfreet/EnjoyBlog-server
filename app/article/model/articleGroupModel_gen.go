// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	articleGroupFieldNames          = builder.RawFieldNames(&ArticleGroup{})
	articleGroupRows                = strings.Join(articleGroupFieldNames, ",")
	articleGroupRowsExpectAutoSet   = strings.Join(stringx.Remove(articleGroupFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	articleGroupRowsWithPlaceHolder = strings.Join(stringx.Remove(articleGroupFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheArticleGroupIdPrefix             = "cache:articleGroup:id:"
	cacheArticleGroupArticleGroupIdPrefix = "cache:articleGroup:articleGroupId:"
	cacheArticleGroupUserIdPrefix         = "cache:articleGroup:userId:"
)

type (
	articleGroupModel interface {
		Insert(ctx context.Context, data *ArticleGroup) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ArticleGroup, error)
		FindOneByArticleGroupId(ctx context.Context, articleGroupId string) (*ArticleGroup, error)
		FindOneByUserId(ctx context.Context, userId string) (*ArticleGroup, error)
		Update(ctx context.Context, data *ArticleGroup) error
		Delete(ctx context.Context, id int64) error
	}

	defaultArticleGroupModel struct {
		sqlc.CachedConn
		table string
	}

	ArticleGroup struct {
		Id                int64        `db:"id"`                  // 自增id
		UserId            string       `db:"user_id"`             // 用户id
		ArticleGroupId    string       `db:"article_group_id"`    // 分类目录id
		ArticleGroupTitle string       `db:"article_group_title"` // 分类目录标题
		ParentId          string       `db:"parent_id"`           // 上级分类目录id
		CreateTime        time.Time    `db:"create_time"`
		UpdateTime        time.Time    `db:"update_time"`
		DeletedTime       sql.NullTime `db:"deleted_time"`
	}
)

func newArticleGroupModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultArticleGroupModel {
	return &defaultArticleGroupModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`article_group`",
	}
}

func (m *defaultArticleGroupModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	articleGroupArticleGroupIdKey := fmt.Sprintf("%s%v", cacheArticleGroupArticleGroupIdPrefix, data.ArticleGroupId)
	articleGroupIdKey := fmt.Sprintf("%s%v", cacheArticleGroupIdPrefix, id)
	articleGroupUserIdKey := fmt.Sprintf("%s%v", cacheArticleGroupUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, articleGroupArticleGroupIdKey, articleGroupIdKey, articleGroupUserIdKey)
	return err
}

func (m *defaultArticleGroupModel) FindOne(ctx context.Context, id int64) (*ArticleGroup, error) {
	articleGroupIdKey := fmt.Sprintf("%s%v", cacheArticleGroupIdPrefix, id)
	var resp ArticleGroup
	err := m.QueryRowCtx(ctx, &resp, articleGroupIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", articleGroupRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultArticleGroupModel) FindOneByArticleGroupId(ctx context.Context, articleGroupId string) (*ArticleGroup, error) {
	articleGroupArticleGroupIdKey := fmt.Sprintf("%s%v", cacheArticleGroupArticleGroupIdPrefix, articleGroupId)
	var resp ArticleGroup
	err := m.QueryRowIndexCtx(ctx, &resp, articleGroupArticleGroupIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `article_group_id` = ? limit 1", articleGroupRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, articleGroupId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultArticleGroupModel) FindOneByUserId(ctx context.Context, userId string) (*ArticleGroup, error) {
	articleGroupUserIdKey := fmt.Sprintf("%s%v", cacheArticleGroupUserIdPrefix, userId)
	var resp ArticleGroup
	err := m.QueryRowIndexCtx(ctx, &resp, articleGroupUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", articleGroupRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultArticleGroupModel) Insert(ctx context.Context, data *ArticleGroup) (sql.Result, error) {
	articleGroupArticleGroupIdKey := fmt.Sprintf("%s%v", cacheArticleGroupArticleGroupIdPrefix, data.ArticleGroupId)
	articleGroupIdKey := fmt.Sprintf("%s%v", cacheArticleGroupIdPrefix, data.Id)
	articleGroupUserIdKey := fmt.Sprintf("%s%v", cacheArticleGroupUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, articleGroupRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ArticleGroupId, data.ArticleGroupTitle, data.ParentId, data.DeletedTime)
	}, articleGroupArticleGroupIdKey, articleGroupIdKey, articleGroupUserIdKey)
	return ret, err
}

func (m *defaultArticleGroupModel) Update(ctx context.Context, newData *ArticleGroup) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	articleGroupArticleGroupIdKey := fmt.Sprintf("%s%v", cacheArticleGroupArticleGroupIdPrefix, data.ArticleGroupId)
	articleGroupIdKey := fmt.Sprintf("%s%v", cacheArticleGroupIdPrefix, data.Id)
	articleGroupUserIdKey := fmt.Sprintf("%s%v", cacheArticleGroupUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, articleGroupRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.ArticleGroupId, newData.ArticleGroupTitle, newData.ParentId, newData.DeletedTime, newData.Id)
	}, articleGroupArticleGroupIdKey, articleGroupIdKey, articleGroupUserIdKey)
	return err
}

func (m *defaultArticleGroupModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheArticleGroupIdPrefix, primary)
}

func (m *defaultArticleGroupModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", articleGroupRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultArticleGroupModel) tableName() string {
	return m.table
}

package model

import (
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/request"
	"context"
	"fmt"
)

type userOtherModel interface {
	FindList(ctx context.Context, req *request.PageLimitReq) ([]*User, error)
}

func (m *defaultUserModel) FindList(ctx context.Context, req *request.PageLimitReq) ([]*User, error) {
	if req.Page < 1 || req.PageSize < 0 {
		return nil, errorx.StatusErrParam
	}
	var (
		resp     []*User
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
	query := fmt.Sprintf("select %s from %s order by %s %s limit ? offset ?", userRows, m.table, orderKey, order)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

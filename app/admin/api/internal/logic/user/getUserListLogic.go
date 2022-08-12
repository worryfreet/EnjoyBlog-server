package user

import (
	"EnjoyBlog/app/user/rpc/types/user"
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/utils"
	"context"

	"EnjoyBlog/app/admin/api/internal/svc"
	"EnjoyBlog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	pageLimit := new(user.UserListReq)
	err = utils.FillModel(&pageLimit, req)
	if err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	users, err := l.svcCtx.UserRPC.GetUserList(l.ctx, pageLimit)
	if err != nil {
		return nil, err
	}
	resp = new(types.UserListResp)
	err = utils.FillModel(&resp.UserList, users.UserList)
	if err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, err
	}
	return
}

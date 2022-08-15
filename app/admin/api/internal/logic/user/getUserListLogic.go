package user

import (
	"EnjoyBlog/app/user/rpc/userclient"
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/utils"
	"context"

	"EnjoyBlog/app/admin/api/internal/svc"
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

func (l *GetUserListLogic) GetUserList(req *userclient.UserListReq) (resp *userclient.UserListResp, err error) {
	pageLimit := new(userclient.UserListReq)
	err = utils.FillModel(&pageLimit, req)
	if err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	users, err := l.svcCtx.UserRPC.GetUserList(l.ctx, pageLimit)
	if err != nil {
		return nil, err
	}
	resp = new(userclient.UserListResp)
	err = utils.FillModel(&resp.UserList, users.UserList)
	if err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, err
	}
	return
}

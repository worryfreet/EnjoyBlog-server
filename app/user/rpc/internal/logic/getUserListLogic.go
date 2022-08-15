package logic

import (
	"EnjoyBlog/app/user/rpc/userclient"
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/request"
	"EnjoyBlog/common/utils"
	"context"

	"EnjoyBlog/app/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *userclient.UserListReq) (*userclient.UserListResp, error) {
	pageLimit := new(request.PageLimitReq)
	if err := utils.FillModel(&pageLimit, in); err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	users, err := l.svcCtx.UserModel.FindList(l.ctx, pageLimit)
	if err != nil {
		l.Logger.Error("user-rpc查询用户列表信息错误 err: ", err)
		return nil, errorx.StatusErrSystemBusy
	}
	resp := new(userclient.UserListResp)
	if err := utils.FillModel(&resp.UserList, users); err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	return resp, nil
}

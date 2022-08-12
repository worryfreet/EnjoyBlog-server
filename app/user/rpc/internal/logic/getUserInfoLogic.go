package logic

import (
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/utils"
	"context"

	"EnjoyBlog/app/user/rpc/internal/svc"
	"EnjoyBlog/app/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.UserInfoReq) (*user.UserInfoResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		l.Logger.Error("rpc 根据userId查询用户信息失败, err: ", err)
		return nil, errorx.StatusErrSystemBusy
	}
	resp := new(user.UserInfoResp)
	if err = utils.FillModel(&resp, userInfo); err != nil {
		l.Logger.Error("rpc FillModel(&resp, userInfo)填充数据失败, err: ", err)
		return nil, errorx.StatusErrParam
	}
	return resp, nil
}

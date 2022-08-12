package user

import (
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/utils"
	"context"

	"EnjoyBlog/app/user/api/internal/svc"
	"EnjoyBlog/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateInfoReq) (resp *types.User, err error) {
	email := l.ctx.Value("email").(string)
	userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err != nil {
		l.Logger.Error("根据邮箱", email, "查询用户信息错误 err: ", err)
		return nil, errorx.StatusErrSystemBusy
	}
	if err = utils.FillModel(&userInfo, req); err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	// 给resp赋值
	resp = new(types.User)
	if err = utils.FillModel(&resp, userInfo); err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	err = l.svcCtx.UserModel.Update(l.ctx, userInfo)
	if err != nil {
		l.Logger.Error("更新用户信息错误 err: ", err)
		return nil, errorx.StatusErrSystemBusy
	}
	return
}

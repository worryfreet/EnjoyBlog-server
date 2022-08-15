package user

import (
	"EnjoyBlog/app/user/rpc/userclient"
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/utils"
	"context"

	"EnjoyBlog/app/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *userclient.UserInfoReq) (resp *userclient.UserInfoResp, err error) {
	l.Logger.Info(l.ctx.Value("userId"))
	userInfo, err := l.svcCtx.UserRPC.GetUserInfo(l.ctx, &userclient.UserInfoReq{UserId: req.UserId})
	if err != nil {
		l.Logger.Error("admin.api 调用 user.rpc 根据userId查询用户信息失败 err: ", err)
		return nil, errorx.StatusErrSystemBusy
	}
	resp = new(userclient.UserInfoResp)
	if err = utils.FillModel(&resp, userInfo); err != nil {
		l.Logger.Error("admin.api FillModel(&resp, userInfo) 填充数据失败 err: ", err)
		return nil, errorx.StatusErrParam
	}
	return
}

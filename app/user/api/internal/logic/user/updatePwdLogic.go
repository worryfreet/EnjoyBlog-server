package user

import (
	"EnjoyBlog/app/user/api/internal/svc"
	"EnjoyBlog/app/user/api/internal/types"
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/utils/encrypt"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePwdLogic {
	return &UpdatePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePwdLogic) UpdatePwd(req *types.UpdatePwdReq) (resp *types.User, err error) {
	email := l.ctx.Value("email").(string)
	userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err != nil {
		return nil, errorx.StatusErrSystemBusy
	}
	// 没有忘记密码, 且老密码填写错误
	req.Passwd = encrypt.MD5V([]byte(encrypt.RsaPriDecode(req.Passwd)))
	if !req.Forget && req.Passwd != userInfo.Password {
		l.Logger.Error("密码错误, req.Passwd: ", req.Passwd)
		return nil, errorx.StatusErrUserPwd
	}

	// TODO: 忘记密码, 发邮箱或者手机验证码确认
	if req.Forget {
	}

	// 更新密码核心操作
	userInfo.Password = encrypt.MD5V([]byte(encrypt.RsaPriDecode(req.NewPasswd)))
	err = l.svcCtx.UserModel.Update(l.ctx, userInfo)
	if err != nil {
		return nil, errorx.StatusErrSystemBusy
	}
	return
}

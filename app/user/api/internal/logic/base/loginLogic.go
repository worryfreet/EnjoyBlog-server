package base

import (
	"EnjoyBlog/app/user/model"
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/utils"
	"EnjoyBlog/common/utils/encrypt"
	"context"
	"github.com/form3tech-oss/jwt-go"
	"strings"
	"time"

	"EnjoyBlog/app/user/api/internal/svc"
	"EnjoyBlog/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if len(strings.TrimSpace(req.Email)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.StatusErrParam
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.StatusErrUserNotFound
	default:
		return nil, errorx.StatusErrNotKnown
	}
	// RSA解密, MD5加密, 再进行比较密码
	req.Password = encrypt.MD5V([]byte(encrypt.RsaPriDecode(req.Password)))
	if req.Password != userInfo.Password {
		return nil, errorx.StatusErrUserPwd
	}
	return l.tokenNext(resp, userInfo)
}

func (l *LoginLogic) tokenNext(resp *types.LoginResp, userInfo *model.User) (*types.LoginResp, error) {
	jwtToken, err := l.GetJwtToken(userInfo)
	if err != nil {
		l.Logger.Error("loginLogic.Login--jwtAuth.GetJwtToken(r, claims) err: ", err)
		return nil, errorx.StatusErrUserNoAuth
	}
	resp = &types.LoginResp{
		AccessToken:  jwtToken,
		AccessExpire: l.svcCtx.Config.JwtAuth.AccessExpire,
	}
	if err = utils.FillModel(&resp.User, userInfo); err != nil {
		l.Logger.Error("loginLogic.Login--utils.FillModel(&resp.User, userInfo) err: ", err)
		return nil, errorx.StatusErrParam
	}
	return resp, nil
}

func (l *LoginLogic) GetJwtToken(userInfo *model.User) (string, error) {
	claims := make(jwt.MapClaims)
	claims["id"] = userInfo.Id
	claims["userId"] = userInfo.UserId
	claims["email"] = userInfo.Email
	claims["userInfo"] = userInfo
	claims["expire"] = time.Now().Unix() + l.svcCtx.Config.JwtAuth.AccessExpire
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
}

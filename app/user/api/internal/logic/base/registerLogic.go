package base

import (
	"EnjoyBlog/app/user/model"
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/utils"
	"EnjoyBlog/common/utils/encrypt"
	"context"
	"database/sql"
	"time"

	"EnjoyBlog/app/user/api/internal/svc"
	"EnjoyBlog/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) error {
	errInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err == nil {
		l.Logger.Error("该用户已存在, 若数据库为空, 则检查Redis缓存, 该用户id为: ", errInfo.Id)
		return errorx.StatusErrUserExist
	}
	userData := &model.User{
		UserId:      utils.NewUUID(),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		DeletedTime: sql.NullTime{Valid: false},
	}
	if err = utils.FillModel(&userData, req); err != nil {
		l.Logger.Error("registerLogic.Register -- utils.FillModel(&userData, req) err: ", err)
		return errorx.StatusErrParam
	}
	// RSA解密, MD5加密, 再存入数据库
	userData.Password = encrypt.MD5V([]byte(encrypt.RsaPriDecode(req.Password)))
	_, err = l.svcCtx.UserModel.Insert(l.ctx, userData)
	if err != nil {
		l.Logger.Error("registerLogic.Register -- l.svcCtx.UserModel.Insert(l.ctx, userData) err: ", err)
		return errorx.StatusErrNotKnown
	}
	return nil
}

func defaultFillNullVal(user *model.User) {

}

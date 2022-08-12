package user

import (
	"EnjoyBlog/app/user/api/internal/svc"
	"EnjoyBlog/app/user/api/internal/types"
	"EnjoyBlog/app/user/model"
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/utils"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
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

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.User, err error) {
	id, _ := strconv.Atoi(l.ctx.Value("id").(json.Number).String())
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, int64(id))
	if err == model.ErrNotFound {
		l.Logger.Error("用户", id, "不存在, err: ", err)
		return nil, errorx.StatusErrUserNotFound
	}
	if err != nil {
		l.Logger.Error("查询用户", id, "信息失败, err: ", err)
		return nil, errorx.StatusErrSystemBusy
	}
	// 给resp赋值
	err = utils.FillModel(&resp, userInfo)
	if err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	return
}

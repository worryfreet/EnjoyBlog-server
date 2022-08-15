package logic

import (
	"context"

	"EnjoyBlog/app/user/rpc/internal/svc"
	"EnjoyBlog/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllLogic {
	return &GetAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllLogic) GetAll(in *user.UserListReq, stream user.User_GetAllServer) error {
	// todo: add your logic here and delete this line

	return nil
}

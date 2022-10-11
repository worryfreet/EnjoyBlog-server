package logic

import (
	"context"

	"EnjoyBlog/app/article/api/internal/svc"
	"EnjoyBlog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SupportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSupportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SupportLogic {
	return &SupportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SupportLogic) Support(req *types.SupportReq) error {
	// todo: add your logic here and delete this line

	return nil
}

package article

import (
	"context"

	"EnjoyBlog/app/admin/api/internal/svc"
	"EnjoyBlog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleInfoLogic {
	return &GetArticleInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleInfoLogic) GetArticleInfo(req *types.ArticleInfoReq) (resp *types.ArticleInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}

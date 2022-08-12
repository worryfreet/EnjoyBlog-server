package logic

import (
	"context"

	"EnjoyBlog/app/article/api/internal/svc"
	"EnjoyBlog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyArticleListLogic {
	return &GetMyArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyArticleListLogic) GetMyArticleList(req *types.MyArticleListReq) (resp *types.ArticleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}

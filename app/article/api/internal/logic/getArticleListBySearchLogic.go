package logic

import (
	"context"

	"EnjoyBlog/app/article/api/internal/svc"
	"EnjoyBlog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleListBySearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleListBySearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListBySearchLogic {
	return &GetArticleListBySearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleListBySearchLogic) GetArticleListBySearch(req *types.SearchArticleListReq) (resp *types.ArticleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}

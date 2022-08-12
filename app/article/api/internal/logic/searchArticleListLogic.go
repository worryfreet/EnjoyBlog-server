package logic

import (
	"context"

	"EnjoyBlog/app/article/api/internal/svc"
	"EnjoyBlog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchArticleListLogic {
	return &SearchArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchArticleListLogic) SearchArticleList(req *types.SearchArticleListReq) (resp *types.ArticleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}

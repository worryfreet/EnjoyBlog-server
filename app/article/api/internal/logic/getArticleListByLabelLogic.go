package logic

import (
	"context"

	"EnjoyBlog/app/article/api/internal/svc"
	"EnjoyBlog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleListByLabelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleListByLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListByLabelLogic {
	return &GetArticleListByLabelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleListByLabelLogic) GetArticleListByLabel(req *types.ArticleListByLabelReq) (resp *types.ArticleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}

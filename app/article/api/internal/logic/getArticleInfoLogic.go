package logic

import (
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/utils"
	"context"

	"EnjoyBlog/app/article/api/internal/svc"
	"EnjoyBlog/app/article/api/internal/types"

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

func (l *GetArticleInfoLogic) GetArticleInfo(req *types.ArticleInfoReq) (resp *types.ArticleInfoWithContent, err error) {
	articleInfo, err := l.svcCtx.ArticleModel.FindOneByArticleId(l.ctx, req.ArticleId)
	if err != nil {
		return nil, errorx.StatusErrSystemBusy
	}
	resp = new(types.ArticleInfoWithContent)
	if err = utils.FillModel(&resp, articleInfo); err != nil {
		return nil, errorx.StatusErrParam
	}
	return
}

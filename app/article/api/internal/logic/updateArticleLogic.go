package logic

import (
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/request"
	"EnjoyBlog/common/utils"
	"context"

	"EnjoyBlog/app/article/api/internal/svc"
	"EnjoyBlog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.UpdateArticleInfoReq) error {
	// 1. 转换入参结构体
	var queryReq *request.EditArticleInfoReq
	if err := utils.FillModel(&queryReq, req); err != nil {
		l.Logger.Error("FillModel err: ", err)
		return errorx.StatusErrParam
	}

	// 2. 设置article入参
	article, err := l.svcCtx.ArticleModel.FindOneByArticleId(l.ctx, req.ArticleId)
	if err != nil {
		l.Logger.Error("查询文章信息失败, err ", err)
		return errorx.StatusErrSystemBusy
	}
	_ = utils.FillModel(&article, req)
	if len(req.ArticleCttHtml) > 100 {
		article.ArticleCttHead = req.ArticleCttHtml[:100]
	} else {
		article.ArticleCttHead = req.ArticleCttHtml
	}

	// 3. 设置article_content入参
	ctt, err := l.svcCtx.ArticleContentModel.FindOneByArticleId(l.ctx, req.ArticleId)
	if err != nil {
		l.Logger.Error("查询文章主体内容信息失败, err ", err)
		return errorx.StatusErrSystemBusy
	}
	_ = utils.FillModel(&ctt, req)

	// 4. 调用方法, 更新文章进Mysql
	if err := l.svcCtx.ArticleModel.TransUpdate(l.ctx, article, ctt); err != nil {
		l.Logger.Error("更新文章事务失败, err ", err)
		return errorx.StatusErrSystemBusy
	}
	return nil
}

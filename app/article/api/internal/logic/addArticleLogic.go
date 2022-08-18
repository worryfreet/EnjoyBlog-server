package logic

import (
	"EnjoyBlog/app/article/api/internal/global"
	"EnjoyBlog/app/article/model"
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/request"
	"EnjoyBlog/common/utils"
	"context"

	"EnjoyBlog/app/article/api/internal/svc"
	"EnjoyBlog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleLogic) AddArticle(req *types.AddArticleInfoReq) error {
	// 1. 转换入参结构体
	var queryReq *request.EditArticleInfoReq
	if err := utils.FillModel(&queryReq, req); err != nil {
		l.Logger.Error("FillModel err: ", err)
		return errorx.StatusErrParam
	}

	// 2. 设置article入参
	article := new(model.Article)
	_ = utils.FillModel(&article, req)
	article.ArticleId = utils.NewUUID()
	article.UserId = global.Jwt.Claims["userId"].(string)
	if len(req.ArticleCttHtml) > 100 {
		article.ArticleCttHead = req.ArticleCttHtml[:100]
	} else {
		article.ArticleCttHead = req.ArticleCttHtml
	}

	// 3. 设置article_group_rel入参
	rel := new(model.ArticleGroupRel)
	_ = utils.FillModel(&rel, req)
	rel.ArticleId = article.ArticleId

	// 4. 设置article_content入参
	ctt := new(model.ArticleContent)
	_ = utils.FillModel(&ctt, req)
	ctt.ArticleId = article.ArticleId

	// 5. 调用方法, 添加文章进Mysql
	if err := l.svcCtx.ArticleModel.TransAdd(l.ctx, article, rel, ctt); err != nil {
		return errorx.StatusErrSystemBusy
	}
	return nil
}

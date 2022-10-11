package logic

import (
	"EnjoyBlog/app/article/api/internal/global"
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
	// 1. 查询文章概览信息
	articleInfo, err := l.svcCtx.ArticleModel.FindOneByArticleId(l.ctx, req.ArticleId)
	if err != nil {
		l.Logger.Error("查询文章概览信息失败, err: ", err)
		return nil, errorx.StatusErrSystemBusy
	}
	// 2. 鉴权
	userId := global.Jwt.Claims["userId"].(string)
	if userId != articleInfo.UserId && articleInfo.IsPub != 1 {
		l.Logger.Error("用户访问文章权限不够")
		return nil, errorx.StatusErrUserNoAuth
	}
	// 3. 赋值
	resp = new(types.ArticleInfoWithContent)
	if err = utils.FillModel(&resp, articleInfo); err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	// 4. 查询文章主体内容
	ctt, err := l.svcCtx.ArticleContentModel.FindOneByArticleId(l.ctx, articleInfo.ArticleId)
	if err != nil {
		l.Logger.Error("查询文章主体内容失败, err: ", err)
		return nil, err
	}
	resp.ArticleCttHtml = ctt.ArticleCttHtml
	resp.ArticleCttMd = ctt.ArticleCttMd
	return
}

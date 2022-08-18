package logic

import (
	"EnjoyBlog/app/article/api/internal/global"
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/request"
	"EnjoyBlog/common/utils"
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
	// 1. 转换入参结构体
	var queryReq *request.SearchArticleListReq
	if err = utils.FillModel(&queryReq, req); err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	// 2. 查询文章列表
	tokenUserId := global.Jwt.Claims["userId"].(string)
	articles, err := l.svcCtx.ArticleModel.FindListBySearch(l.ctx, queryReq, tokenUserId)
	if err != nil {
		l.Logger.Error("查询文章列表失败, err: ", err)
		return nil, errorx.StatusErrSystemBusy
	}
	// 3. 转换出参结构体
	resp = new(types.ArticleListResp)
	if err = utils.FillModel(&resp.ArticleList, articles); err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	return
}

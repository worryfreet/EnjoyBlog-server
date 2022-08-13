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
	var queryReq *request.MyArticleListReq
	err = utils.FillModel(&queryReq, req)
	if err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	queryReq.UserId = l.ctx.Value("userId").(string)
	articles, err := l.svcCtx.ArticleModel.FindMyList(l.ctx, queryReq)
	if err != nil {
		l.Logger.Error("查询我的文章列表失败, err: ", err)
		return nil, errorx.StatusErrSystemBusy
	}
	resp = new(types.ArticleListResp)
	err = utils.FillModel(&resp.ArticleList, articles)
	if err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	return
}

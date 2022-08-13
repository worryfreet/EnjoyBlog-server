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

type GetArticleListByGroupIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleListByGroupIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListByGroupIdLogic {
	return &GetArticleListByGroupIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleListByGroupIdLogic) GetArticleListByGroupId(req *types.ArticleListByGroupIdReq) (resp *types.ArticleListResp, err error) {
	var queryReq *request.ArticleListByGroupIdReq
	err = utils.FillModel(&queryReq, req)
	if err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	tokenUserId := l.ctx.Value("userId").(string)
	pub := 1
	if tokenUserId == req.UserId {
		pub = 0
	}
	articles, err := l.svcCtx.ArticleModel.FindListByGroupId(l.ctx, queryReq, pub)
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

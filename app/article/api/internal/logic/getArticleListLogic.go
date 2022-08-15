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

type GetArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListLogic {
	return &GetArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleListLogic) GetArticleList(req *types.ArticleListReq) (resp *types.ArticleListResp, err error) {
	var queryReq *request.ArticleListReq
	if err = utils.FillModel(&queryReq, req); err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	tokenUserId := global.Jwt.Claims["userId"]
	pub := 1
	if tokenUserId == req.UserId {
		pub = 0
	}
	articles, err := l.svcCtx.ArticleModel.FindList(l.ctx, queryReq, pub)
	if err != nil {
		l.Logger.Error("查询我的文章列表失败, err: ", err)
		return nil, errorx.StatusErrSystemBusy
	}
	resp = new(types.ArticleListResp)
	if err = utils.FillModel(&resp.ArticleList, articles); err != nil {
		l.Logger.Error("FillModel err: ", err)
		return nil, errorx.StatusErrParam
	}
	return
}

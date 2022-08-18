package logic

import (
	"EnjoyBlog/app/article/api/internal/global"
	"EnjoyBlog/common/errorx"
	"context"

	"EnjoyBlog/app/article/api/internal/svc"
	"EnjoyBlog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteArticleReq) error {
	userId := global.Jwt.Claims["userId"].(string)
	err := l.svcCtx.ArticleModel.DeleteWithUserId(l.ctx, userId, req.ArticleId)
	if err != nil {
		l.Logger.Error("删除文章失败, err: ", err)
		return errorx.StatusErrSystemBusy
	}
	return nil
}

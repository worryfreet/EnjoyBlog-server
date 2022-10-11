package logic

import (
	"EnjoyBlog/app/comment/global"
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/request"
	"context"

	"EnjoyBlog/app/comment/api/internal/svc"
	"EnjoyBlog/app/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SupportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSupportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SupportLogic {
	return &SupportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SupportLogic) Support(req *types.SupportReq) error {
	cmtInfo := &request.CmtSupport{
		CmtId: req.CommentId,
		IP:    global.IP,
	}
	err := l.svcCtx.CommentModel.RedisSupport(l.ctx, cmtInfo)
	if err != nil {
		l.Logger.Error("点赞评论失败 err: ", err)
		return errorx.StatusErrSystemBusy
	}
	return nil
}

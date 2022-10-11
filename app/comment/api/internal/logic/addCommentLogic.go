package logic

import (
	"EnjoyBlog/app/comment/global"
	"EnjoyBlog/app/comment/model"
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/request"
	"EnjoyBlog/common/utils"
	"context"

	"EnjoyBlog/app/comment/api/internal/svc"
	"EnjoyBlog/app/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCommentLogic) AddComment(req *types.AddCommentReq) error {
	reqModel := new(model.Comment)
	err := utils.FillModel(&reqModel, req)
	if err != nil {
		l.Logger.Error("FillModel err: ", err)
		return errorx.StatusErrParam
	}
	reqModel.CommentId = utils.NewObjID()
	cmtInfo := &request.CmtSupport{
		CmtId: reqModel.CommentId,
		IP:    global.IP,
	}
	err = l.svcCtx.CommentModel.AddComment(l.ctx, reqModel, cmtInfo)
	if err != nil {
		l.Logger.Error("添加评论失败 err: ", err)
		return errorx.StatusErrSystemBusy
	}
	return nil
}

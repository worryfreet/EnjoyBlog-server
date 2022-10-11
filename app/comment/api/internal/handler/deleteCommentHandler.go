package handler

import (
	"EnjoyBlog/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"EnjoyBlog/app/comment/api/internal/logic"
	"EnjoyBlog/app/comment/api/internal/svc"
	"EnjoyBlog/app/comment/api/internal/types"
)

func DeleteCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteCommentLogic(r.Context(), svcCtx)
		err := l.DeleteComment(&req)
		response.Response(w, nil, err)
	}
}

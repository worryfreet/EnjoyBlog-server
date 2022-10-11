package handler

import (
	"EnjoyBlog/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"EnjoyBlog/app/comment/api/internal/logic"
	"EnjoyBlog/app/comment/api/internal/svc"
	"EnjoyBlog/app/comment/api/internal/types"
)

func supportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SupportReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSupportLogic(r.Context(), svcCtx)
		err := l.Support(&req)
		response.Response(w, nil, err)
	}
}

package handler

import (
	"EnjoyBlog/common/response"
	"net/http"

	"EnjoyBlog/app/article/api/internal/logic"
	"EnjoyBlog/app/article/api/internal/svc"
	"EnjoyBlog/app/article/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchArticleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchArticleListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearchArticleListLogic(r.Context(), svcCtx)
		resp, err := l.SearchArticleList(&req)
		response.Response(w, resp, err)
	}
}

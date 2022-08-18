package handler

import (
	"EnjoyBlog/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"EnjoyBlog/app/article/api/internal/logic"
	"EnjoyBlog/app/article/api/internal/svc"
	"EnjoyBlog/app/article/api/internal/types"
)

func AddArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddArticleInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddArticleLogic(r.Context(), svcCtx)
		err := l.AddArticle(&req)
		response.Response(w, nil, err)
	}
}

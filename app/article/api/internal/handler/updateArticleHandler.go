package handler

import (
	"EnjoyBlog/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"EnjoyBlog/app/article/api/internal/logic"
	"EnjoyBlog/app/article/api/internal/svc"
	"EnjoyBlog/app/article/api/internal/types"
)

func UpdateArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateArticleInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateArticleLogic(r.Context(), svcCtx)
		err := l.UpdateArticle(&req)
		response.Response(w, nil, err)
	}
}

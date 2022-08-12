package article

import (
	"EnjoyBlog/common/response"
	"net/http"

	"EnjoyBlog/app/admin/api/internal/logic/article"
	"EnjoyBlog/app/admin/api/internal/svc"
	"EnjoyBlog/app/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetArticleInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := article.NewGetArticleInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetArticleInfo(&req)
		response.Response(w, resp, err)
	}
}

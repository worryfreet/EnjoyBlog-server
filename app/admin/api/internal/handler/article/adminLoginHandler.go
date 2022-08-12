package article

import (
	"net/http"

	"EnjoyBlog/app/admin/api/internal/logic/article"
	"EnjoyBlog/app/admin/api/internal/svc"
	"EnjoyBlog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AdminLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := article.NewAdminLoginLogic(r.Context(), svcCtx)
		resp, err := l.AdminLogin(&req)
		response.Response(w, resp, err)
	}
}

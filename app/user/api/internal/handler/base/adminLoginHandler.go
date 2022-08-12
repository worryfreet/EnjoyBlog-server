package base

import (
	"EnjoyBlog/common/response"
	"net/http"

	"EnjoyBlog/app/user/api/internal/logic/base"
	"EnjoyBlog/app/user/api/internal/svc"
	"EnjoyBlog/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AdminLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := base.NewAdminLoginLogic(r.Context(), svcCtx)
		resp, err := l.AdminLogin(&req)
		response.Response(w, resp, err)
	}
}

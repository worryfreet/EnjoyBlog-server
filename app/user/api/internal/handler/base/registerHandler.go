package base

import (
	"EnjoyBlog/common/response"
	"net/http"

	"EnjoyBlog/app/user/api/internal/logic/base"
	"EnjoyBlog/app/user/api/internal/svc"
	"EnjoyBlog/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := base.NewRegisterLogic(r.Context(), svcCtx)
		err := l.Register(&req)
		response.Response(w, nil, err)
	}
}

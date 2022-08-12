package user

import (
	"EnjoyBlog/common/response"
	"net/http"

	"EnjoyBlog/app/user/api/internal/logic/user"
	"EnjoyBlog/app/user/api/internal/svc"
	"EnjoyBlog/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUpdateUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserInfo(&req)
		response.Response(w, resp, err)
	}
}

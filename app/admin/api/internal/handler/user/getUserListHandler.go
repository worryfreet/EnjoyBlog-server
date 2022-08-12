package user

import (
	"EnjoyBlog/common/response"
	"net/http"

	"EnjoyBlog/app/admin/api/internal/logic/user"
	"EnjoyBlog/app/admin/api/internal/svc"
	"EnjoyBlog/app/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGetUserListLogic(r.Context(), svcCtx)
		resp, err := l.GetUserList(&req)
		response.Response(w, resp, err)
	}
}

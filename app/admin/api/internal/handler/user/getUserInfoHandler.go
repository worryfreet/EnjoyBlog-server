package user

import (
	"EnjoyBlog/app/user/rpc/userclient"
	"EnjoyBlog/common/response"
	"net/http"

	"EnjoyBlog/app/admin/api/internal/logic/user"
	"EnjoyBlog/app/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userclient.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo(&req)
		response.Response(w, resp, err)
	}
}

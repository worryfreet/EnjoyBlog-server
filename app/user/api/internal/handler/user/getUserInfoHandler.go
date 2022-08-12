package user

import (
	"EnjoyBlog/common/response"
	"net/http"

	"EnjoyBlog/app/user/api/internal/logic/user"
	"EnjoyBlog/app/user/api/internal/svc"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()
		response.Response(w, resp, err)
	}
}

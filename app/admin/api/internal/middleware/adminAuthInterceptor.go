package middleware

import (
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/response"
	"net/http"
)

type AdminAuth struct {
}

func NewAdminAuth() *AdminAuth {
	return &AdminAuth{}
}

func (m *AdminAuth) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value("userInfo").(map[string]interface{})
		if userInfo["IsAdmin"] == 1 {
			next(w, r)
		}
		response.Response(w, nil, errorx.StatusErrAdminAuth)
	}
}

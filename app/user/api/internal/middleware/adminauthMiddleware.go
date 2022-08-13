package middleware

import (
	"EnjoyBlog/common/errorx"
	"EnjoyBlog/common/response"
	"encoding/json"
	"net/http"
)

type AdminAuthMiddleware struct {
}

func NewAdminAuthMiddleware() *AdminAuthMiddleware {
	return &AdminAuthMiddleware{}
}

func (m *AdminAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ia, _ := r.Context().Value("userInfo").(json.Number).Int64()
		if ia == 1 {
			next(w, r)
		}
		response.Response(w, nil, errorx.StatusErrAdminAuth)
	}
}

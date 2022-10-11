package middleware

import (
	"EnjoyBlog/app/comment/global"
	"net/http"
)

type SingleIP struct {
}

func NewSingleIP() *SingleIP {
	return &SingleIP{}
}

func (m *SingleIP) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		global.IP = r.Host

		next(w, r)
	}
}

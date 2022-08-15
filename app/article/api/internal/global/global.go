package global

import (
	"EnjoyBlog/app/article/api/internal/config"
	"github.com/form3tech-oss/jwt-go"
)

var Jwt JWT

type JWT struct {
	config.JwtAuth
	Claims jwt.MapClaims
}

func InitJwt(auth config.JwtAuth) {
	Jwt.Claims = make(jwt.MapClaims, 10)
	Jwt.JwtAuth = auth
}

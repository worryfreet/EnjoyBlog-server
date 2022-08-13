package middleware

import (
	"EnjoyBlog/app/article/api/internal/global"
	"EnjoyBlog/common/errorx"
	"github.com/form3tech-oss/jwt-go"
	"net/http"
)

type ParseJwtToken struct {
}

func NewParseJwtToken() *ParseJwtToken {
	return &ParseJwtToken{}
}

func (m *ParseJwtToken) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		j := NewJWT()
		if token == "" {
			global.JwtClaims["userId"] = "-1"
			next(w, r)
		}
		claims, err := j.ParseToken(token[7:]) // 去除token附加开头
		if err != nil {
			global.JwtClaims["userId"] = "-1"
			return
		}
		global.JwtClaims = claims
		next(w, r)
	}
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("worryfreet"),
	}
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, errorx.StatusErrTokenNotValid
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errorx.StatusErrTokenNotValid
}

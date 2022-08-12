// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	base "EnjoyBlog/app/user/api/internal/handler/base"
	user "EnjoyBlog/app/user/api/internal/handler/user"
	"EnjoyBlog/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: base.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: base.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/base"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/updateInfo",
				Handler: user.UpdateUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/updatePwd",
				Handler: user.UpdatePwdHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.adminAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/adminLogin",
					Handler: base.AdminLoginHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/base"),
	)
}

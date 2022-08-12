package main

import (
	"EnjoyBlog/app/admin/api/internal/config"
	"EnjoyBlog/app/admin/api/internal/handler"
	"EnjoyBlog/app/admin/api/internal/middleware"
	"EnjoyBlog/app/admin/api/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	// 调用全局admin权限拦截器中间件
	server.Use(middleware.NewAdminAuth().Handle)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

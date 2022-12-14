package main

import (
	"EnjoyBlog/app/article/api/internal/middleware"
	"flag"
	"fmt"

	"EnjoyBlog/app/article/api/internal/config"
	"EnjoyBlog/app/article/api/internal/handler"
	"EnjoyBlog/app/article/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/article.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	server.Use(middleware.NewParseJwtToken().Handle)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

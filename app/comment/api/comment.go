package main

import (
	"EnjoyBlog/app/comment/api/internal/config"
	"EnjoyBlog/app/comment/api/internal/handler"
	"EnjoyBlog/app/comment/api/internal/middleware"
	"EnjoyBlog/app/comment/api/internal/svc"
	"EnjoyBlog/app/comment/cronJob"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/comment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	server.Use(middleware.NewSingleIP().Handle)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	// 执行定时任务
	cronJob.CronTimer()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

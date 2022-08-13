package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth    JwtAuth
	Mysql      Mysql
	CacheRedis cache.CacheConf
	UserRPC    zrpc.RpcClientConf
}

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}
type Mysql struct {
	DataSource string
}

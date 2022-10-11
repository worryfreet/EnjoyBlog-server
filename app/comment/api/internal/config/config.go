package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	JwtAuth    JwtAuth
	Mysql      Mysql
	CacheRedis cache.CacheConf
}

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}
type Mysql struct {
	DataSource string
}

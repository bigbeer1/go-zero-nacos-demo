package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf

	NacosService struct {
		Host                string
		Port                int64
		NamespaceId         string
		Timeout             int64
		NotLoadCacheAtStart bool
		LogDir              string
		CacheDir            string
		LogLevel            string
	}
}

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ElasticsearchConf struct {
	Addresses []string
	Username  string
	Password  string
}

type Config struct {
	zrpc.RpcServerConf
	Mongo struct {
		URL string
		DB  string
	}
	Cache         cache.CacheConf
	Elasticsearch ElasticsearchConf
	Redis         redis.RedisConf
}

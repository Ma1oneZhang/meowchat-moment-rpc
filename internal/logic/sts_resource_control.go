package logic

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/url"
	"sync"
)

var (
	mu  sync.Mutex
	cli *redis.Redis
)

const (
	delayQueue = "sts:dq:timeOutObjectUrl"
)

// checkSingletonRedis singleton redis pattern
func checkSingletonRedis(redisConf *redis.RedisConf) {
	if cli == nil {
		mu.Lock()
		if cli == nil {
			cli = redis.MustNewRedis(*redisConf)
		}
		mu.Unlock()
	}
}

// removeUsedUrls 对于string数组使用
func removeUsedUrls(redisConf *redis.RedisConf, urls []string) {
	checkSingletonRedis(redisConf)
	for _, val := range urls {
		nUrl, _ := url.Parse(val)
		_, _ = cli.Zrem(delayQueue, nUrl.Path)
	}
}

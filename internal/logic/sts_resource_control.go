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
	usedUrl    = "sts:store:usedUrl"
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

// addUrlsToUsedUrl 对于string数组使用
func addUrlsToUsedUrl(redisConf *redis.RedisConf, urls []string) {
	checkSingletonRedis(redisConf)
	for _, val := range urls {
		nUrl, _ := url.Parse(val)
		_, _ = cli.Sadd(usedUrl, nUrl.Path)
	}
}

// addUnExistUrlToUsedUrl 当更新某项url才使用，较为耗费IO
func addUnExistUrlToUsedUrl(redisConf *redis.RedisConf, urls []string) {
	checkSingletonRedis(redisConf)
	for _, val := range urls {
		pUrl, _ := url.Parse(val)
		_, err := cli.Zrank(delayQueue, pUrl.Path)
		if err == nil {
			_, _ = cli.Sadd(usedUrl, pUrl.Path)
		}
	}
}

package util

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

const REDIS_KEY = "slide_redis"

var rClient *redis.Client

/**
滑动窗口计数器-集群版 基于redis
*/
type CounterRedis struct {
	rate  int           //计数周期内最多允许的请求数
	cycle time.Duration //计数周期

	itemSlice []int64
	lock      sync.Mutex
}

func (l *CounterRedis) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	return true
}

func (l *CounterRedis) Set(r int, cycle time.Duration) {
	if r <= 0 || cycle <= 0 {
		panic("参数异常")
	}

	l.rate = r
	l.cycle = cycle

	l.itemSlice = make([]int64, 0)
	redisClient()
}

func redisClient() {
	rClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rClient.Close()

	// 检测心跳
	_, err := rClient.Ping(context.Background()).Result()
	if err != nil {
		panic("connect redis failed")
	}
}

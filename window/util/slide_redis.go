package util

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"sync"
	"time"
)

const REDIS_KEY = "slide_redis"

var rClient *redis.Client

/**
滑动窗口计数器-集群版 基于redis
*/
type counterRedis struct {
	rate  int           //计数周期内最多允许的请求数
	cycle time.Duration //计数周期

	itemSlice []int64
	lock      sync.Mutex
}

// new
func NewSlideRedis(r int, cycle time.Duration) *counterRedis {
	if r <= 0 || cycle <= 0 {
		panic("参数异常")
	}

	re := &counterRedis{
		rate:  r,
		cycle: cycle,
	}
	re.itemSlice = make([]int64, 0)

	redisClient()

	return re
}

func redisClient() {
	rClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// 检测心跳
	_, err := rClient.Ping(context.Background()).Result()
	if err != nil {
		panic("connect redis failed")
	}

	fmt.Println("redis ok")
	os.Exit(0)
}

package util

import (
	"sync"
	"time"
)

const REDIS_KEY = "slide_redis"

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
	return re
}

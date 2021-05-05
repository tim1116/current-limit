package util

import (
	"current-limit/common"
	redigo "github.com/garyburd/redigo/redis"
	"strconv"
	"sync"
	"time"
)

const REDIS_KEY = "slide_redis"

/**
滑动窗口计数器-集群版 基于redis
*/
type CounterRedis struct {
	rate  int64         //计数周期内最多允许的请求数
	cycle time.Duration //计数周期

	lock sync.Mutex
}

func (l *CounterRedis) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	// 当前时间
	now := time.Now().UnixNano()

	redis := common.Get()
	count, err := redis.Do("llen", REDIS_KEY)
	if err != nil {
		panic(err)
	}
	countList, _ := count.(int64)
	if countList < l.rate {
		// list没有满 插入
		_, err := redis.Do("rpush", REDIS_KEY, now)
		if err != nil {
			panic(err)
		}
		return true
	}
	timeStart, err := redigo.Values(redis.Do("lrange", REDIS_KEY, 0, 0))

	timeStartI, _ := timeStart[0].([]uint8)
	timeStartS := common.B2S(timeStartI)
	tint64, _ := strconv.ParseInt(timeStartS, 10, 64)
	// 时间范围内 拒绝
	if now-tint64 <= l.cycle.Nanoseconds() {
		return false
	}

	// 删掉第一个 重新插入当前时间
	redis.Do("lpop", REDIS_KEY)
	redis.Do("rpush", REDIS_KEY, now)
	return true
}

func (l *CounterRedis) Set(r int64, cycle time.Duration) {
	if r <= 0 || cycle <= 0 {
		panic("参数异常")
	}

	l.rate = r
	l.cycle = cycle

}

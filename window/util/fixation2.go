package util

import (
	"sync"
	"time"
)

// 请求次数
var countLock = 0

/**
固定窗口计数器2
*/
type Counter struct {
	rate  int           //计数周期内最多允许的请求数
	begin time.Time     //计数开始时间
	cycle time.Duration //计数周期
	count int           //计数周期内累计收到的请求数
	lock  sync.Mutex
}

func (l *Counter) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	countLock++

	if l.count >= l.rate {
		// 时间判断
		now := time.Now()
		if now.Sub(l.begin) > l.cycle {
			//速度允许范围内， 重置计数器
			l.Reset(now)
			// 重置的时候也算一次请求
			l.count++
			return true
		} else {
			return false
		}
	} else {
		//没有达到速率限制，计数加1
		l.count++
		return true
	}
}

func (l *Counter) Set(r int, cycle time.Duration) {
	if r <= 0 || cycle <= 0 {
		panic("参数异常")
	}

	l.rate = r
	l.begin = time.Now()
	l.cycle = cycle
	l.count = 0
}

func (l *Counter) Reset(t time.Time) {
	l.begin = t
	l.count = 0
}

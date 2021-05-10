package leak

import (
	"sync"
	"time"
)

/**
漏桶算法
*/

var countLeak = 0

type Bucketleak struct {
	rate     int           //计数周期内最多允许的请求数
	cycle    time.Duration //计数周期
	count    int           //计数周期内累计收到的请求数
	lastTime time.Time

	lock sync.Mutex
}

func (l *Bucketleak) Set(r int, cycle time.Duration) {
	if r <= 0 || cycle <= 0 {
		panic("参数异常")
	}

	l.rate = r
	l.cycle = cycle
}

func (l *Bucketleak) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	return true
}

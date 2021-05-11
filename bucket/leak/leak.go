package leak

import (
	"math"
	"sync"
	"time"
)

/**
漏桶算法
*/

var countLeak = 0

type Bucketleak struct {
	rate     float64 //每秒出水速率  eg: 多少毫升  毫升/每秒
	capacity float64 //容量		 eg:毫升
	water    float64 // 当前水容量
	lastTime time.Time

	lock sync.Mutex
}

func (l *Bucketleak) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now()
	// 先漏水
	// 当前容量减去时间段内进入的流量
	timeLimit := now.Sub(l.lastTime)
	waterLimit := float64(timeLimit) / 1e9 * l.rate
	l.water = l.water - waterLimit
	l.water = math.Max(0, l.water) // 最小为0

	l.lastTime = now
	if l.water+1 <= l.capacity {
		// 漏桶没有满 可以继续加水
		l.water++
		return true
	} else {
		// 水满了
		return false
	}
}

func (l *Bucketleak) Set(rate, capacity float64) {
	if rate <= 0 || capacity <= 0 {
		panic("参数异常")
	}

	// 漏桶速率
	l.rate = rate
	// 漏桶容量
	l.capacity = capacity
	l.lastTime = time.Now()
}

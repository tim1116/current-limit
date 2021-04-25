package util

import (
	"sync/atomic"
	"time"
)

/**
固定窗口计数器-基于给定时间范围
精确度最差
*/
type CounterFix struct {
	count   int64         // 当前次数
	maxNum  int64         // 次数限制
	durTime time.Duration // 间隔时间
	endTime time.Time     //限制结束时间
}

// 设置计数器
// num / td 限制当前时间范围的次数 eg 1/time.Second 每秒1次
func (c *CounterFix) Set(num int64, td time.Duration) {
	c.maxNum = num
	c.durTime = td
	c.setEndTime()
}

// 重置结束时间
func (c *CounterFix) setEndTime() {
	c.endTime = time.Now().Add(c.durTime)
}

// 判断是否满足限流器要求
func (c *CounterFix) Allow() bool {
	now := time.Now()

	// 在规定时间内 判断请求次数
	if now.Before(c.endTime) {
		atomic.AddInt64(&c.count, 1)
		if c.count > c.maxNum {
			//  重置计数器
			return false
		} else {
			return true
		}
	} else {
		c.count = 1
		c.setEndTime()
		return true
	}
}

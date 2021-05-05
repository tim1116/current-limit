package main

import (
	"current-limit/common"
	"current-limit/window/util"
	"strconv"
	"sync"
	"time"
)

// 活动窗口限流
func main() {
	//testOne()
	testRedis()
}

// 集群版本测试
func testRedis() {
	var lr util.CounterRedis
	var wg sync.WaitGroup

	lr.Set(3, time.Second)

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(i int) {
			common.FmtShow("创建请求:" + strconv.Itoa(i))
			if lr.Allow() {
				common.Logger().Info("响应请求:" + strconv.Itoa(i))
			}
			wg.Done()
		}(i)

		time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()
}

// 单机版本test
func testOne() {
	var wg sync.WaitGroup
	var lr util.CounterSlide

	// 设置限流规则  比如这里1s1次
	lr.Set(1, time.Second)

	// 超过限制则不响应
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(i int) {
			common.FmtShow("创建请求:" + strconv.Itoa(i))
			if lr.Allow() {
				common.Logger().Info("响应请求:" + strconv.Itoa(i))
			}
			wg.Done()
		}(i)

		time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()
}

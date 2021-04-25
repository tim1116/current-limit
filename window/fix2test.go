package main

import (
	"current-limit/common"
	"current-limit/window/util"
	"strconv"
	"sync"
	"time"
)

// 固定窗口限流-单机版
func main() {
	var wg sync.WaitGroup
	var lr util.Counter

	// 设置限流规则  比如这里1s3次
	lr.Set(1, time.Second)

	// 超过限制则不响应
	for i := 1; i <= 30; i++ {
		wg.Add(1)
		go func(i int) {
			common.FmtShow("创建请求:" + strconv.Itoa(i))
			if lr.Allow() {
				common.Logger().Info("响应请求:" + strconv.Itoa(i))
			}
			wg.Done()
		}(i)

		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}

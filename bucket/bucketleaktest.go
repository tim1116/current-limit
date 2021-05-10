package main

import (
	"current-limit/bucket/leak"
	"current-limit/common"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var lr leak.Bucketleak

	lr.Set(2, time.Second)

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

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

	// 漏桶速率 1秒1毫升  漏桶1毫升
	lr.Set(1, 1)

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

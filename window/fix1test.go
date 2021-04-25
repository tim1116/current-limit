package main

import (
	"current-limit/common"
	"current-limit/window/util"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	counter := &util.CounterFix{}
	// 0.5S 1 次
	counter.Set(1, 500*time.Millisecond)

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(i int) {
			msg := "进入" + strconv.Itoa(i)
			if counter.Allow() {
				msg += "----->yes"
			} else {
				msg += " no"
			}
			common.FmtShow(msg)
			wg.Done()
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()

}

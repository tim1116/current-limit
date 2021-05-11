package main

import (
	"current-limit/bucket/leak"
	"current-limit/common"
	"current-limit/window/util"
	"fmt"
	"time"
)

// 设置频率 1S5次
func main() {
	var timeDura []time.Duration
	// 第0.1S一次
	timeDura = append(timeDura, 100*time.Millisecond)
	timeDura = append(timeDura, 700*time.Millisecond)
	timeDura = append(timeDura, 1*time.Millisecond)
	timeDura = append(timeDura, 1*time.Millisecond)
	timeDura = append(timeDura, 1*time.Millisecond)
	timeDura = append(timeDura, 1*time.Millisecond)
	timeDura = append(timeDura, 1*time.Millisecond)
	timeDura = append(timeDura, 100*time.Millisecond)
	timeDura = append(timeDura, 100*time.Millisecond)
	timeDura = append(timeDura, 1*time.Millisecond)
	timeDura = append(timeDura, 1*time.Millisecond)

	testFix1(timeDura)
	fmt.Println("---------------")
	testFix2(timeDura)
	fmt.Println("---------------")
	testSlide(timeDura)
	fmt.Println("---------------")
	testSlideRedis(timeDura)
	fmt.Println("--- 漏桶算法 -----")
	testBucketLeak(timeDura)
}

// 输出检查结果
func check(cc bool) {
	var msg string
	if cc {
		msg = " --->yes"
	} else {
		msg = " -- no"
	}
	common.FmtShow(msg)
}

//测试固定窗口1
func testFix1(t []time.Duration) {
	counter := &util.CounterFix{}
	counter.Set(5, time.Second)

	check(counter.Allow())
	for _, v := range t {
		time.Sleep(v)
		check(counter.Allow())
	}
}

//测试固定窗口2
func testFix2(t []time.Duration) {
	var lr util.Counter
	lr.Set(5, time.Second)

	check(lr.Allow())
	for _, v := range t {
		time.Sleep(v)
		check(lr.Allow())
	}
}

// 测试滑动窗口计数器
func testSlide(t []time.Duration) {
	var lr util.CounterSlide
	lr.Set(5, time.Second)

	check(lr.Allow())
	for _, v := range t {
		time.Sleep(v)
		check(lr.Allow())
	}
}

// 滑动窗口redis版本
func testSlideRedis(t []time.Duration) {
	var lr util.CounterRedis
	lr.Set(5, time.Second)

	check(lr.Allow())
	for _, v := range t {
		time.Sleep(v)
		check(lr.Allow())
	}
}

// 漏桶算法测试
func testBucketLeak(t []time.Duration) {
	var lr leak.Bucketleak
	lr.Set(5, 1)

	check(lr.Allow())
	for _, v := range t {
		time.Sleep(v)
		check(lr.Allow())
	}
}

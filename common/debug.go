package common

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

// 变量输出
func Show(v interface{}) {
	log.Printf("%T --> %v\n", v, v)
}

// 精确时间
func FmtShow(v interface{}) {
	now := time.Now()
	ss := now.Format(TIME_FORMAT)
	mill := (now.UnixNano() / 1e6) % 1000
	ss = ss + "." + strconv.FormatInt(mill, 10)
	fmt.Println(ss, v)
}

// 格式化当前时间戳 包含毫秒
func ShowTime(tt time.Time) {
	ss := tt.Format(TIME_FORMAT)
	mill := (tt.UnixNano() / 1e6) % 1000
	ss = ss + "." + strconv.FormatInt(mill, 10)
	Show(ss)
}

// 格式化显示time数组
func ShowTimeArr(tArr []time.Time) {
	var re []string
	for _, v := range tArr {
		ss := v.Format(TIME_FORMAT)
		mill := (v.UnixNano() / 1e6) % 1000
		ss = ss + "." + strconv.FormatInt(mill, 10)
		re = append(re, ss)
	}

	for _, v := range re {
		Show(v)
	}
}

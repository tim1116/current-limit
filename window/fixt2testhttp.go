package main

import (
	"current-limit/window/util"
	"fmt"
	"log"
	"net/http"
	"time"
)

var lr util.Counter

func test(w http.ResponseWriter, r *http.Request) {
	if !lr.Allow() {
		fmt.Fprintf(w, "not allow")
	} else {
		fmt.Fprintf(w, "hello")
	}
}

// 固定窗口限流 http客户端请求
func main() {
	//创建监听端口
	http.HandleFunc("/current-limit", test) //设置访问路由

	// 设置限流规则
	lr.Set(1, 3*time.Second)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe:	", err)
	}
}

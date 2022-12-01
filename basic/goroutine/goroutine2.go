package main

import (
	"fmt"
	"time"
)

var ch chan int

func recv() {
	rec := <-ch
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " 接收成功:", rec)
}

func main() {
	ch = make(chan int)
	go recv()
	ch <- 10
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " 发送成功")
}

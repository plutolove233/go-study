package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{}, 10)
	select {
	case <-ch:
		fmt.Println("get info from channel")
	case <-time.After(3 * time.Second):
		fmt.Println("time out")
	}
}

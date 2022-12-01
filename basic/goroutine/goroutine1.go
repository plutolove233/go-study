package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("this hello goroutine", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	fmt.Println("this is main goroutine")
	wg.Wait()
}

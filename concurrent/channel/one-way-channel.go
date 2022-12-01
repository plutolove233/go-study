package main

import "fmt"

func Producer() <- chan int{// only send data
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++{
			if i%2 == 1{
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func Consumer(ch <-chan int) int { // only receive
	sum := 0
	for v := range ch{
		sum += v
	}
	return sum
}

func main() {
	ch2 := Producer()
	res := Consumer(ch2)
	fmt.Println(res)
}

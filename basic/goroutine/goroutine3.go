package main

import "fmt"

func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("channel has been closed")
			break
		}
		fmt.Println(v)
	}
	// another method
	//for v := range ch {
	//	fmt.Println(v)
	//}
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	f2(ch)
}

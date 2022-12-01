package main

import "fmt"

var ch1 chan int

func recv(c chan int) {
	x := <- c // receive data from channel
	fmt.Println("received success, data is %d",x)
}

func main() {
	fmt.Println(ch1)
	ch1 = make(chan int,1)// channel must be initialized by make function
	fmt.Println(ch1)

	ch := make(chan int)// unbuffered channel
	go recv(ch)
	ch <- 10			// unbuffered channel will make deadlock
	// because the item in channel must be received
	// so we should build one receive function
	fmt.Println("send successes")
}

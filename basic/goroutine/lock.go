package main

import (
	"fmt"
	"sync"
)

var WG sync.WaitGroup
var x = 0
var lock sync.Mutex

func add() {

	for i := 0; i < 5000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}

	WG.Done()
}

func main() {
	WG.Add(2)

	go add()
	go add()

	WG.Wait()
	fmt.Println(x)
}

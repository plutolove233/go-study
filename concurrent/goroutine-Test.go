/*
@Coding : utf-8
@Time : 2022/4/9 16:14
@Author : 刘浩宇
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

func Hello(){
	defer wt.Done()// use sync.Wait() method
	fmt.Println("Hello")
}

var wt sync.WaitGroup

func main() {
	wt.Add(1)
	go Hello()// make one goroutine need time, so "main goroutine done!" will print earlier
	fmt.Println("main goroutine done!")
	// main goroutine ends too fast, so it may lead to one condition like only write "main goroutine done!"
	// So we should add one line code like "sleep" so that make main goroutine ends later than Hello goroutine
	//time.Sleep(2*time.Second)
	// But such method is not elegant, so we may use sync.WaitGroup() to wait for "Hello" goroutine stop
	wt.Wait()
}
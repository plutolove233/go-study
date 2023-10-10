package main

import (
	"flag"
	"fmt"
	"strings"
)

var files = flag.String("file", "asdasd", "evaluate file path")

var NReduce = 10

func main() {
	println("master node is running....")
	flag.Parse()
	println(*files)
	filePath := strings.Split(*files, "_")
	for _, v := range filePath {
		println(v)
	}

	m := NewMaster(filePath, NReduce)

	for !m.Done() {
	}

	fmt.Println("已完成")
}

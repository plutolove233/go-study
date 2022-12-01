// Package args
/*
@Coding : utf-8
@time : 2022/7/23 17:16
@Author : yizhigopher
@Software : GoLand
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " "))
}

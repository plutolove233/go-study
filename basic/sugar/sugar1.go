package main

import "fmt"

func test1(args ...string) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	strss := []string{
		"qwe",
		"234",
		"yui",
	}
	test1(strss...)
}

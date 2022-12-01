// Package rpcxDemo
/*
@Coding : utf-8
@time : 2022/10/18 16:56
@Author : yizhigopher
@Software : GoLand
*/
package main

import (
	"github.com/smallnest/rpcx/server"
	example "rpcxDemo"
)

func main() {
	s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("tcp", ":8972")
}

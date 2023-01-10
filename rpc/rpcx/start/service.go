package main

import (
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"rpcx"
)

var addr = flag.String("addrCliet", "localhost:8972", "server address")

func main() {
	fmt.Println("service start...")
	flag.Parse()
	s := server.NewServer()
	err := s.Register(new(rpcx.Arith), "")
	if err != nil {
		fmt.Println("register service failed, err = ", err.Error())
		return
	}
	err = s.Serve("tcp", *addr)
	if err != nil {
		fmt.Println("service start failed, err = ", err.Error())
		return
	}
}

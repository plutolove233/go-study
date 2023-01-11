/*
几台服务器提供相同的服务
*/
package main

import (
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"rpcx"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "arith service")
	addr2 = flag.String("addr2", "localhost:9981", "arith service")
)

func main() {
	flag.Parse()
	go func() {
		fmt.Println("arith1 service start...")
		s := server.NewServer()
		s.Register(&rpcx.Arith{}, "")
		s.Serve("tcp", *addr1)
	}()

	go func() {
		fmt.Println("arith2 service start...")
		s := server.NewServer()
		s.Register(&rpcx.Arith{}, "")
		s.Serve("tcp", *addr2)
	}()
	select {}
}

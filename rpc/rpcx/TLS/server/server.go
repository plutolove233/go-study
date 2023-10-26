package main

import (
	"crypto/tls"
	"flag"
	"github.com/smallnest/rpcx/server"
	"log"
	"rpcx"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	s := server.NewServer(server.WithTLSConfig(config))
	s.Register(&rpcx.Arith{}, "")
	s.Serve("tcp", *addr)
}

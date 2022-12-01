package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen TCP failed, err: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("listen client failed, err: ", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	}
}

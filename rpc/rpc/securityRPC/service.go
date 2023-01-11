package main

import (
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, replay *string) error
}

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func main() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", "1234")
	if err != nil {
		log.Fatal("listen TCP failed, err=", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go rpc.ServeConn(conn)
	}
}

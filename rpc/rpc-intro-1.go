package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloServie struct {

}

func (p *HelloServie) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService",new(HelloServie))

	listener,err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal("Listen TCP failed,error:",err.Error())
		return
	}

	conn,err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:",err)
	}

	rpc.ServeConn(conn)
}

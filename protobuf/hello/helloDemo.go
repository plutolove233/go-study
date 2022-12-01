package main

import (
	"hello/pb"
)

type HelloService struct{}

func (h *HelloService) Hello(request *pb.String, reply *pb.String) error {
	reply.Value = "hello: " + request.GetValue()
	return nil
}

func main() {

}

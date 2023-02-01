package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo/pb/models"
	"grpc-demo/pb/services/hello"
	"io"
	"net"
)

type server struct {
	hello.UnimplementedClientStreamServer
}

func (s *server) SayHi(stream hello.ClientStream_SayHiServer) error {
	reply := "你好: "
	for true {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&models.HelloResponse{Reply: reply})
			}
			return err
		}
		reply += res.GetName()
	}
	return nil
}

func main() {
	lis, _ := net.Listen("tcp", ":8972")

	s := grpc.NewServer()
	hello.RegisterClientStreamServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		fmt.Println("load service failed")
		return
	}

	fmt.Println("server started...")
}

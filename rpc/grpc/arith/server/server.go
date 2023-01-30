package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo/pb/Arith"
	"net"
)

type server struct {
	Arith.UnimplementedArithServiceServer
}

func (s *server) Add(ctx context.Context, req *Arith.ArithRequest) (*Arith.ArithResponse, error) {
	return &Arith.ArithResponse{
		C: req.GetA() + req.GetB(),
	}, nil
}

func (s *server) Mul(ctx context.Context, req *Arith.ArithRequest) (*Arith.ArithResponse, error) {
	return &Arith.ArithResponse{
		C: req.GetA() * req.GetB(),
	}, nil
}

func main() {
	fmt.Println("arith serve starting...")
	lis, err := net.Listen("tcp", ":8971")
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}

	serve := grpc.NewServer()
	Arith.RegisterArithServiceServer(serve, &server{})

	err = serve.Serve(lis)
	if err != nil {
		fmt.Println("failed to serve: ", err)
		return
	}
}

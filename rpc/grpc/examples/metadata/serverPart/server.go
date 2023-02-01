package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-demo/pb/services/arith"
	"net"
)

type server struct {
	arith.UnimplementedArithServiceServer
}

func (s *server) Add(ctx context.Context, req *arith.ArithRequest) (*arith.ArithResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.DataLoss, "fail get metadata")
	}
	if t, ok := md["token"]; ok {
		if t[0] != "api" {
			return nil, status.Errorf(4001, "authorize failed")
		}
	}

	header := metadata.MD{
		"location": []string{"China", "Jiangsu"},
	}

	grpc.SendHeader(ctx, header)
	return &arith.ArithResponse{C: req.GetA() + req.GetB()}, nil
}

func main() {
	fmt.Println("arith serve starting...")
	lis, err := net.Listen("tcp", ":8971")
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}

	serve := grpc.NewServer()
	arith.RegisterArithServiceServer(serve, &server{})

	err = serve.Serve(lis)
	if err != nil {
		fmt.Println("failed to serve: ", err)
		return
	}
}

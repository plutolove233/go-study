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

func valid(authorization []string) bool {
	return authorization[0] == "ncepu-token"
}

func authorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["token"]) {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	m, err := handler(ctx, req)
	return m, err
}

type server struct {
	arith.UnimplementedArithServiceServer
}

func (s *server) Add(ctx context.Context, req *arith.ArithRequest) (*arith.ArithResponse, error) {
	header := metadata.MD{
		"location": []string{"China", "Jiangsu"},
	}

	grpc.SetHeader(ctx, header)
	return &arith.ArithResponse{C: req.GetA() + req.GetB()}, nil
}

func main() {
	fmt.Println("arith serve starting...")
	lis, err := net.Listen("tcp", ":8971")
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}

	serve := grpc.NewServer(grpc.UnaryInterceptor(authorInterceptor))
	arith.RegisterArithServiceServer(serve, &server{})

	err = serve.Serve(lis)
	if err != nil {
		fmt.Println("failed to serve: ", err)
		return
	}
}

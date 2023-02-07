package rpc

import (
	"context"
	"fmt"
	"gateway/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(c context.Context, req *pb.Req) (*pb.Resp, error) {
	return &pb.Resp{
		Message: "hello " + req.GetArgs(),
	}, nil
}

func GrpcServe() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("failed to listen, error=", err)
		return
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})

	log.Println("Serve gRPC on 0.0.0.0:9000")
	go func() {
		log.Fatal(s.Serve(listen))
	}()

	//conn, err := grpc.DialContext(
	//	context.Background(),
	//	"0.0.0.0:9000",
	//	grpc.WithBlock(),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)
	//if err != nil {
	//	log.Fatalln("fail to dial server, err=", err)
	//}
	//
	//gwmux := runtime.NewServeMux()
	//err = pb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	//if err != nil {
	//	log.Fatalln("failed to register gateway, error=", err)
	//}
	//
	//gwServer := &http.Server{
	//	Addr:    ":9001",
	//	Handler: gwmux,
	//}
	//log.Println("Serving gRPC-gateway on http://0.0.0.0:9001")
	//log.Fatalln(gwServer.ListenAndServe())
}

package api

import (
	"context"
	"gateway/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func RegisterOnHttp() {
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:9000",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("fail to dial server, err=", err)
	}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("failed to register gateway, error=", err)
	}

	gwServer := &http.Server{
		Addr:    ":9001",
		Handler: gwmux,
	}
	log.Println("Serving gRPC-gateway on http://0.0.0.0:9001")
	log.Fatalln(gwServer.ListenAndServe())
}

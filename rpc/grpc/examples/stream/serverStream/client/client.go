package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/pb/models"
	"grpc-demo/pb/services/hello"
	"io"
	"log"
	"time"
)

var addr = flag.String("server addr", "localhost:8972", "server stream demo")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := hello.NewServerStreamClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.SayHi(ctx, &models.HelloRequest{Name: "yizhigopher"})
	if err != nil {
		log.Fatalf("could not call arith add service: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("client say hi failed, err=%v\n", err)
		}
		log.Printf("got reply: %q\n", res.GetReply())
	}
}

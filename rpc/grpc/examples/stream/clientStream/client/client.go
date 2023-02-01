package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/pb/models"
	"grpc-demo/pb/services/hello"
	"log"
	"time"
)

var addr = flag.String("server addr", "localhost:8972", "client stream demo")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := hello.NewClientStreamClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, _ := c.SayHi(ctx)
	names := []string{"Qimi", "yizhigopher"}

	for _, name := range names {
		err := stream.Send(&models.HelloRequest{
			Name: name,
		})
		if err != nil {
			log.Fatalf("stream send %s failed, err=%v", name, err)
		}
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("say hi failed: %v", err)
	}
	log.Printf("got reply: %v", res.GetReply())
}

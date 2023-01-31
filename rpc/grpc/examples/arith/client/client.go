package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/pb/services/arith"
	"log"
	"time"
)

var addr = flag.String("arithAddr", "localhost:8971", "arith server address")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := arith.NewArithServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Add(ctx, &arith.ArithRequest{
		A: 10,
		B: 20,
	})
	if err != nil {
		log.Fatalf("could not call arith add service: %v", err)
	}

	log.Printf("result: %d\n", r.C)
}

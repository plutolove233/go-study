package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"grpc-demo/pb/services/arith"
	"log"
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

	md := metadata.New(map[string]string{
		"token": "api",
	})

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	var header, trailer metadata.MD
	r, err := c.Add(
		ctx,
		&arith.ArithRequest{
			A: 10,
			B: 20,
		},
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		log.Fatalf("could not call arith add service: %v", err)
	}

	if t, ok := header["location"]; ok {
		fmt.Printf("location from header:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Printf("location expected but doesn't exist in header")
		return
	}
	log.Printf("result: %d\n", r.C)
}

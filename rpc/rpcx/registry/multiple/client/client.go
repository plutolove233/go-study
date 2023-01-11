package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	"rpcx"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "arith service")
	addr2 = flag.String("addr2", "localhost:9981", "arith service")
)

func main() {
	flag.Parse()
	d, _ := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})

	xClient := client.NewXClient("Arith", client.Failtry, client.RoundRobin, d, client.DefaultOption)
	defer xClient.Close()

	args := &rpcx.Args{
		A: 10,
		B: 20,
	}
	reply := &rpcx.Reply{}

	call, err := xClient.Go(context.Background(), "Add", args, reply, nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	c := <-call.Done
	if c.Error != nil {
		log.Fatal(c.Error.Error())
	} else {
		log.Printf("%d + %d = %d\n", args.A, args.B, reply.C)
	}
}

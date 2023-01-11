package main

import (
	"context"
	"crypto/tls"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	"rpcx"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	option := client.DefaultOption
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	option.TLSConfig = conf
	xClient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xClient.Close()

	args := &rpcx.Args{
		A: 10,
		B: 30,
	}
	reply := &rpcx.Reply{}
	err := xClient.Call(context.Background(), "Add", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}

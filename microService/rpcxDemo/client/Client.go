// Package rpcxDemo
/*
@Coding : utf-8
@time : 2022/10/18 17:11
@Author : yizhigopher
@Software : GoLand
*/
package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	example "rpcxDemo"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}

	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("%d * %d = %d\n", args.A, args.B, reply.C)
}

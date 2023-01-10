package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	"rpcx"
)

var (
	addrCliet = flag.String("addrCliet", "localhost:8972", "server address")
)

func main() {
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addrCliet, "")
	xClient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer func(xClient client.XClient) {
		err := xClient.Close()
		if err != nil {
			println("client close failed, err = ", err.Error())
		}
	}(xClient)

	req := &rpcx.Args{
		A: 10,
		B: 20,
	}
	resp := &rpcx.Reply{}

	err := xClient.Call(context.Background(), "Add", req, resp)
	if err != nil {
		log.Fatal("call service failed, err=", err.Error())
	}
}

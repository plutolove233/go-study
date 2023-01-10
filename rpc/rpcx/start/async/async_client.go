package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/log"
	"rpcx"
	"time"
)

var addrAsync = flag.String("addr", "localhost:8972", "async client caller")

func main() {
	d, err := client.NewPeer2PeerDiscovery("tcp@"+*addrAsync, "")
	if err != nil {
		log.Errorf("create peer to peer failed, err=", err.Error())
		return
	}
	xClient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer func(xClient client.XClient) {
		if err := xClient.Close(); err != nil {
			log.Fatal("close client failed, err=", err.Error())
		}
	}(xClient)

	req := &rpcx.Args{
		A: 20,
		B: 15,
	}
	i := 0

	channelPool := make(chan *client.Call, 10)
	go func() {
		for {
			i++
			resp := &rpcx.Reply{}
			fmt.Println("run service ", i)
			call, err := xClient.Go(context.Background(), "Add", req, resp, nil)
			if err != nil {
				log.Fatal("call service failed, err=", err.Error())
			}
			replyCall := <-call.Done
			channelPool <- replyCall
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case x := <-channelPool:
			if x.Error != nil {
				log.Errorf("call service failed, err=", err.Error())
			} else {
				fmt.Printf("%d + %d = %d\n", req.A, req.B, x.Reply.(*rpcx.Reply).C)
			}
		}
	}
}

// Package server
/*
@Coding : utf-8
@Time : 2023/1/30 19:45
@Author : yizhigopher
@Software : GoLand
*/
package main

import (
	model "etcd-demo"
	"flag"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/server"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
	etcdAddr = flag.String("etcd.server", "localhost:2379", "etcd server address")
	basePath = flag.String("base", "/rpcx", "prefix path")
)

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers: []string{*etcdAddr},
		BasePath: *basePath,
		UpdateInterval: time.Minute,
	}

	err := r.Start()
	if err != nil{
		log.Fatal(err)
	}

	s.Plugins.Add(r)
}

func main() {
	flag.Parse()

	//go http.ListenAndServe(":9981", nil)

	s := server.NewServer()
	addRegistryPlugin(s)

	s.Register(&model.Arith{}, "")

	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
package main

import (
	"gateway/api"
	"gateway/rpc"
)

func main() {
	rpc.GrpcServe()
	api.RegisterOnHttp()
}

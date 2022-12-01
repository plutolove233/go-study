package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
}

// 用于获取user rpc在etcd上注册服务的配置信息

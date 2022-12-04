package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-looklook/app/usercenter/cmd/api/internal/config"
	"go-looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	UserCenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserCenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
	}
}

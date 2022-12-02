package main

import (
	"demo2/apps/user/api/internal/config"
	"demo2/apps/user/api/internal/handler"
	"demo2/apps/user/api/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//httpx.SetErrorHandler(func(err error) (int, interface{}) {
	//	switch e := err.(type) {
	//	case *errorx.CodeError: // 定义了的异常则捕获
	//		return http.StatusOK, e.Data()
	//	default: // 未定义的异常报500
	//		return http.StatusInternalServerError, nil
	//	}
	//})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

package main

import (
	v3 "buf/api/v3"
	"buf/services"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"log"
)

func main() {
	httpSrv := http.NewServer(
		http.Address(":8000"),
		http.Middleware(recovery.Recovery()),
	)

	v3.RegisterBookServiceHTTPServer(httpSrv, &services.BookHandler{})
	app := kratos.New(
		kratos.Name("book"),
		kratos.Server(
			httpSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

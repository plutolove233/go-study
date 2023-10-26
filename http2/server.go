package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

const cert = "./cert/localhost.pem"
const _key = "./cert/localhost-key.pem"

func main() {
	_, err := tls.LoadX509KeyPair(cert, _key)
	if err != nil {
		fmt.Println("load X509 cert failed, err=", err.Error())
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%+v\n", r.Header.Get("Extra"))
		fmt.Println("+++++++++++++++++++++++++++++++++++++")
		b, _ := io.ReadAll(r.Body)
		fmt.Println(string(b))
		fmt.Fprintf(w, "Hello World, protocol=%s\n", r.Proto)
	})

	server := http.Server{
		ReadTimeout:  1 * time.Hour,
		WriteTimeout: 1 * time.Hour,
		Handler:      mux,
	}
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("listen tcp on 9090 failed ,err=", err.Error())
		return
	}
	server.ServeTLS(listen, cert, _key)
	select {}
}

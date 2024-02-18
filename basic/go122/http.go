package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// 支持RESTful风格参数
	// 以及可以在pattern前部添加请求类型
	mux.HandleFunc("GET /say/{name}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello "+r.PathValue("name"))
	})
	
	srv := http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("Create http service failed, err=", err.Error())
		return
	}
}

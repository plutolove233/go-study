package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.PATH = %q\n", request.URL.Path)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
	})
}

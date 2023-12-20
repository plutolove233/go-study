package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
// go env -w GOOS=wasip1
// go env -w GOARCH=wasm
// go build -o main.wasm main.go
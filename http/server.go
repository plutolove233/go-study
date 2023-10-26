package main

import (
	"fmt"
	"net"
	"strconv"
)

const (
	SOURCE_PORT_LEN = 2
	DEST_PORT_LEN   = 2
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Listen on 8080 failed, err=", err)
		return
	}
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("accept failed, err=", err)
		return
	}

	data := make([]byte, 10000)
	conn.Read(data)
	if err != nil {
		fmt.Println("read from conn failed, err=", err)
		return
	}
	fmt.Println(string(data))

	// 响应体
	var respBody = "<h1>Hello World</h1>"
	length := len(respBody)
	//响应头
	var respHeader = "HTTP/1.1 200 OK\n" +
		"Content-Type: text/html;charset=ISO-8859-1\n" +
		"Content-Length: " + strconv.FormatInt(int64(length), 10)

	resp := respHeader + "\n\r\n" + respBody
	conn.Write([]byte(resp))
}

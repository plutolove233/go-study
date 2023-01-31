/*
* 客户端发出一个RPC请求，服务端与客户端之间建立一个单向的流，服务端可以向流中写入多个响应消息，最后主动关闭流；
* 而客户端需要监听这个流，不断获取响应直到流关闭。
* 应用场景举例：客户端向服务端发送一个股票代码，服务端就把该股票的实时数据源源不断的返回给客户端。
 */

package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo/pb/models"
	"grpc-demo/pb/services/hello"
	"net"
)

type server struct {
	hello.UnimplementedServerStreamServer
}

func (s *server) SayHi(req *models.HelloRequest, stream hello.ServerStream_SayHiServer) error {
	words := []string{
		"你好",
		"Hello",
		"こんにちは",
		"안녕하세요",
	}

	for _, word := range words {
		data := &models.HelloResponse{
			Reply: word + req.GetName(),
		}

		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	fmt.Println("server starting...")
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("listening failed, err=%v\n", err.Error())
		return
	}

	s := grpc.NewServer()
	hello.RegisterServerStreamServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		fmt.Println("fail to serve", err.Error())
		return
	}
}

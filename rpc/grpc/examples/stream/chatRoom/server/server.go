package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo/pb/models"
	"grpc-demo/pb/services/chat"
	"io"
	"log"
	"net"
	"strings"
)

type chatRoom struct {
	chat.UnimplementedChatRoomServer
}

func (c *chatRoom) Chat(stream chat.ChatRoom_ChatServer) error {
	for {
		in, err := stream.Recv()
		log.Printf("get request, value=%v, error=%v", in, err)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		reply := magic(in.GetName())

		if err = stream.Send(&models.HelloResponse{Reply: reply}); err != nil {
			return err
		}
	}
}

func magic(s string) string {
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "？", "!")
	s = strings.ReplaceAll(s, "?", "!")
	return s
}

func main() {
	lis, _ := net.Listen("tcp", ":8972")

	s := grpc.NewServer()
	chat.RegisterChatRoomServer(s, &chatRoom{})

	fmt.Println("chat room starting...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("chat room server failed...")
		return
	}

}

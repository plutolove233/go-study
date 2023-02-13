package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"web-socket/chatMultiRooms/apis/ws"
)

func main() {
	engine := gin.Default()

	engine.GET("/ws/:roomId", ws.ServeChat)

	if err := engine.Run(":8080"); err != nil {
		log.Fatalln("run engine failed")
		return
	}
	log.Println("multipart rooms chat service start...")
}

package main

import (
	"github.com/gin-gonic/gin"
	"web-socket/oneRoomChat/chat"
)

var port = ":8080"

func main() {
	engine := gin.Default()

	go chat.HUB.Run()

	engine.GET("chat", chat.ServeWs)

	engine.Run(port)
}

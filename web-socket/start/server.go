package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

var addr = flag.String("ws port", "localhost:8080", "http server port")
var upGrader = websocket.Upgrader{}

func wsEcho(c *gin.Context) {
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade failed, error=", err)
		return
	}

	defer conn.Close()
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read failed, error=", err)
			break
		}
		log.Println("recv message=", msg)

		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			log.Println("write failed, error=", err)
			break
		}
	}
}

func main() {
	engine := gin.Default()

	engine.GET("/echo", wsEcho)

	if err := engine.Run(*addr); err != nil {
		log.Fatalf("run gin server failed, error=%s\n", err.Error())
		return
	}
}

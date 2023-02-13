// Package ws
/*
@Coding : utf-8
@Time : 2023/2/13 19:48
@Author : yizhigopher
@Software : GoLand
*/
package ws

import (
	"github.com/gin-gonic/gin"
	"web-socket/chatMultiRooms/globals"
	"web-socket/chatMultiRooms/models/wsModels"
)

func ServeChat(ctx *gin.Context) {
	roomId := ctx.Param("roomId")
	hub, ok := globals.HOUSE[roomId]
	if !ok {
		hub = wsModels.NewHub(roomId)
		globals.HOUSE[roomId] = hub
		go hub.Run()
	}
	wsModels.ServeWSChat(hub, ctx.Writer, ctx.Request)
}

// Package globals
/*
@Coding : utf-8
@Time : 2023/2/13 19:58
@Author : yizhigopher
@Software : GoLand
*/
package globals

import (
	"web-socket/chatMultiRooms/models/wsModels"
)

// HOUSE will store the hubs of chat system,we can search one hub by roomId
var HOUSE = make(map[string]*wsModels.Hub)

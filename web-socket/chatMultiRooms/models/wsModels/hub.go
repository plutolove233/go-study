// Package wsModels
/*
@Coding : utf-8
@Time : 2023/2/13 20:00
@Author : yizhigopher
@Software : GoLand
*/
package wsModels

import "web-socket/chatMultiRooms/globals"

type Hub struct {
	roomId     string
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub(roomId string) *Hub {
	return &Hub{
		roomId:     roomId,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	defer func() {
		close(h.register)
		close(h.broadcast)
		close(h.unregister)
	}()

	for {
		select {
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				if len(h.clients) == 0 {
					delete(globals.HOUSE, h.roomId)
					return
				}
			}
		case msg := <-h.broadcast:
			for c, _ := range h.clients {
				select {
				case c.send <- msg:
				default:
					close(c.send)
					delete(h.clients, c)
				}
			}
		}
	}
}

func (h *Hub) Unregister(client *Client) {
	h.unregister <- client
}

func (h *Hub) SetBroadcast(msg []byte) {
	h.broadcast <- msg
}

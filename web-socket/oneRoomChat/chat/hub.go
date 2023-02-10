// Package chat
/*
@Coding : utf-8
@Time : 2023/2/10 21:22
@Author : yizhigopher
@Software : GoLand
*/
package chat

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from clients.
	broadcast chan []byte

	// register request from clients.
	register chan *Client

	// unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

var HUB = NewHub()

func (h *Hub) Run() {
	for true {
		select {
		case xClient := <-h.register:
			h.clients[xClient] = true
		case xClient := <-h.unregister:
			if _, ok := h.clients[xClient]; ok {
				delete(h.clients, xClient)
				close(xClient.Send)
			}
		case msg := <-h.broadcast:
			for xClient, _ := range h.clients {
				select {
				case xClient.Send <- msg:
				default:
					close(xClient.Send)
					delete(h.clients, xClient)
				}
			}
		}
	}
}

func (h *Hub) Register(client2 *Client) {
	h.register <- client2
}

func (h *Hub) Unregister(client2 *Client) {
	h.unregister <- client2
}

func (h *Hub) SetBroadcast(msg []byte) {
	h.broadcast <- msg
}

package main

import (
	// "github.com/gorilla/websocket"
	// "appliedgo.net/what"

	"encoding/json"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	clientIdRecord int
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Register(client *Client) {
	client.SetId(h.clientIdRecord)
	h.clientIdRecord++
	h.register <- client
}

func (h *Hub) BroadcastObj(obj interface{}) {
	bs, err := json.Marshal(obj)
	if err != nil {
		logrus.Warn("marshal obj failed, no broadcasting")
		spew.Dump(obj)
		return
	}
	h.Broadcast(bs)
}

func (h *Hub) Broadcast(bs []byte) {
	h.broadcast <- bs
	logrus.Debugf("--> %s", string(bs))
}

func (h *Hub) UnRegister(client *Client) {
	h.unregister <- client
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			// logrus.Info()
			logrus.Infof("client: %d registed", client.id)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

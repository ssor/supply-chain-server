package main

import (
	"bytes"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub
	id  int
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send     chan []byte
	IdNumber int
	actor    ActorId
}

func (c *Client) SetId(id int) {
	c.id = id
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.UnRegister(c)
		// c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		logrus.Infof("client (%d) <- %s", c.id, string(message))
		jsonParsed, err := gabs.ParseJSON(message)
		if err != nil {
			logrus.Warnf("message parse error: %s", err)
			continue
		}
		messageType, ok := jsonParsed.Path("type").Data().(float64)
		if !ok {
			spew.Dump(messageType)
			spew.Dump(ok)
			logrus.Warnf("message protocol error")
			continue
		}
		t := MessageType(messageType)
		switch t {
		case RoleJoinMessageType:
			number, ok := jsonParsed.Path("roleID").Data().(float64)
			if !ok {
				logrus.Warnf("message protocol error")
				spew.Dump(number)
				continue
			}
			c.IdNumber = int(number)
			c.actor = RoleIdToActor(c.IdNumber)
			jsonParsed.Set(200, "maxInventory")
			jsonParsed.Set(100, "currentInventory")

			switch c.actor {
			case NoSuchActor:
				logrus.Warnf("no such actor occured from number: %d", number)
			case Producer:
				
			case LevelOneDispather:
			case LevelTwoDispather:
			case Detailer:
				logrus.Infof("client(%d) is a %s", c.IdNumber, c.actor.String())
			}
		default:
			logrus.Warnf("unsupported message type: %d", messageType)
		}
		jsonParsed.Set(RoleJoinResponseMessageType, "type")
		jsonParsed.Set(game.Id, "gameID")
		jsonParsed.Set(game.StartTime.Unix(), "gameStartTime")
		c.hub.Broadcast([]byte(jsonParsed.String()))
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				logrus.Warnf("client(%d) closed", c.id)
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			logrus.Infof("client(%d) -> %s", c.id, string(message))

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

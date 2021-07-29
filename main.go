package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)
var orderSubscriber = make(chan *Order,1)
var orderDispatcher = NewOrderDispatcher(orderSubscriber)
var game = NewGame()
var msgBase = NewMessageBase(game.Id,game.StartTime.Unix())
var upgrader = websocket.Upgrader{} // use default options
var hub = newHub()

func startListenOrderDispatch(){
	// orderDispatcher.Run()
	order:= <- orderSubscriber
	hub.OrderIn(order)
}
// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	// client.hub.register <- client
	hub.Register(client)
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}


func main() {
	go hub.run()
	go startListenOrderDispatch()

	router := gin.New()

	router.GET("/ws", func(context *gin.Context) {
		w, r := context.Writer, context.Request
		serveWs(hub, w, r)
	})
	router.StaticFS("/public", http.Dir("roots"))

	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}

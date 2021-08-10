package main

import (
	"bytes"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
)
var addr = flag.String("addr", "localhost:8000", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	var messageIndex int;
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("%d <---: %s",messageIndex, bytes.TrimSpace(message))
			messageIndex++
			// subMessages := bytes.Split(message, newline)
			// for _, sm := range subMessages {
			// 	trimmed := bytes.TrimSpace(sm)
			// 	if len(trimmed) > 0 {
			// 		log.Printf("<---: [%s]", bytes.TrimSpace(sm))
			// 	}
			// }
		}
	}()
	time.AfterFunc(time.Second*2, func() {
		jsonObj := gabs.New()
		jsonObj.Set(1, "type")
		jsonObj.Set(1, "roleID")
		c.WriteMessage(websocket.TextMessage, []byte(jsonObj.String()))
		log.Printf("--->: %s", jsonObj.String())
	})

	time.AfterFunc(time.Second*3, func() {
		jsonObj := gabs.New()
		jsonObj.Set(4, "type")
		jsonObj.Set(1, "roleID")
		jsonObj.Set(21, "truckID")
		c.WriteMessage(websocket.TextMessage, []byte(jsonObj.String()))
		log.Printf("--->: %s", jsonObj.String())
	})

	<-interrupt
	log.Println("interrupt")

}

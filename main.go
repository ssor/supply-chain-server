package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	// "github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
	"github.com/ispringteam/eventbus"
)


// var game = NewGame()
// var msgBase = NewMessageBase(gameState.Id, gameState.StartTime.Unix())
var upgrader = websocket.Upgrader{} // use default options
var hub = newHub()
var eventBus = eventbus.New()
var gameState = NewGameState()
// var truckMonitor *TruckMonitor = NewTruckMonitor(gameState, hub).Run()
var orderDispatcher = NewOrderDispatcher(gameState)

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
	go client.readPump(func(b []byte) {
		binaryMessageDispatch(b)
	})
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
	logrus.SetLevel(logrus.DebugLevel)
	go hub.run()
	// go startListenOrderDispatch()
	newGameEngine(gameState,hub).StartRunning()
	orderDispatcher.Run()
	// eventBus.Subscribe(eventNameBroadcast, func(e eventbus.Event) {
	// 	be := e.(*BroadcastEvent)
	// 	hub.BroadcastObj(be.Obj)
	// })
	// eventBus.Subscribe(eventNameNotifyOrderDispatcherStart, notifyOrderDispatcherStart)
	// eventBus.Subscribe(eventNameNewOrderFromCustomer, newOrderFromCustomer)
	eventBus.Subscribe(eventNameLevel1DispatcherJoin, func(event eventbus.Event) {
		e := event.(*Level1DispatherJoinEvent)
		gameState.AddLevel2Dispather(e.clientID)
	})
	eventBus.Subscribe(eventNameLevel2DispatcherJoin, func(event eventbus.Event) {
		e := event.(*Level2DispatherJoinEvent)
		gameState.AddLevel1Dispather(e.clientID)
	})
	eventBus.Subscribe(eventNameDetailerJoin, func(event eventbus.Event) {
		e := event.(*DetailerJoinEvent)
		gameState.AddDetailer(e.Id)
	})
	eventBus.Subscribe(eventNameProducerJoin, func(event eventbus.Event) {
		e := event.(*ProducerJoinEvent)
		gameState.AddProducer(e.clientID)
	})

	eventBus.Subscribe(eventNameUpdateDetailerInventory, func(event eventbus.Event) {
		e := event.(*UpdateDetailerInventoryEvent)
		gameState.UpdateDetailerInventory(e.Id, e.Inventory)
	})
	eventBus.Subscribe(eventNameResetTruckLoadAndInventory, func(event eventbus.Event) {
		e := event.(*ResetTruckLoadAndInventoryEvent)
		gameState.UpdateTruckLoad(e.truckID, e.load)
		gameState.UpdateInventory(e.clientID, e.inventory)
	})
	eventBus.Subscribe(eventNameResetTruckDestination, resetTruckDest)
	eventBus.Subscribe(eventNameresetGame, func(event eventbus.Event) {
		gameState.Reset()
		hub.BroadcastObj(gameState.BasicGameInfoMessage())
	})
	// truckMonitor.Run()

	router := gin.New()

	router.GET("/ws", func(context *gin.Context) {
		w, r := context.Writer, context.Request
		serveWs(hub, w, r)
	})
	router.GET("/game/reset", func(c *gin.Context) {
		eventBus.Publish(&ResetGameEvent{})
		c.JSON(http.StatusOK, nil)
	})
	router.StaticFS("/public", http.Dir("roots"))

	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}

func resetTruckDest(e eventbus.Event) {
	logrus.Debug("-> reset Truck Dest")
	event := e.(*ResetTruckDestinationEvent)
	// truckMonitor.AddMonitorItem(event.TruckID, event.RoleID)
	gameState.UpdateTruckDestRole(event.TruckID, event.RoleID)
}

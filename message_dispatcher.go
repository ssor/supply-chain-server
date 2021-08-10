package main

import (
	"supplychain_server/protocol"

	"github.com/Jeffail/gabs"
	"github.com/davecgh/go-spew/spew"

	// "github.com/ispringteam/eventbus"
	"github.com/sirupsen/logrus"
)

// "strconv"
// "time"

func binaryMessageDispatch(bs []byte) {
	// logrus.Info("< ---- binaryMessageDispatch")
	jsonParsed, err := gabs.ParseJSON(bs)
	if err != nil {
		logrus.Warnf("message parse error: %s", err)
		return
	}
	messageType, ok := jsonParsed.Path("type").Data().(float64)
	if !ok {
		spew.Dump(messageType)
		spew.Dump(ok)
		logrus.Warnf("message protocol error")
		return
	}
	t := protocol.MessageType(messageType)
	switch t {
	case protocol.RoleJoinMessageType:
		number, ok := jsonParsed.Path("roleID").Data().(float64)
		if !ok {
			logrus.Warnf("message protocol error")
			spew.Dump(number)
			return
		}
		var deviceName string
		data := jsonParsed.Path("device").Data()
		if data != nil {
			deviceName = data.(string)
		}
		if len(deviceName) <= 0 {
			createNewRole(int(number))
		} else {
			createNewRoleWithName(deviceName)
		}
	case protocol.ResetTruckDestinationMessageType:
		number, ok := jsonParsed.Path("roleID").Data().(float64)
		if !ok {
			logrus.Warnf("message protocol error")
			spew.Dump(number)
			return
		}
		truckNumber, ok := jsonParsed.Path("truckID").Data().(float64)
		if !ok {
			logrus.Warnf("message protocol error")
			spew.Dump(number)
			return
		}
		eventBus.Publish(&ResetTruckDestinationEvent{
			RoleID:  int(number),
			TruckID: int(truckNumber),
		})
	case protocol.InventoryUpdateMessageType:
		number, ok := jsonParsed.Path("roleID").Data().(float64)
		if !ok {
			logrus.Warnf("message protocol error for roleID")
			spew.Dump(number)
			return
		}
		inventory, ok := jsonParsed.Path("currentInventory").Data().(float64)
		if !ok {
			logrus.Warnf("message protocol error for currentInventory")
			spew.Dump(number)
			return
		}
		// gameState.UpdateDetailerInventory(int(number), int(inventory))
		truckNumber, ok := jsonParsed.Path("truckID").Data().(float64)
		if !ok {
			logrus.Warnf("message protocol error for truckID")
			spew.Dump(number)
			return
		}
		truckLoad, ok := jsonParsed.Path("tQuantity").Data().(float64)
		if !ok {
			logrus.Warnf("message protocol error for tQuantity ")
			spew.Dump(number)
			return
		}
		// gameState.UpdateTruckLoad(int(truckNumber), int(truckLoad))
		eventBus.Publish(&ResetTruckLoadAndInventoryEvent{
			clientID:  int(number),
			inventory: int(inventory),
			truckID:   int(truckNumber),
			load:      int(truckLoad),
		})
	case protocol.DetailerInventoryUpdateMessageType:
		number, ok := jsonParsed.Path("roleID").Data().(float64)
		if !ok {
			logrus.Warnf("message protocol error")
			spew.Dump(number)
			return
		}
		inventory, ok := jsonParsed.Path("currentInventory").Data().(float64)
		if !ok {
			logrus.Warnf("message protocol error")
			spew.Dump(number)
			return
		}
		eventBus.Publish(&UpdateDetailerInventoryEvent{Id: int(number), Inventory: int(inventory)})
		// gameState.UpdateDetailerInventory(int(number), int(inventory))
	default:
		logrus.Warnf("unsupported message type: %d", messageType)
	}
	// jsonParsed.Set(RoleJoinResponseMessageType, "type")
	// jsonParsed.Set(game.Id, "gameID")
	// jsonParsed.Set(game.StartTime.Unix(), "gameStartTime")
	// c.hub.Broadcast([]byte(jsonParsed.String()))
}

func createNewRoleWithName(name string) {
	switch name {
	case "admin":
		logrus.Infof("create role for device %s", name)
		for _, sprite := range gameState.gameMap.trucks {
			{
				mb := protocol.NewMessageBase(protocol.TruckRealtimeStatusMessageType, gameState.Id, gameState.StartTime.Unix())
				t := protocol.NewTruckMoveAnimationMessage(
					sprite.id,
					sprite.currentCoord,
					sprite.direction,
					animationSpeed,
					sprite.cap,
					sprite.loaded).WithMessageBase(mb)
				gameState.broadcastObjList <- t
			}
		}
	default:
		logrus.Warnf("no handler for device %s", name)
		// spew.Dump(name)
		// spew.Dump(name == "admin")
	}

}
func createNewRole(number int) {
	actor := RoleIdToActor(number)

	switch actor {
	case NoSuchActor:
		logrus.Warnf("no such actor occured from number: %d", number)
	case ProducerActor:
		logrus.Infof("id(%d) as %s try to join...", number, actor.String())
		eventBus.Publish(&ProducerJoinEvent{clientID: number})
	case LevelOneDispatherActor:
		logrus.Infof("id(%d) as %s try to join...", number, actor.String())
		eventBus.Publish(&Level1DispatherJoinEvent{clientID: number})
	case LevelTwoDispatherActor:
		logrus.Infof("id(%d) as %s try to join...", number, actor.String())
		eventBus.Publish(&Level2DispatherJoinEvent{clientID: number})
	case DetailerActor:
		eventBus.Publish(&DetailerJoinEvent{Id: number})
	}
}

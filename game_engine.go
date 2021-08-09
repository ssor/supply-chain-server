package main

import "time"

type InfoPublisher interface {
	BroadcastObj(interface{})
}

type GameEngine struct {
	firstMap  *FirstMap
	gs        *GameState
	publisher InfoPublisher
}

func newGameEngine(gs *GameState, publisher InfoPublisher) *GameEngine {
	ge := &GameEngine{
		gs:        gs,
		firstMap:  gs.gameMap,
		publisher: publisher,
	}
	return ge
}

func (ge *GameEngine) StartRunning() *GameEngine {
	go ge.truckMove()
	go ge.gameStateEventHandle()
	return ge
}

func (ge *GameEngine) gameStateEventHandle() {
	for {
		obj := <-ge.gs.broadcastObjList
		ge.publisher.BroadcastObj(obj)
	}
}

func (ge *GameEngine) truckMove() {
	ticker := time.NewTicker(truckMoveInterval)
	for {
		<-ticker.C
		for _, sprite := range ge.firstMap.trucks {
			sprite.move()
			p := sprite.destPosition
			if p.IsEmpty() {
				continue
			}
			{
				mb := NewMessageBase2(TruckInfoNotifyMessageType, ge.gs.Id, p.Id, ge.gs.StartTime.Unix())
				t := Truck{
					MessageBase:  mb,
					Id:           sprite.id,
					Loaded:       sprite.loaded,
					Cap:          sprite.cap,
					DistanceLeft: sprite.leftDistance(),
				}
				ge.publisher.BroadcastObj(t)
			}
			{
				mb := NewMessageBase(TruckMoveNextMessageType, ge.gs.Id, ge.gs.StartTime.Unix())
				t := newTruckMoveNextMessage(
					sprite.id,
					sprite.destPosition.Coord).WithMessageBase(mb)
				ge.publisher.BroadcastObj(t)

			}
		}
	}
}

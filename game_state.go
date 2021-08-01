package main

import (
	"strconv"
	"time"

	"github.com/ispringteam/eventbus"
	"github.com/sirupsen/logrus"
	// "github.com/sirupsen/logrus"
)

type GameStateMode int

const (
	CreatingMode GameStateMode = 0
	RunningMode  GameStateMode = 1
)

type InfoPublisher interface {
	Publish(eventbus.Event)
}

type GameState struct {
	Id        string
	mode      GameStateMode
	StartTime time.Time
	detailers Detailers
	orders    Orders
	publisher InfoPublisher
	trucks    Trucks
	producers Producers
	level1Dispatchers Dispatchers
	level2Dispatchers Dispatchers
}

func NewGameState(publisher InfoPublisher) *GameState {
	g := &GameState{
		publisher: publisher,
	}
	g.init()
	return g
}

func (gs *GameState) init() {
	gs.mode = CreatingMode
	gs.StartTime = time.Now()
	gs.Id = strconv.FormatInt(gs.StartTime.UnixNano(), 10)
	gs.detailers = Detailers{}
	gs.producers = Producers{}
	gs.level1Dispatchers = Dispatchers{}
	gs.level2Dispatchers = Dispatchers{}
	gs.orders = Orders{
		NewOrder(),
		NewOrder(),
		NewOrder(),
	}
	gs.trucks = Trucks{
		NewTruck(21, 10, 0, 0),
		NewTruck(22, 10, 0, 0),
		NewTruck(23, 10, 0, 0),
		NewTruck(24, 10, 0, 0),
		NewTruck(25, 10, 0, 0),
	}
}

func (gs *GameState) BasicGameInfoMessage() MessageBase {
	return NewMessageBase(GameResetMessageType, gs.Id, gs.StartTime.Unix())
}

func (gs *GameState) ProducerJoinMessage(id int) interface{} {
	d := gs.producers.Find(id)
	if d == nil {
		logrus.Warnf("producer %d not found",id)
		return nil
	}
	return d.WithBase(
		NewMessageBase(RoleJoinResponseMessageType, gs.Id, gs.StartTime.Unix()))

}
func (gs *GameState) Level2DispatcherJoinMessage(id int) interface{} {
	d := gs.level2Dispatchers.Find(id)
	if d == nil {
		return nil
	}
	return d.WithBase(
		NewMessageBase(RoleJoinResponseMessageType, gs.Id, gs.StartTime.Unix()))
}
func (gs *GameState) Level1DispatcherJoinMessage(id int) interface{} {
	d := gs.level1Dispatchers.Find(id)
	if d == nil {
		return nil
	}
	return d.WithBase(
		NewMessageBase(RoleJoinResponseMessageType, gs.Id, gs.StartTime.Unix()))
}
func (gs *GameState) DetailerJoinMessage(id int) interface{} {
	d := gs.detailers.Find(id)
	if d == nil {
		return nil
	}
	return d.WithBase(
		NewMessageBase(RoleJoinResponseMessageType, gs.Id, gs.StartTime.Unix()))
}

func (gs *GameState) AddProducer(id int) {
	p := gs.producers.Find(id)
	if p != nil {
		logrus.Warnf("producer(%d) has already exits", p.id)
		return
	}
	newProducer := NewProducer(id)
	gs.producers = gs.producers.Add(newProducer)
	gs.publisher.Publish(&BroadcastEvent{gs.ProducerJoinMessage(id)})
}

func (gs *GameState) AddLevel2Dispather(id int) {
	dispatcher := gameState.level2Dispatchers.Find(id)
	if dispatcher != nil {
		logrus.Warnf("dispatcher(%d) has already exits", dispatcher.id)
		return
	}
	newDispatcher := NewDispatcher(id)
	gs.level2Dispatchers = gs.level2Dispatchers.Add(newDispatcher)
	gs.publisher.Publish(&BroadcastEvent{gs.Level2DispatcherJoinMessage(id)})
}
func (gs *GameState) AddLevel1Dispather(id int) {
	dispatcher := gameState.level1Dispatchers.Find(id)
	if dispatcher != nil {
		logrus.Warnf("dispatcher(%d) has already exits", dispatcher.id)
		return
	}
	newDispatcher := NewDispatcher(id)
	gs.level1Dispatchers = gs.level1Dispatchers.Add(newDispatcher)
	gs.publisher.Publish(&BroadcastEvent{gs.Level1DispatcherJoinMessage(id)})
}

func (gs *GameState) AddDetailer(id int) {
	detailer := gameState.detailers.Find(id)
	if detailer != nil {
		logrus.Warnf("detailer(%d) has already exits", detailer.id)
		return
	}
	newDetailer := NewDetailer(id)
	gs.detailers = gs.detailers.Add(newDetailer)
	gs.publisher.Publish(&BroadcastEvent{gs.DetailerJoinMessage(id)})
	gs.mode = RunningMode
	// eventBus.Publish(&NotifyOrderDispatcherStartEvent{})
}

func (gs *GameState) Reset() {
	gs.init()
}

func (gs *GameState) UpdateDetailerInventory(id, inventory int) {
	gs.detailers.UpdateDetailerInventory(id, inventory)
}

func (gs *GameState) CheckTruck() {
	for _, t := range gs.trucks {
		b := t.RunTowardsDest()
		if !b {
			continue
		}
		mb := NewMessageBase(TruckInfoNotifyMessageType, gs.Id, gs.StartTime.Unix())
		gs.publisher.Publish(&BroadcastEvent{t.WithMessageBase(mb)})
	}

}
func (gs *GameState) CheckOrder() {
	if gs.mode == CreatingMode {
		return
	}

	headOrder, left := gs.orders.Head()
	if headOrder == nil {
		return
	}

	gs.orders = left

	//TODO random to select detailer
	headDetailer := gs.detailers.Head()
	if headDetailer == nil {
		return
	}
	// oe := e.(*NewOrderFromCustomerEvent)
	headDetailer.AddOrder(headOrder)
	base:= NewMessageBase2(OrderDispatch2DetailerMessageType, gs.Id,headDetailer.id, gs.StartTime.Unix())
	msg := headOrder.WithMessageBase(base)
	gs.publisher.Publish(&BroadcastEvent{msg})
}

func (gs *GameState) UpdateInventory(clientID, inventory int) {
	{
		c := gs.detailers.Find(clientID)
		if c != nil {
			gs.detailers.UpdateDetailerInventory(clientID, inventory)
			return
		}
	}

}
func (gs *GameState) UpdateTruckLoad(truckID, load int) {
	gs.trucks.UpdateTruckLoad(truckID, load)
}

func (gs *GameState) UpdateTruckDestRole(truckID, roleID int) {
	gs.trucks.UpdateTruckDestRole(truckID, roleID)
}

// func (gs *GameState) OrderIn(order *Order) {
// 	msg := NewOrderDispatch2DetailerMessage(1, order.Count,
// 		NewMessageBase(OrderDispatch2DetailerMessageType, gs.Id, gs.StartTime.Unix()))
// 	bs, err := json.Marshal(msg)
// 	if err != nil {
// 		logrus.Errorf("serialize message error: %s", err)
// 		spew.Dump(msg)
// 		return
// 	}
// 	// eventBus.Publish("broadcast-message", bs)
// 	// h.Broadcast(bs)
// 	// logrus.Infof("--> %s", string(bs))
// }

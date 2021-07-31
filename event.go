package main

import "github.com/ispringteam/eventbus"

var (
	eventNameBroadcast                  = eventbus.EventID("Broadcast")
	eventNameDetailerJoin               = eventbus.EventID("detailerJoin")
	eventNameProducerJoin               = eventbus.EventID("producerJoin")
	eventNameLevel1DispatcherJoin       = eventbus.EventID("level1DispatcherJoin")
	eventNameLevel2DispatcherJoin       = eventbus.EventID("level2DispatcherJoin")
	eventNamebinaryMessageDispatch      = eventbus.EventID("binaryMessageDispatch")
	eventNameNotifyOrderDispatcherStart = eventbus.EventID("notifyOrderDispatcherStart")
	eventNameNewOrderFromCustomer       = eventbus.EventID("newOrderFromCustomer")
	eventNameResetTruckDestination      = eventbus.EventID("resetTruckDestination")
	eventNameResetTruckLoadAndInventory = eventbus.EventID("resetTruckLoadAndInventory")
	eventNameresetGame                  = eventbus.EventID("resetGame")
	eventNameUpdateDetailerInventory    = eventbus.EventID("updateDetailerInventory")
)

type Level2DispatherJoinEvent struct {
	clientID int
}

func (Level2DispatherJoinEvent) EventID() eventbus.EventID {
	return eventNameLevel2DispatcherJoin
}

type Level1DispatherJoinEvent struct {
	clientID int
}

func (Level1DispatherJoinEvent) EventID() eventbus.EventID {
	return eventNameLevel1DispatcherJoin
}

type ProducerJoinEvent struct {
	clientID int
}

func (ProducerJoinEvent) EventID() eventbus.EventID {
	return eventNameProducerJoin
}

type ResetTruckLoadAndInventoryEvent struct {
	clientID  int
	inventory int
	truckID   int
	load      int
}

func (ResetTruckLoadAndInventoryEvent) EventID() eventbus.EventID {
	return eventNameResetTruckLoadAndInventory
}

type UpdateDetailerInventoryEvent struct {
	Id        int
	Inventory int
}

func (UpdateDetailerInventoryEvent) EventID() eventbus.EventID {
	return eventNameUpdateDetailerInventory
}

type DetailerJoinEvent struct {
	Id int
}

func (DetailerJoinEvent) EventID() eventbus.EventID {
	return eventNameDetailerJoin
}

type ResetGameEvent struct {
}

func (ResetGameEvent) EventID() eventbus.EventID {
	return eventNameresetGame
}

type ResetTruckDestinationEvent struct {
	TruckID int
	RoleID  int
}

func (e *ResetTruckDestinationEvent) EventID() eventbus.EventID {
	return eventNameResetTruckDestination
}

type BroadcastEvent struct {
	Obj interface{}
}

func (be *BroadcastEvent) EventID() eventbus.EventID {
	return eventNameBroadcast
}

type BinaryMessageDispatchEvent struct {
	Binary []byte
}

func (be *BinaryMessageDispatchEvent) EventID() eventbus.EventID {
	return eventNamebinaryMessageDispatch
}

type NotifyOrderDispatcherStartEvent struct {
}

func (be *NotifyOrderDispatcherStartEvent) EventID() eventbus.EventID {
	return eventNameNotifyOrderDispatcherStart
}

type NewOrderFromCustomerEvent struct {
	order *Order
}

func (e *NewOrderFromCustomerEvent) EventID() eventbus.EventID {
	return eventNameNewOrderFromCustomer
}

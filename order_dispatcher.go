package main

import (
	"supplychain_server/protocol"
	"time"

	"github.com/sirupsen/logrus"
)

type Order struct {
	Count int `json:"purchaseCnt"`
	protocol.MessageBase
}

func NewOrder() *Order {
	return &Order{Count: 1}
}

func (o *Order) WithMessageBase(base protocol.MessageBase) Order {
	c := Order{
		Count:       o.Count,
		MessageBase: base,
	}
	return c
}

type Orders []*Order

func (orders Orders) Head() (*Order, Orders) {
	if len(orders) <= 0 {
		return nil, Orders{}
	}
	return orders[0], orders[1:]
}

type OrderChecker interface {
	CheckOrder()
}

func NewOrderDispatcher(oc OrderChecker) *OrderDispatcher {
	od := &OrderDispatcher{
		// orders: Orders{
		// 	NewOrder(),
		// 	NewOrder(),
		// 	NewOrder(),
		// },
		// subscriber: sub,
		// gs:         gameState,
		checker: oc,
	}
	return od
}

type OrderDispatcher struct {
	checker OrderChecker
	// orders     Orders
	// subscriber func(*Order)
	// gs         *GameState
}

func (od *OrderDispatcher) run(ticker *time.Ticker) {
	logrus.Info("order dispatcher start running...")
	for {
		<-ticker.C
		if od.checker == nil {
			continue
		}
		od.checker.CheckOrder()
	}
}

func (od *OrderDispatcher) Run() *OrderDispatcher {
	ticker := time.NewTicker(3 * time.Second)
	go od.run(ticker)
	return od
}

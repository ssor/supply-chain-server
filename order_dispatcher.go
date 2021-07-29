package main

import "time"

type Order struct {
	Count int
}

func NewOrder() *Order {
	return &Order{Count: 1}
}

type Orders []*Order

func (orders Orders) Head() (*Order, Orders) {
	if len(orders) <= 0 {
		return nil, Orders{}
	}
	return orders[0], orders[1:]
}

func NewOrderDispatcher(sub chan *Order) *OrderDispatcher {
	od := &OrderDispatcher{
		orders: Orders{
			NewOrder(),
			NewOrder(),
			NewOrder(),
			NewOrder(),
			NewOrder(),
			NewOrder(),
		},
		subscriber: sub,
	}
	return od
}

type OrderDispatcher struct {
	orders     Orders
	subscriber chan *Order
}

func (od *OrderDispatcher) run(ticker *time.Ticker) {
	for {
		<-ticker.C
		head, left := od.orders.Head()
		if head == nil {
			continue
		}
		od.orders = left
		od.subscriber <- head
	}

}

func (od *OrderDispatcher) Run() {
	ticker := time.NewTicker(3 * time.Second)
	go od.run(ticker)
}

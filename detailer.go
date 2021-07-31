package main

import "github.com/sirupsen/logrus"

func NewDetailer(id int) *Detailer {
	return &Detailer{
		id:         id,
		MaxInventory:     200,
		CurrentInventory: 100,
	}
}

type Detailer struct {
	id         int `json:"-"`
	MaxInventory     int `json:"maxInventory"`
	CurrentInventory int `json:"currentInventory"`
	MessageBase
	orders Orders `json:"-"`
}

func (d *Detailer) AddOrder(order *Order) {
	d.orders = append(d.orders, order)
}

func (d *Detailer) WithBase(base MessageBase) Detailer {
	c := Detailer{
		id:         d.id,
		MaxInventory:     d.MaxInventory,
		CurrentInventory: d.CurrentInventory,
		MessageBase:      base,
	}
	c.RoleID = d.id
	return c
}

type Detailers []*Detailer

func (ds Detailers) Head() *Detailer {
	if ds == nil || len(ds) <= 0 {
		return nil
	}
	return ds[0]
}

func (ds Detailers) UpdateDetailerInventory(id, inventory int) {
	d := ds.Find(id)
	if d == nil {
		logrus.Warnf("no detailer(%d) register, update inventory failed", id)
		return
	}
	d.CurrentInventory = inventory
	logrus.Infof("detailer(%d) inventory updated to %d", id, inventory)
}

func (ds Detailers) Find(id int) *Detailer {
	for _, d := range ds {
		if d.id == id {
			return d
		}
	}
	return nil
}

func (ds Detailers) Add(d *Detailer) Detailers {
	return append(ds, d)
}

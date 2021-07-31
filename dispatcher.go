package main

import "github.com/sirupsen/logrus"

func NewDispatcher(id int) *Dispatcher {
	return &Dispatcher{
		id:               id,
		MaxInventory:     200,
		CurrentInventory: 100,
	}
}

type Dispatcher struct {
	id               int `json:"-"`
	MaxInventory     int `json:"maxInventory"`
	CurrentInventory int `json:"currentInventory"`
	MessageBase
}

func (d *Dispatcher) WithBase(base MessageBase) Dispatcher {
	c := Dispatcher{
		id:               d.id,
		MaxInventory:     d.MaxInventory,
		CurrentInventory: d.CurrentInventory,
		MessageBase:      base,
	}
	c.RoleID = d.id
	return c
}

type Dispatchers []*Dispatcher

func (ds Dispatchers) Head() *Dispatcher {
	if ds == nil || len(ds) <= 0 {
		return nil
	}
	return ds[0]
}

func (ds Dispatchers) UpdateDetailerInventory(id, inventory int) {
	d := ds.Find(id)
	if d == nil {
		logrus.Warnf("no detailer(%d) register, update inventory failed", id)
		return
	}
	d.CurrentInventory = inventory
	logrus.Infof("detailer(%d) inventory updated to %d", id, inventory)
}

func (ds Dispatchers) Find(id int) *Dispatcher {
	for _, d := range ds {
		if d.id == id {
			return d
		}
	}
	return nil
}

func (ds Dispatchers) Add(d *Dispatcher) Dispatchers {
	return append(ds, d)
}

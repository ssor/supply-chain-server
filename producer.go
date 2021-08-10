package main

import "supplychain_server/protocol"

func NewProducer(id int) *Producer {
	return &Producer{
		id: id,
	}
}

type Producer struct {
	id int `json:"-"`
	protocol.MessageBase
}

func (d *Producer) WithBase(base protocol.MessageBase) Producer {
	// t := d.RoleID
	p := Producer{
		id:          d.id,
		MessageBase: base,
	}
	p.RoleID = d.id
	return p
}

type Producers []*Producer

func (ds Producers) Head() *Producer {
	if ds == nil || len(ds) <= 0 {
		return nil
	}
	return ds[0]
}

func (ds Producers) Find(id int) *Producer {
	for _, d := range ds {
		if d.id == id {
			return d
		}
	}
	return nil
}

func (ds Producers) Add(d *Producer) Producers {
	return append(ds, d)
}

package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

func NewTruck(id int, cap, loaded, distance int) *Truck {
	t := &Truck{
		Id: id, DistanceLeft: distance,
		Cap: cap, Loaded: loaded,
	}
	return t
}

type Truck struct {
	MessageBase
	Id           int `json:"truckID"`
	DistanceLeft int `json:"remainingDistance"`
	Cap          int `json:"tMaxQuantity"`
	Loaded       int `json:"tQuantity"`
}

func (t *Truck) WithMessageBase(base MessageBase) Truck {
	role := t.RoleID
	ot := Truck{
		MessageBase: base,
		Id:          t.Id, DistanceLeft: t.DistanceLeft, Cap: t.Cap, Loaded: t.Loaded,
	}
	ot.RoleID = role
	return ot

}

func (t *Truck) RunTowardsDest() bool {
	if t.DistanceLeft == 0 {
		return false
	}

	t.DistanceLeft -= 5
	if t.DistanceLeft < 0 {
		t.DistanceLeft = 0
	}
	return true
}

func (t *Truck) ResetDistance(distance int) {
	t.DistanceLeft = distance
}

type Trucks []*Truck

func (trucks Trucks) Find(truckID int) *Truck {
	for _, t := range trucks {
		if t.Id == truckID {
			return t
		}
	}
	return nil
}

func (trucks Trucks) UpdateTruckLoad(truckID, load int) {
	t := trucks.Find(truckID)
	if t == nil {
		logrus.Warnf("no truck %d found, update load failed", truckID)
		return
	}
	t.Loaded = load
	logrus.Infof("truck %d  update loaded to %d", truckID, load)
}

func (trucks Trucks) UpdateTruckDestRole(truckID, roleID int) {
	for _, t := range trucks {
		if t.Id == truckID {
			t.RoleID = roleID
			t.DistanceLeft = 50
			return
		}
	}
	logrus.Warnf("no truck %d found", truckID)
}

type TruckChecker interface {
	CheckTruck()
}

func NewTruckMonitor(checker TruckChecker) *TruckMonitor {
	moniotr := &TruckMonitor{
		checker: checker,
	}

	return moniotr
}

type TruckMonitor struct {
	checker TruckChecker
}

// func (monitor *TruckMonitor) UpdateTruckLoad(truckID, load int) {
// 	monitor.trucks.UpdateTruckLoad(truckID, load)
// }

// func (monitor *TruckMonitor) UpdateTruckDestRole(truckID, roleID int) *TruckMonitor {
// 	monitor.trucks.UpdateTruckDestRole(truckID, roleID)
// 	return monitor
// }
func (monitor *TruckMonitor) Run() *TruckMonitor {
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		for {
			<-ticker.C
			if monitor.checker != nil {
				monitor.checker.CheckTruck()
			}

		}
	}()
	return monitor
}

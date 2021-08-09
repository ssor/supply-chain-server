package main

import (

	"github.com/sirupsen/logrus"
)


func NewFirstMap() *FirstMap {
	return &FirstMap{
		positions: Positions{},
		trucks:    TruckSprites{},
	}
}

type FirstMap struct {
	width, height int
	positions     Positions
	trucks        TruckSprites
}

func (fm *FirstMap) init() *FirstMap {
	fm.width = 1280
	fm.height = 720

	fm.trucks = fm.trucks.add(21, 10, Coord{X: 0, Y: 183}, East)
	fm.trucks = fm.trucks.add(22, 10, Coord{X: 0, Y: 183}, East)
	fm.trucks = fm.trucks.add(23, 10, Coord{X: 0, Y: 183}, East)
	fm.trucks = fm.trucks.add(24, 10, Coord{X: 0, Y: 183}, East)
	fm.trucks = fm.trucks.add(25, 10, Coord{X: 0, Y: 183}, East)

	fm.positions = fm.positions.add(1, 100, 103)
	fm.positions = fm.positions.add(6, 400, 103)
	fm.positions = fm.positions.add(11, 860, 103)
	fm.positions = fm.positions.add(16, 1200, 103)

	// fm.run()
	return fm
}

func (fm *FirstMap) FindWithDestRole(truckID, role int) *TruckSprite {
	return fm.trucks.FindWithDestRole(truckID, role)
}

func (fm *FirstMap) updateTruckDestRole(truckID, roleID int) {
	truck := fm.trucks.Find((truckID))
	if truck == nil {
		logrus.Warnf("truck %d not found on map", truckID)
		return
	}
	p := fm.positions.find(roleID)
	if !p.Maybe {
		logrus.Warnf("position: %d not found on map", roleID)
		return
	}
	truck.addDestPosition(p.Position)
}

// func (fm *FirstMap) run() {
// 	ticker := time.NewTicker(time.Second * 3)
// 	go func() {
// 		for {
// 			<-ticker.C
// 			for _, sprite := range fm.trucks {
// 				sprite.move()
// 			}
// 		}
// 	}()
// }
func (fm *FirstMap) Find(id int) MaybePosition {
	return fm.positions.find(id)
}

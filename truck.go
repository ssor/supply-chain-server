package main

import (
	"container/list"
	"math"

	"github.com/sirupsen/logrus"
)

type TruckReachDestEvent struct {
	id int
	Position
	cap    int
	loaded int
}

func NewTruck(id int, cap, loaded, distance int) *Truck {
	t := &Truck{
		Id: id, DistanceLeft: distance,
		Cap: cap, Loaded: loaded,
	}
	return t
}

func NewTruckSprite(id int, cap int, c Coord, d Direction) *TruckSprite {
	t := &TruckSprite{
		id:           id,
		cap:          cap,
		events:       list.New(),
		currentCoord: c,
		direction:    d,
		nextPositions: *newPositionLine(),
	}
	return t
}

type TruckSprite struct {
	id            int
	cap           int
	loaded        int
	currentCoord  Coord
	destPosition  Position
	nextPositions PositionLine
	direction     Direction
	events        *list.List
}

func (ts *TruckSprite) addDestPosition(c Position) {
	ts.nextPositions.AddPosition(c)
	logrus.Debugf("truck(%d) on map add dest:  %s", ts.id, c)
}

func (ts *TruckSprite) setNextDest() {
	nextPosition := ts.nextPositions.Shift()
	if !nextPosition.Maybe{
		ts.destPosition = newEmptyPosition(ts.destPosition.X, ts.destPosition.Y)
		logrus.Debugf("truck %d no next position to move to", ts.id)
		return
	}
	if nextPosition.Coord.Equal(ts.currentCoord) {
		// ts.nextPositions = ts.nextPositions[1:]
		return
	}
	ts.destPosition = nextPosition.Position
	// ts.nextPositions = ts.nextPositions[1:]
	//switch direction
	if ts.destPosition.X > ts.currentCoord.X {
		ts.direction = East
		logrus.Infof("truck %d set next dest to %s: %s", ts.id, ts.destPosition, ts.direction)
		return
	} else if ts.destPosition.X < ts.currentCoord.X {
		ts.direction = West
		logrus.Infof("truck %d set next dest to %s: %s", ts.id, ts.destPosition, ts.direction)
		return
	}
	if ts.destPosition.Y > ts.currentCoord.Y {
		ts.direction = North
		logrus.Infof("truck %d set next dest to %s: %s", ts.id, ts.destPosition, ts.direction)
		return
	} else if ts.destPosition.Y < ts.currentCoord.Y {
		ts.direction = South
		logrus.Infof("truck %d set next dest to %s: %s", ts.id, ts.destPosition, ts.direction)
		return
	}
	ts.direction = NoDirection
	logrus.Warnf("truck %d direction update failure, current: %s  dest: %s", ts.id, ts.currentCoord, ts.destPosition.Coord)
}

func (ts *TruckSprite) pushEvent() {
	ts.events.PushBack(TruckReachDestEvent{
		id:       ts.id,
		cap:      ts.cap,
		loaded:   ts.loaded,
		Position: ts.destPosition,
	})
}

func (ts *TruckSprite) move() {
	switch ts.direction {
	case North:
		// logrus.Debugf("truck %d moving towards north", ts.id)
		if ts.destPosition.Y < ts.currentCoord.Y {
			ny := ts.currentCoord.Y - DefaultTruckSpeed
			if ny < ts.destPosition.Y { //reached
				ny = ts.destPosition.Y
				ts.currentCoord = newCoord(ts.currentCoord.X, ny)
				ts.pushEvent()
			} else {
				ts.currentCoord = newCoord(ts.currentCoord.X, ny)
			}
		} else {
			ts.setNextDest()
		}
	case South:
		// logrus.Debugf("truck %d moving towards south", ts.id)
		if ts.destPosition.Y > ts.currentCoord.Y {
			ny := ts.currentCoord.Y + DefaultTruckSpeed
			if ny > ts.destPosition.Y { //reached
				ny = ts.destPosition.Y
				ts.currentCoord = newCoord(ts.currentCoord.X, ny)
				ts.pushEvent()
			} else {
				ts.currentCoord = newCoord(ts.currentCoord.X, ny)
			}
		} else {
			ts.setNextDest()
		}
	case West:
		if ts.destPosition.X < ts.currentCoord.X {
			nx := ts.currentCoord.X - DefaultTruckSpeed
			if nx < ts.destPosition.X { //reached
				nx = ts.destPosition.X
				ts.currentCoord = newCoord(nx, ts.currentCoord.Y)
				ts.pushEvent()
			} else {
				ts.currentCoord = newCoord(nx, ts.currentCoord.Y)
			}
		} else {
			ts.setNextDest()
		}
	case East:
		if ts.destPosition.X > ts.currentCoord.X {
			nx := ts.currentCoord.X + DefaultTruckSpeed
			if nx > ts.destPosition.X { //reached
				nx = ts.destPosition.X
				ts.currentCoord = newCoord(nx, ts.currentCoord.Y)
				ts.pushEvent()
			} else {
				ts.currentCoord = newCoord(nx, ts.currentCoord.Y)
			}
		} else {
			ts.setNextDest()
		}
	default:
		logrus.Warnf("used uninitialized direction to move")
		ts.setNextDest()
	}
}
func (ts *TruckSprite) leftDistance() (ld int) {
	switch ts.direction {
	case West, East, South, North:
		fld1 := math.Abs(float64(ts.destPosition.X - ts.currentCoord.X))
		fld2 := math.Abs(float64(ts.destPosition.Y - ts.currentCoord.Y))
		ld = int(fld1 + fld2)
	case NoDirection:
		logrus.Infof("truck %d now no-direction set", ts.id)
	default:
		logrus.Warnf("used uninitialized direction to compute left distance")
	}
	return
}

type TruckSprites []*TruckSprite

func (tss TruckSprites) add(id, cap int, coord Coord, d Direction) TruckSprites {
	return append(tss, NewTruckSprite(id, cap, coord, d))
}

func (tss TruckSprites) FindWithDestRole(truckID, role int) *TruckSprite {
	logrus.Debugf("FindWithDestRole -> truck: %d role: %d", truckID, role)
	for _, t := range tss {
		if t.id != truckID {
			continue
		}
		if t.destPosition.Id != role {
			continue
		}
		return t
	}
	return nil
}
func (tss TruckSprites) Find(truckID int) *TruckSprite {
	for _, t := range tss {
		if t.id == truckID {
			return t
		}
	}
	return nil
}

func (tss TruckSprites) UpdateTruckLoad(truckID, load int) {
	t := tss.Find(truckID)
	if t == nil {
		logrus.Warnf("no truck %d found, update load failed", truckID)
		return
	}
	t.loaded = load
	logrus.Infof("truck %d  update loaded to %d", truckID, load)
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

// type TruckRunningStatus struct {
// 	MessageBase
// 	Id        int       `json:"truckID"`
// 	Cap       int       `json:"tMaxQuantity"`
// 	Loaded    int       `json:"tQuantity"`
// 	From      Coord     `json:"from"`
// 	To        Coord     `json:"to"`
// 	Current   Coord     `json:"current"`
// 	Direction Direction `json:"direction"`
// 	Speed     int       `json:"speed"`
// }

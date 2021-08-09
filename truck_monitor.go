package main

import (

)

// type TruckMonitorItem struct {
// 	Role    int
// 	TruckId int
// }

// func NewTruckMonitor(gs *GameState, publisher InfoPublisher) *TruckMonitor {
// 	moniotr := &TruckMonitor{
// 		gs:        gs,
// 		publisher: publisher,
// 		items:     []TruckMonitorItem{},
// 		chAddItem: make(chan TruckMonitorItem, 8),
// 		// toAdded:   *newTrackedMonitorItems(),
// 	}

// 	return moniotr
// }

// type TruckMonitor struct {
// 	gs        *GameState
// 	publisher InfoPublisher
// 	items     []TruckMonitorItem
// 	chAddItem chan TruckMonitorItem
// }

// func (monitor *TruckMonitor) UpdateTruckLoad(truckID, load int) {
// 	monitor.trucks.UpdateTruckLoad(truckID, load)
// }

// func (monitor *TruckMonitor) AddMonitorItem(truckID, roleID int) {
// 	item := TruckMonitorItem{Role: roleID, TruckId: truckID}
// 	monitor.chAddItem <- item
// 	// monitor.toAdded.AddMonitorItem(truckID, roleID)
// 	logrus.Debugf("add monitor Item: truck: %d  role: %d", truckID, roleID)
// }

// func (monitor *TruckMonitor) realtimePositionNotify() {
// 	items := []TruckMonitorItem{}
// 	for _, item := range monitor.items {
// 		realTruck := monitor.gs.gameMap.trucks.FindWithDestRole(item.TruckId, item.Role)
// 		if realTruck == nil {
// 			logrus.Infof("truck %d -> position(%d) to be removed", item.TruckId, item.Role)
// 			continue
// 		}
// 		items = append(items, item)
// 		mb := NewMessageBase2(TruckInfoNotifyMessageType, monitor.gs.Id, item.Role, monitor.gs.StartTime.Unix())
// 		t := Truck{
// 			MessageBase:  mb,
// 			Id:           item.TruckId,
// 			Loaded:       realTruck.loaded,
// 			Cap:          realTruck.cap,
// 			DistanceLeft: realTruck.leftDistance(),
// 		}
// 		monitor.publisher.BroadcastObj(t)
// 	}
// 	monitor.items = items
// }
// func (monitor *TruckMonitor) Run() *TruckMonitor {
// 	ticker := time.NewTicker(2 * time.Second)
// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				monitor.realtimePositionNotify()
// 			case item := <-monitor.chAddItem:
// 				monitor.items = append(monitor.items, item)
// 			}
// 		}
// 	}()
// 	return monitor
// }

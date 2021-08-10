package protocol

type MessageType int

const (
	DefaultMessageType                 MessageType = 0
	RoleJoinMessageType                MessageType = 1
	RoleJoinResponseMessageType        MessageType = 2
	OrderDispatch2DetailerMessageType  MessageType = 3
	ResetTruckDestinationMessageType   MessageType = 4
	TruckInfoNotifyMessageType         MessageType = 5
	InventoryUpdateMessageType         MessageType = 6
	DetailerInventoryUpdateMessageType MessageType = 7
	GameResetMessageType               MessageType = 8
	TruckRealtimeStatusMessageType     MessageType = 9
	TruckMoveNextMessageType           MessageType = 10
)

func NewTruckMoveNextMessage(truck int,  dest Coord ) TruckMoveNextMessage {
	return TruckMoveNextMessage{
		Id:        truck,
		To: dest,
	}
}

func (msg TruckMoveNextMessage) WithMessageBase(base MessageBase) TruckMoveNextMessage {
	msg.MessageBase = base
	return msg
}
type TruckMoveNextMessage struct {
	MessageBase
	Id     int   `json:"truckID"`
	To     Coord `json:"to"`
}
type TruckMoveAnimationMessage struct {
	MessageBase
	Id     int   `json:"truckID"`
	Cap    int   `json:"tMaxQuantity"`
	Loaded int   `json:"tQuantity"`
	Current   Coord     `json:"position"`
	Direction Direction `json:"direction"`
	Speed     int       `json:"speed"`
}

func NewTruckMoveAnimationMessage(truck int,  current Coord, d Direction, speed, cap, loaded int) TruckMoveAnimationMessage {
	return TruckMoveAnimationMessage{
		Id:        truck,
		Cap:       cap,
		Loaded:    loaded,
		Current: current,
		Direction: d,
		Speed:     speed,
	}
}
func (msg TruckMoveAnimationMessage) WithMessageBase(base MessageBase) TruckMoveAnimationMessage {
	msg.MessageBase = base
	return msg
}

func NewMessageBase(t MessageType, game string, start int64) MessageBase {
	return MessageBase{
		Type:          t,
		GameID:        game,
		GameStartTime: start,
	}
}

func NewMessageBase2(t MessageType, game string, role int, start int64) MessageBase {
	return MessageBase{
		Type:          t,
		GameID:        game,
		GameStartTime: start,
		RoleID:        role,
	}
}

type MessageBase struct {
	Type          MessageType `json:"type"`
	GameID        string      `json:"gameID"`
	GameStartTime int64       `json:"gameStartTime"`
	RoleID        int         `json:"roleID"`
}

// type OrderDispatch2DetailerMessage struct {
// 	MessageBase
// 	PurchaseCnt int `json:"purchaseCnt"`
// }

// func NewOrderDispatch2DetailerMessage(role int, count int, base MessageBase) OrderDispatch2DetailerMessage {
// 	msg := OrderDispatch2DetailerMessage{}
// 	msg.RoleID = role
// 	msg.Type = OrderDispatch2DetailerMessageType
// 	msg.MessageBase = base
// 	msg.PurchaseCnt = count
// 	return msg
// }

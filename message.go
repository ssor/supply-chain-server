package main

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
)

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
		RoleID: role,
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

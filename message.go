package main

type MessageType int

const (
	DefaultMessageType                MessageType = 0
	RoleJoinMessageType               MessageType = 1
	RoleJoinResponseMessageType       MessageType = 2
	OrderDispatch2DetailerMessageType MessageType = 3
)

func NewMessageBase(game string, start int64) MessageBase {
	return MessageBase{
		GameID:        game,
		GameStartTime: start,
	}
}

type MessageBase struct {
	Type          MessageType `json: "type"`
	RoleID        int         `json: "roleID"`
	GameID        string      `json: "gameID"`
	GameStartTime int64       `json: "gameStartTime"`
}

type OrderDispatch2DetailerMessage struct {
	MessageBase
	PurchaseCnt int `json: "purchaseCnt"`
}

func NewOrderDispatch2DetailerMessage(role int, count int, base MessageBase) OrderDispatch2DetailerMessage {
	msg := OrderDispatch2DetailerMessage{}
	base.RoleID = role
	base.Type = OrderDispatch2DetailerMessageType
	msg.MessageBase = base
	msg.PurchaseCnt = count
	return msg
}

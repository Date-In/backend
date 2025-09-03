package chat

import (
	"dating_service/internal/model"
	"github.com/gorilla/websocket"
)

type ChatStorage interface {
	SaveMessage(*model.Message) error
	GetMessageHistory(uint, int) ([]*model.Message, error)
}

type MatchProvider interface {
	IsUserInMatch(uint, uint) (bool, error)
	GetUsers(uint) ([]model.User, error)
}

type MessageProcessor interface {
	ProcessEvent(hub *Hub, event *EventWithSender) error
}

type MessageProvider interface {
	CreateAndSaveMessage(message *model.Message) (*model.Message, error)
	GetHistory(matchID uint, limit int) ([]*model.Message, error)
	MarkMessageIsRead(messagesID []uint) error
	Delete(messagesID []uint) error
}

type ChatProvider interface {
	HandleNewConnection(uint, uint, *websocket.Conn)
	GetMessageHistory(uint, uint, int) ([]*model.Message, error)
}

type Notify interface {
	NotifyUser(userID uint, eventType string, payload interface{})
}

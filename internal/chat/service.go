package chat

import (
	"dating_service/internal/model"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

var ErrForbidden = errors.New("user is not a participant in this match")

type ChatService struct {
	hubs            map[uint]*Hub
	mu              sync.RWMutex
	matchProvider   MatchProvider
	messageProvider MessageProvider
	notify          Notify
}

func NewService(matchProvider MatchProvider, messageProvider MessageProvider, notify Notify) *ChatService {
	return &ChatService{
		hubs:            make(map[uint]*Hub),
		matchProvider:   matchProvider,
		messageProvider: messageProvider,
		notify:          notify,
	}
}

func (s *ChatService) HandleNewConnection(userID, matchID uint, conn *websocket.Conn) {
	isParticipant, err := s.matchProvider.IsUserInMatch(userID, matchID)
	if err != nil {
		log.Printf("Service Error: failed to check match participation: %v", err)
		conn.Close()
		return
	}
	if !isParticipant {
		log.Printf("Service Forbidden: user %d tried to access chat for match %d", userID, matchID)
		conn.Close()
		return
	}

	s.mu.Lock()
	hub, ok := s.hubs[matchID]
	if !ok {
		hub = NewHub(matchID, s)
		s.hubs[matchID] = hub
		go hub.Run()
	}
	s.mu.Unlock()

	client := &Client{
		ID:   userID,
		Hub:  hub,
		Conn: conn,
		Send: make(chan []byte, 256),
	}
	client.Hub.register <- client

	go client.writePump()
	go client.readPump()

	log.Printf("Client %d successfully connected to hub %d", userID, matchID)
}

func (s *ChatService) GetMessageHistory(userID, matchID uint, limit int) ([]*model.Message, error) {
	isParticipant, err := s.matchProvider.IsUserInMatch(userID, matchID)
	if err != nil {
		return nil, err
	}
	if !isParticipant {
		return nil, ErrForbidden
	}
	return s.messageProvider.GetHistory(matchID, limit)
}

func (s *ChatService) ProcessEvent(hub *Hub, eventData *EventWithSender) error {
	event := eventData.Event
	sender := eventData.Sender
	switch event.EventType {
	case "new_message":
		err := s.processMessage(event, sender, hub)
		if err != nil {
			return err
		}
	default:
		log.Printf("Unknown event type received: %s", event.EventType)
		return errors.New("unknown event type")
	}
	return nil
}

func (s *ChatService) processMessage(event *EventMessage, sender *Client, hub *Hub) error {
	var payload struct {
		MessageText string `json:"message_text"`
	}
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		return err
	}
	msg := &model.Message{
		MessageText: payload.MessageText,
		SenderID:    sender.ID,
		MatchID:     hub.ID,
		IsRead:      false,
	}
	message, err := s.messageProvider.CreateAndSaveMessage(msg)
	if err != nil {
		return err
	}
	msgOut := &MessageDto{
		ID:          message.ID,
		UpdatedAt:   message.UpdatedAt,
		MessageText: message.MessageText,
		IsRead:      message.IsRead,
		MatchID:     message.MatchID,
		SenderID:    sender.ID,
	}
	jsonMsg, err := json.Marshal(msgOut)
	if err != nil {
		return err
	}
	hub.Broadcast(jsonMsg)
	newEvent := EventMessage{
		EventType: event.EventType,
		Payload:   jsonMsg,
	}
	s.notifyNewMessage(&newEvent, sender, hub)
	return nil
}

func (s *ChatService) notifyNewMessage(event *EventMessage, sender *Client, hub *Hub) {
	secondUserOnline := false
	for client, _ := range hub.clients {
		if client.ID == sender.ID {
			continue
		} else {
			secondUserOnline = true
		}
	}
	if !secondUserOnline {
		users, err := s.matchProvider.GetUsers(hub.ID)
		if err != nil {
			log.Printf("Service Error: failed to get match users: %v", err)
			return
		}
		var recipientId uint
		if len(users) > 0 {
			if users[0].ID == sender.ID {
				recipientId = users[1].ID
			} else {
				recipientId = users[0].ID
			}
		} else {
			log.Printf("Service Error: failed to get match users: %v", ErrForbidden)
			return
		}
		s.notify.NotifyUser(recipientId, event.EventType, event.Payload)
	}
}

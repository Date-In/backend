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
}

func NewService(matchProvider MatchProvider, messageProvider MessageProvider) *ChatService {
	return &ChatService{
		hubs:            make(map[uint]*Hub),
		matchProvider:   matchProvider,
		messageProvider: messageProvider,
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
		return s.processMessage(event, sender, hub)
	default:
		log.Printf("Unknown event type received: %s", event.EventType)
		return errors.New("unknown event type")
	}
}

func (s *ChatService) processMessage(event *EventMessage, sender *Client, hub *Hub) error {
	var payload struct {
		MessageText string `json:"messageText"`
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
	if err := s.messageProvider.CreateAndSaveMessage(msg); err != nil {
		return err
	}
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	hub.Broadcast(jsonMsg)
	return nil
}

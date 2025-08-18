package chat

import (
	"dating_service/internal/model"
	"encoding/json"
	"log"
)

type Hub struct {
	ID         uint
	clients    map[*Client]bool
	broadcast  chan *model.Message
	register   chan *Client
	unregister chan *Client
	repo       *ChatRepository
}

func NewHub(id uint, repo *ChatRepository) *Hub {
	return &Hub{
		ID:         id,
		clients:    make(map[*Client]bool),
		broadcast:  make(chan *model.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		repo:       repo,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
		case message := <-h.broadcast:
			if err := h.repo.SaveMessage(message); err != nil {
				log.Printf("failed to save message: %v", err)
			}
			jsonMsg, _ := json.Marshal(message)
			for client := range h.clients {
				select {
				case client.Send <- jsonMsg:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
		}
	}
}

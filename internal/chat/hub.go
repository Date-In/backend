package chat

import (
	"log"
)

type Hub struct {
	ID           uint
	clients      map[*Client]bool
	processEvent chan *EventWithSender
	register     chan *Client
	unregister   chan *Client

	processor MessageProcessor
}

func NewHub(id uint, processor MessageProcessor) *Hub {
	return &Hub{
		ID:           id,
		clients:      make(map[*Client]bool),
		processEvent: make(chan *EventWithSender),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		processor:    processor,
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
		case eventWithSender := <-h.processEvent:
			if err := h.processor.ProcessEvent(h, eventWithSender); err != nil {
				log.Printf("hub failed to process event for match %d: %v", h.ID, err)
			}
		}
	}
}

func (h *Hub) Broadcast(message []byte) {
	for client := range h.clients {
		select {
		case client.Send <- message:
		default:
			close(client.Send)
			delete(h.clients, client)
		}
	}
}

func (h *Hub) BroadcastExcept(message []byte, sender *Client) {
	for client := range h.clients {
		if client != sender {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(h.clients, client)
			}
		}
	}
}

package notifier

import (
	"sync"
)

type Hub struct {
	clients    map[uint]*Client
	mu         sync.RWMutex
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[uint]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.ID] = client
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.ID]; ok {
				close(client.Send)
				delete(h.clients, client.ID)
			}
			h.mu.Unlock()
		}
	}
}

func (h *Hub) RegisterChannel() chan<- *Client {
	return h.register
}

func (h *Hub) UnregisterChannel() chan<- *Client {
	return h.unregister
}

func (h *Hub) SendTo(userID uint, message []byte) {
	h.mu.RLock()
	client, found := h.clients[userID]
	h.mu.RUnlock()

	if found {
		client.SafeSend(message)
	}
}

func (h *Hub) SendToMultiple(userIDs []uint, message []byte) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, id := range userIDs {
		if client, found := h.clients[id]; found {
			client.SafeSend(message)
		}
	}
}

func (h *Hub) Broadcast(message []byte) {
	h.mu.RLock()
	clientsSnapshot := make([]*Client, 0, len(h.clients))
	for _, client := range h.clients {
		clientsSnapshot = append(clientsSnapshot, client)
	}
	h.mu.RUnlock()

	for _, client := range clientsSnapshot {
		client.SafeSend(message)
	}
}

// IsOnline проверяет, находится ли пользователь онлайн.
func (h *Hub) IsOnline(userID uint) bool {
	h.mu.RLock()
	_, found := h.clients[userID]
	h.mu.RUnlock()
	return found
}

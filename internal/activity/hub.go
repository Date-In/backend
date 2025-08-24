package activity

import (
	"encoding/json"
	"log"
	"time"
)

type Hub struct {
	clients       map[uint]*Client
	register      chan *Client
	unregister    chan *Client
	storage       ActivityStorage
	matchProvider MatchProvider
}

func NewHub(storage ActivityStorage, matchProvider MatchProvider) *Hub {
	return &Hub{
		clients:       make(map[uint]*Client),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		storage:       storage,
		matchProvider: matchProvider,
	}
}

type StatusUpdateMessage struct {
	EventType string      `json:"event_type"`
	Payload   interface{} `json:"payload"`
}

type UserStatus struct {
	UserID   uint       `json:"user_id"`
	IsOnline bool       `json:"is_online"`
	LastSeen *time.Time `json:"last_seen,omitempty"`
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			func() {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("Activity Hub: recovered from panic in handleRegister: %v", r)
					}
				}()
				h.handleRegister(client)
			}()
		case client := <-h.unregister:
			func() {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("Activity Hub: recovered from panic in handleUnregister: %v", r)
					}
				}()
				h.handleUnregister(client)
			}()
		}
	}
}

func (h *Hub) handleRegister(client *Client) {
	log.Printf("Activity Hub: Handling registration for client %d...", client.ID)
	h.clients[client.ID] = client
	log.Printf("User %d connected to activity hub", client.ID)
	h.notifyContacts(client.ID, true)
	h.sendInitialStatuses(client)
}

func (h *Hub) handleUnregister(client *Client) {
	if _, ok := h.clients[client.ID]; ok {
		close(client.Send)
		delete(h.clients, client.ID)
		now := time.Now().UTC()
		if err := h.storage.UpdateLastSeen(client.ID, now); err != nil {
			log.Printf("Error updating last seen for user %d: %v", client.ID, err)
		}
		log.Printf("User %d disconnected from activity hub", client.ID)
		h.notifyContacts(client.ID, false)
	}
}

func (h *Hub) notifyContacts(userID uint, isOnline bool) {
	matchIDs, err := h.matchProvider.GetMatchUserIDs(userID)
	if err != nil {
		log.Printf("Could not find matches for user %d: %v", userID, err)
		return
	}

	var message StatusUpdateMessage
	if isOnline {
		message = StatusUpdateMessage{
			EventType: "user_online",
			Payload:   UserStatus{UserID: userID, IsOnline: true},
		}
	} else {
		lastSeen := time.Now().UTC()
		message = StatusUpdateMessage{
			EventType: "user_offline",
			Payload:   UserStatus{UserID: userID, IsOnline: false, LastSeen: &lastSeen},
		}
	}
	jsonMsg, _ := json.Marshal(message)
	for _, id := range matchIDs {
		if contactClient, found := h.clients[id]; found {
			contactClient.Send <- jsonMsg
		}
	}
}

func (h *Hub) sendInitialStatuses(client *Client) {
	log.Printf("[HUB-DEBUG] Step 1: Running sendInitialStatuses for UserID: %d", client.ID)
	matchIDs, err := h.matchProvider.GetMatchUserIDs(client.ID)
	log.Printf("[HUB-DEBUG] Step 2: For UserID %d, findMatches returned -> MatchIDs: %v, Error: %v", client.ID, matchIDs, err)
	if err != nil {
		log.Printf("[HUB-ERROR] Could not find matches for user %d: %v", client.ID, err)
		log.Printf("Could not find matches for user %d: %v", client.ID, err)
		return
	}
	lastSeenMap, err := h.storage.GetLastSeenForUsers(matchIDs)
	if err != nil {
		log.Printf("Could not get last seen for user %d contacts: %v", client.ID, err)
		return
	}

	statuses := make([]UserStatus, 0, len(matchIDs))
	for _, id := range matchIDs {
		if _, online := h.clients[id]; online {
			statuses = append(statuses, UserStatus{UserID: id, IsOnline: true})
		} else {
			lastSeenTime := lastSeenMap[id]
			statuses = append(statuses, UserStatus{UserID: id, IsOnline: false, LastSeen: &lastSeenTime})
		}
	}

	message := StatusUpdateMessage{
		EventType: "initial_statuses",
		Payload:   statuses,
	}

	jsonMsg, _ := json.Marshal(message)
	client.Send <- jsonMsg
}

package notifier

import (
	"encoding/json"
	"log"
	"time"
)

type Service struct {
	hub              *Hub
	activityProvider ActivityProvider
	matchProvider    MatchProvider
}

func NewService(hub *Hub, activityProvider ActivityProvider, matchProvider MatchProvider) *Service {
	return &Service{
		hub:              hub,
		activityProvider: activityProvider,
		matchProvider:    matchProvider,
	}
}

func (s *Service) HandleUserConnect(client *Client) {
	s.hub.RegisterChannel() <- client
	s.broadcastStatus(client.ID, true)
	s.sendInitialStatuses(client.ID)
}

func (s *Service) HandleUserDisconnect(client *Client) {
	if err := s.activityProvider.UpdateLastSeen(client.ID, time.Now().UTC()); err != nil {
		log.Printf("Error updating last seen for user %d: %v", client.ID, err)
	}
	s.broadcastStatus(client.ID, false)
	s.hub.UnregisterChannel() <- client
}

func (s *Service) NotifyUser(userID uint, eventType string, payload interface{}) {
	message := s.buildMessage(eventType, payload)
	if message != nil {
		s.hub.SendTo(userID, message)
	}
}

func (s *Service) broadcastStatus(userID uint, isOnline bool) {
	matchIDs, err := s.matchProvider.GetMatchUserIDs(userID)
	if err != nil {
		return
	}

	var payload UserStatus
	var eventType string

	if isOnline {
		eventType = "user_online"
		payload = UserStatus{UserID: userID, IsOnline: true}
	} else {
		eventType = "user_offline"
		lastSeen := time.Now().UTC()
		payload = UserStatus{UserID: userID, IsOnline: false, LastSeen: &lastSeen}
	}

	message := s.buildMessage(eventType, payload)
	if message != nil {
		s.hub.SendToMultiple(matchIDs, message)
	}
}

func (s *Service) sendInitialStatuses(userID uint) {
	matchIDs, err := s.matchProvider.GetMatchUserIDs(userID)
	if err != nil {
		return
	}

	if len(matchIDs) == 0 {
		return
	}

	lastSeenMap, err := s.activityProvider.GetLastSeenForUsers(matchIDs)
	if err != nil {
		return
	}

	statuses := make([]UserStatus, 0, len(matchIDs))
	for _, id := range matchIDs {
		var status UserStatus
		if s.hub.IsOnline(id) {
			status = UserStatus{UserID: id, IsOnline: true}
		} else {
			lastSeenTime := lastSeenMap[id]
			status = UserStatus{UserID: id, IsOnline: false, LastSeen: &lastSeenTime}
		}
		statuses = append(statuses, status)
	}

	message := s.buildMessage("initial_statuses", statuses)
	if message != nil {
		s.hub.SendTo(userID, message)
	}
}

func (s *Service) buildMessage(eventType string, payload interface{}) []byte {
	msg := Notify{
		EventType: eventType,
		Payload:   payload,
	}
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return nil
	}
	return jsonMsg
}

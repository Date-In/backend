package notifier

import "time"

type Notify struct {
	EventType string      `json:"event_type"`
	Payload   interface{} `json:"payload"`
}

type UserStatus struct {
	UserID   uint       `json:"user_id"`
	IsOnline bool       `json:"is_online"`
	LastSeen *time.Time `json:"last_seen,omitempty"`
}

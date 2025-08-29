package chat

import "encoding/json"

type EventMessage struct {
	EventType string          `json:"event_type"`
	Payload   json.RawMessage `json:"payload"`
}

type EventWithSender struct {
	Event  *EventMessage
	Sender *Client
}

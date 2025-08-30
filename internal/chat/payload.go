package chat

import "time"

type MessageDto struct {
	ID          uint      `json:"id"`
	UpdatedAt   time.Time `json:"updated_at"`
	MessageText string    `json:"message_text"`
	IsRead      bool      `json:"is_read"`
	MatchID     uint      `json:"match_id"`
	SenderID    uint      `json:"sender_id"`
}

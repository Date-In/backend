package match

import "time"

type MatchPreviewDTO struct {
	MatchID     uint            `json:"match_id"`
	Partner     PartnerDTO      `json:"partner"`
	LastMessage *LastMessageDTO `json:"last_message"`
}

type PartnerDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type LastMessageDTO struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	IsRead    bool      `json:"is_read"`
}

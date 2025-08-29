package message

import "dating_service/internal/model"

type MessageStorage interface {
	Save(message *model.Message) error
	GetHistory(matchID uint, limit int) ([]*model.Message, error)
}

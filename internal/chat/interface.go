package chat

import "dating_service/internal/model"

type ChatStorage interface {
	SaveMessage(*model.Message) error
	GetMessageHistory(uint, int) ([]*model.Message, error)
}

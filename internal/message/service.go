package message

import (
	"dating_service/internal/model"
)

type MessageService struct {
	messageStorage MessageStorage
}

func NewMessageService(messageStorage MessageStorage) *MessageService {
	return &MessageService{messageStorage: messageStorage}
}

func (s *MessageService) CreateAndSaveMessage(msg *model.Message) (*model.Message, error) {
	message, err := s.messageStorage.Save(msg)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *MessageService) GetHistory(matchID uint, limit int) ([]*model.Message, error) {
	return s.messageStorage.GetHistory(matchID, limit)
}

func (s *MessageService) MarkMessageIsRead(messagesID []uint) error {
	return s.messageStorage.MarkMessageIsRead(messagesID)
}

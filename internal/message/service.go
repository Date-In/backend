package message

import (
	"dating_service/internal/model"
)

type Service struct {
	messageStorage MessageStorage
}

func NewMessageService(messageStorage MessageStorage) *Service {
	return &Service{messageStorage: messageStorage}
}

func (s *Service) CreateAndSaveMessage(msg *model.Message) (*model.Message, error) {
	message, err := s.messageStorage.Save(msg)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *Service) GetHistory(matchID uint, limit int) ([]*model.Message, error) {
	return s.messageStorage.GetHistory(matchID, limit)
}

func (s *Service) MarkMessageIsRead(messagesID []uint) error {
	return s.messageStorage.MarkMessageIsRead(messagesID)
}

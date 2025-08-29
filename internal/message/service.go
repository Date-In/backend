package message

import "dating_service/internal/model"

type MessageService struct {
	messageStorage MessageStorage
}

func NewMessageService(messageStorage MessageStorage) *MessageService {
	return &MessageService{messageStorage: messageStorage}
}

func (s *MessageService) CreateAndSaveMessage(msg *model.Message) error {
	if err := s.messageStorage.Save(msg); err != nil {
		return err
	}
	return nil
}

func (s *MessageService) GetHistory(matchID uint, limit int) ([]*model.Message, error) {
	return s.messageStorage.GetHistory(matchID, limit)
}

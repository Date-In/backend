package message

import (
	"dating_service/internal/model"
)

type Service struct {
	messageStorage MessageStorage
	cryptoProvider CryptoProvider
}

func NewService(messageStorage MessageStorage, cryptoProvider CryptoProvider) *Service {
	return &Service{messageStorage: messageStorage, cryptoProvider: cryptoProvider}
}

func (s *Service) CreateAndSaveMessage(msg *model.Message) (*model.Message, error) {
	cryptMessageText, err := s.cryptoProvider.EncryptString(msg.MessageText)
	if err != nil {
		return nil, err
	}
	msg.MessageText = cryptMessageText
	message, err := s.messageStorage.Save(msg)
	if err != nil {
		return nil, err
	}
	message.MessageText, err = s.cryptoProvider.DecryptString(message.MessageText)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *Service) GetHistory(matchID uint, limit int) ([]*model.Message, error) {
	messages, err := s.messageStorage.GetHistory(matchID, limit)
	if err != nil {
		return nil, err
	}
	for _, msg := range messages {
		decryptTxt, err := s.cryptoProvider.DecryptString(msg.MessageText)
		if err != nil {
			return nil, err
		}
		msg.MessageText = decryptTxt
	}
	return messages, nil
}

func (s *Service) MarkMessageIsRead(messagesID []uint) error {
	return s.messageStorage.MarkMessageIsRead(messagesID)
}

func (s *Service) Delete(messagesID []uint) error {
	return s.messageStorage.Delete(messagesID)
}

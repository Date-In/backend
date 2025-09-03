package message

import "dating_service/internal/model"

type MessageStorage interface {
	Save(message *model.Message) (*model.Message, error)
	GetHistory(matchID uint, limit int) ([]*model.Message, error)
	MarkMessageIsRead(messagesID []uint) error
	Delete(messagesID []uint) error
}

type CryptoProvider interface {
	EncryptString(plain string) (string, error)
	DecryptString(encrypted string) (string, error)
}

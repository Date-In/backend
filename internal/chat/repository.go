package chat

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
)

type ChatRepository struct {
	db *db.Db
}

func NewChatRepository(db *db.Db) *ChatRepository {
	return &ChatRepository{db}
}

func (r *ChatRepository) SaveMessage(message *model.Message) error {
	result := r.db.PgDb.Create(message)
	return result.Error
}

func (r *ChatRepository) GetMessageHistory(matchID uint, limit int) ([]*model.Message, error) {
	var messages []*model.Message

	err := r.db.PgDb.Where("match_id = ?", matchID).
		Order("created_at desc").
		Limit(limit).
		Find(&messages).Error

	if err != nil {
		return nil, err
	}
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	return messages, nil
}

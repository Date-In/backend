package message

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
)

type MessageRepository struct {
	db *db.Db
}

func NewMessageRepository(db *db.Db) *MessageRepository {
	return &MessageRepository{db}
}

func (r *MessageRepository) Save(message *model.Message) error {
	result := r.db.PgDb.Create(message)
	return result.Error
}

func (r *MessageRepository) GetHistory(matchID uint, limit int) ([]*model.Message, error) {
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

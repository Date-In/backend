package message

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *db.Db
}

func NewMessageRepository(db *db.Db) *Repository {
	return &Repository{db}
}

func (r *Repository) Save(message *model.Message) (*model.Message, error) {
	result := r.db.PgDb.Create(message)
	if result.Error != nil {
		return nil, result.Error
	}
	return message, nil
}

func (r *Repository) GetHistory(matchID uint, limit int) ([]*model.Message, error) {
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

func (r *Repository) MarkMessageIsRead(messagesID []uint) error {
	err := r.db.PgDb.Model(&model.Message{}).
		Where("id IN ?", messagesID).
		Update("is_read", true).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return nil
}

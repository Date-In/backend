package action

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"time"
)

type Repository struct {
	db *db.Db
}

func NewRepository(db *db.Db) *Repository {
	return &Repository{db}
}

func (r *Repository) GetNonActiveUserIds(olderThan time.Time) ([]uint, error) {
	var userIds []uint
	err := r.db.PgDb.Model(&model.Actions{}).
		Where("action < ? ", olderThan).
		Pluck("user_id", &userIds).Error
	if err != nil {
		return nil, err
	}
	return userIds, nil
}

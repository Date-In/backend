package action

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"time"
)

type ActionsRepository struct {
	db *db.Db
}

func NewActionsRepository(db *db.Db) *ActionsRepository {
	return &ActionsRepository{db}
}

func (repo *ActionsRepository) GetNonActiveUserIds(olderThan time.Time) ([]uint, error) {
	var userIds []uint
	err := repo.db.PgDb.Model(&model.Actions{}).
		Where("action < ? ", olderThan).
		Pluck("user_id", &userIds).Error
	if err != nil {
		return nil, err
	}
	return userIds, nil
}

package action

import (
	"dating_service/pkg/db"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ActionsRepository struct {
	db *db.Db
}

func NewActionsRepository(db *db.Db) *ActionsRepository {
	return &ActionsRepository{db}
}

func (repo *ActionsRepository) Get(userID uint) (*Actions, error) {
	var actions *Actions
	err := repo.db.PgDb.Where("user_id = ?", userID).First(&actions).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return actions, nil
}

func (repo *ActionsRepository) Update(userID uint, actionTime time.Time) error {
	action := Actions{
		UserID: userID,
		Action: actionTime,
	}
	err := repo.db.PgDb.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"action"}),
	}).Create(&action).Error

	return err
}

func (repo *ActionsRepository) GetNonActiveUserIds(olderThan time.Time) ([]uint, error) {
	var userIds []uint
	err := repo.db.PgDb.Model(&Actions{}).
		Where("action < ? ", olderThan).
		Pluck("user_id", &userIds).Error
	if err != nil {
		return nil, err
	}
	return userIds, nil
}

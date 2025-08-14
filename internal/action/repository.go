package action

import (
	"dating_service/pkg/db"
	"errors"
	"gorm.io/gorm"
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
	err := repo.db.PgDb.Where("user_id = ?", userID).First(&Actions{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			repo.db.PgDb.Create(&Actions{userID, actionTime})
		} else {
			return err
		}
	} else {
		err = repo.db.PgDb.Updates(&Actions{userID, actionTime}).Error
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

package filter

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"errors"
	"gorm.io/gorm"
)

type FilterRepository struct {
	db *db.Db
}

func NewFilterRepository(db *db.Db) *FilterRepository {
	return &FilterRepository{db}
}

func (repo *FilterRepository) GetFilterUser(userId uint) (*model.FilterSearch, error) {
	var filter model.FilterSearch

	err := repo.db.PgDb.Where("user_id = ?", userId).First(&filter).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &filter, err
}

func (repo *FilterRepository) CreateFilter(filter model.FilterSearch) error {
	return repo.db.PgDb.Create(&filter).Error
}

func (repo *FilterRepository) UpdateFilter(filter model.FilterSearch) error {
	return repo.db.PgDb.Updates(&filter).Error
}

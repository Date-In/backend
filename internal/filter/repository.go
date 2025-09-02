package filter

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *db.Db
}

func NewFilterRepository(db *db.Db) *Repository {
	return &Repository{db}
}

func (r *Repository) GetFilterUser(userId uint) (*model.FilterSearch, error) {
	var filter model.FilterSearch

	err := r.db.PgDb.Where("user_id = ?", userId).First(&filter).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &filter, err
}

func (r *Repository) CreateFilter(filter model.FilterSearch) error {
	return r.db.PgDb.Create(&filter).Error
}

func (r *Repository) UpdateFilter(filter model.FilterSearch) error {
	return r.db.PgDb.Updates(&filter).Error
}

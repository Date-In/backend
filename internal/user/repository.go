package user

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Create(user *model.User) error {

	return repo.db.PgDb.Create(user).Error
}

func (repo *UserRepository) FindByPhone(phone string) (*model.User, error) {
	var user model.User
	err := repo.db.PgDb.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

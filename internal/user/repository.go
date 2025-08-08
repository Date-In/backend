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

func (repo *UserRepository) FindById(id uint) (*model.User, error) {
	var user model.User
	err := repo.db.PgDb.
		Preload("Sex").
		Preload("ZodiacSign").
		Preload("Worldview").
		Preload("TypeOfDating").
		Preload("Education").
		Preload("AttitudeToAlcohol").
		Preload("AttitudeToSmoking").
		Preload("Status").
		Preload("Interests").
		Preload("Photos").
		Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return &user, err
}

func (repo *UserRepository) Update(id uint, updateData *model.User) error {
	result := repo.db.PgDb.
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(updateData)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

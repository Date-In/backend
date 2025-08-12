package photo

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"errors"
	"gorm.io/gorm"
)

type PhotoRepository struct {
	db *db.Db
}

func NewPhotoRepository(db *db.Db) *PhotoRepository {
	return &PhotoRepository{db}
}

func (repo *PhotoRepository) Save(photo *model.Photo) error {
	return repo.db.PgDb.Save(photo).Error
}

func (repo *PhotoRepository) GetById(uuid string) (*model.Photo, error) {
	var photo *model.Photo
	err := repo.db.PgDb.Where("id = ?", uuid).First(&photo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return photo, err
}

func (repo *PhotoRepository) CountPhoto(id uint) (int, error) {
	var count int64
	err := repo.db.PgDb.Model(&model.Photo{}).Where("user_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (repo *PhotoRepository) DeleteById(uuid string, userId uint) (int64, error) {
	result := repo.db.PgDb.Where("id = ? and user_id = ?", uuid, userId).Delete(&model.Photo{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *PhotoRepository) FindAllIDs(userId uint) ([]string, error) {
	var ids []string

	err := repo.db.PgDb.Model(&model.Photo{}).
		Where("user_id = ?", userId).
		Order("uploaded_at DESC").
		Pluck("id", &ids).Error

	if err != nil {
		return nil, err
	}
	return ids, nil
}

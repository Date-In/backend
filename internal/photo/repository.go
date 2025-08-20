package photo

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"errors"
	"fmt"
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

func (repo *PhotoRepository) FindAvatar(userId uint) (string, error) {
	var avatarID string

	err := repo.db.PgDb.Model(&model.Photo{}).
		Select("id").
		Where("user_id = ? AND is_avatar = ?", userId, true).
		First(&avatarID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}
		return "", fmt.Errorf("failed to find avatar: %v", err)
	}

	return avatarID, nil
}

func (repo *PhotoRepository) FindUserPhotoWithoutAvatar(userId uint) ([]string, error) {
	var photoIds []string

	err := repo.db.PgDb.Model(&model.Photo{}).
		Select("id").
		Where("user_id = ? and is_avatar = ?", userId, false).
		Find(&photoIds).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return photoIds, nil
		}
		return nil, err
	}
	return photoIds, nil
}

func (repo *PhotoRepository) ChangeAvatarUser(userId uint, photoId string) (string, error) {
	tx := repo.db.PgDb.Begin()
	var newAvatarID string

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var photo model.Photo
	if err := tx.Where("id = ? AND user_id = ?", photoId, userId).First(&photo).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrPhotoNotFound
		}
		return "", fmt.Errorf("failed to check photo: %v", err)
	}

	if err := tx.Model(&model.Photo{}).
		Where("user_id = ?", userId).
		Update("is_avatar", false).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to reset avatars: %v", err)
	}

	if err := tx.Model(&model.Photo{}).
		Where("id = ?", photoId).
		Update("is_avatar", true).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to set new avatar: %v", err)
	}

	if err := tx.Model(&model.Photo{}).
		Select("id").
		Where("user_id = ? AND is_avatar = ?", userId, true).
		First(&newAvatarID).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to verify new avatar: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return newAvatarID, nil
}

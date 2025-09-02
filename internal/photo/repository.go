package photo

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	db *db.Db
}

func NewPhotoRepository(db *db.Db) *Repository {
	return &Repository{db}
}

func (r *Repository) Save(photo *model.Photo) error {
	return r.db.PgDb.Save(photo).Error
}

func (r *Repository) GetById(uuid string) (*model.Photo, error) {
	var photo *model.Photo
	err := r.db.PgDb.Where("id = ?", uuid).First(&photo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return photo, err
}

func (r *Repository) CountPhoto(id uint) (int, error) {
	var count int64
	err := r.db.PgDb.Model(&model.Photo{}).Where("user_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *Repository) DeleteById(uuid string, userId uint) (int64, error) {
	result := r.db.PgDb.Where("id = ? and user_id = ?", uuid, userId).Delete(&model.Photo{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (r *Repository) FindAllIDs(userId uint) ([]string, error) {
	var ids []string

	err := r.db.PgDb.Model(&model.Photo{}).
		Where("user_id = ?", userId).
		Order("uploaded_at DESC").
		Pluck("id", &ids).Error

	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (r *Repository) FindAvatar(userId uint) (*model.Photo, error) {
	var photo *model.Photo

	err := r.db.PgDb.Model(&model.Photo{}).
		Where("user_id = ? AND is_avatar = ?", userId, true).
		First(&photo).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find avatar: %v", err)
	}

	return photo, nil
}

func (r *Repository) FindUserPhotoWithoutAvatar(userId uint) ([]string, error) {
	var photoIds []string

	err := r.db.PgDb.Model(&model.Photo{}).
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

func (r *Repository) ChangeAvatarUser(userId uint, photoId string) (string, error) {
	tx := r.db.PgDb.Begin()
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

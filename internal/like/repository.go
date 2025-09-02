package like

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *db.Db
}

func NewRepository(db *db.Db) *Repository {
	return &Repository{db}
}

func (r *Repository) GetLikes(userId uint) ([]model.Like, error) {
	var likes []model.Like

	err := r.db.PgDb.Model(&model.Like{}).
		Preload("User.Avatar", "is_avatar = ?", true).
		Preload("User.Photos", "is_avatar = ?", false).
		Where("target_id = ?", userId).
		Find(&likes).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return likes, nil
		}
		return nil, err
	}
	return likes, nil
}

func (r *Repository) CreateLike(userId uint, targetId uint) error {
	err := r.db.PgDb.Create(&model.Like{
		UserID:    userId,
		TargetID:  targetId,
		CreatedAt: time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteLike(userId uint, targetId uint) error {
	result := r.db.PgDb.Where("user_id = ? AND target_id = ?", userId, targetId).Delete(&model.Like{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFoundLike
	}
	return nil
}

func (r *Repository) FindLikeByTargetIdAndUserID(targetID, userID uint) (*model.Like, error) {
	var like model.Like
	err := r.db.PgDb.Model(&model.Like{}).
		Where("target_id = ? AND user_id = ?", targetID, userID).
		First(&like).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &like, nil
}

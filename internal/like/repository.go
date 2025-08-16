package like

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"gorm.io/gorm"
	"time"
)

type LikeRepository struct {
	db *db.Db
}

func NewLikeRepository(db *db.Db) *LikeRepository {
	return &LikeRepository{db}
}

func (repo *LikeRepository) GetLikes(userId uint) ([]model.Like, error) {
	var likes []model.Like

	err := repo.db.PgDb.Model(&model.Like{}).
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

func (repo *LikeRepository) CreateLike(userId uint, targetId uint) error {
	err := repo.db.PgDb.Create(&model.Like{
		UserID:    userId,
		TargetID:  targetId,
		CreatedAt: time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *LikeRepository) FindLikeByTargetIdAndUserID(targetID, userID uint) (*model.Like, error) {
	var like model.Like
	err := repo.db.PgDb.Model(&model.Like{}).
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

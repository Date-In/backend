package match

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"errors"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	db *db.Db
}

func NewRepository(db *db.Db) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(userID1, userID2 uint) error {
	err := r.db.PgDb.Create(&model.Match{
		User1ID:   userID1,
		User2ID:   userID2,
		CreatedAt: time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllWithDetails(userId uint) ([]model.Match, error) {
	var matches []model.Match

	err := r.db.PgDb.Model(&model.Match{}).
		Preload("User1", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("User1.Avatar", "is_avatar = ?", true).
		Preload("User2", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("User2.Avatar", "is_avatar = ?", true).
		Preload("LastMessage", func(db *gorm.DB) *gorm.DB {
			return db.Order("messages.created_at DESC")
		}).
		Where("user1_id = ? OR user2_id = ?", userId, userId).
		Find(&matches).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return make([]model.Match, 0), nil
		}
		return nil, err
	}
	return matches, nil
}

func (r *Repository) IsUserInMatch(userID uint, matchID uint) (bool, error) {
	var count int64
	result := r.db.PgDb.Model(&model.Match{}).
		Where("id = ? AND (user1_id = ? OR user2_id = ?)", matchID, userID, userID).
		Count(&count)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return count > 0, nil
}

func (r *Repository) GetMatchUserIDs(userID uint) ([]uint, error) {
	var partnerIDs []uint
	var partnerIDsFromUser1 []uint
	err := r.db.PgDb.Model(&model.Match{}).
		Where("user1_id = ?", userID).
		Pluck("user2_id", &partnerIDsFromUser1).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	var partnerIDsFromUser2 []uint
	err = r.db.PgDb.Model(&model.Match{}).
		Where("user2_id = ?", userID).
		Pluck("user1_id", &partnerIDsFromUser2).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	partnerIDs = append(partnerIDs, partnerIDsFromUser1...)
	partnerIDs = append(partnerIDs, partnerIDsFromUser2...)
	return partnerIDs, nil
}

func (r *Repository) GetUsers(matchID uint) ([]model.User, error) {
	var match model.Match

	err := r.db.PgDb.
		Preload("User1").
		Preload("User2").
		First(&match, matchID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("match not found")
		}
		return nil, err
	}

	users := []model.User{match.User1, match.User2}

	return users, nil
}

func (r *Repository) Delete(matchID uint) error {
	err := r.db.PgDb.Delete(&model.Match{}, matchID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("match not found")
		}
		return err
	}
	return nil
}

func (r *Repository) GetAll() ([]model.Match, error) {
	var matches []model.Match
	err := r.db.PgDb.
		Preload("LastMessage").
		Find(&matches).Error
	if err != nil {
		return nil, err
	}
	return matches, nil
}

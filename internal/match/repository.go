package match

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"errors"
	"gorm.io/gorm"
	"time"
)

type MatchRepository struct {
	db *db.Db
}

func NewMatchRepository(db *db.Db) *MatchRepository {
	return &MatchRepository{db}
}

func (repo *MatchRepository) Create(userID1, userID2 uint) error {
	var maxId uint
	var minId uint
	if userID1 > userID2 {
		maxId = userID1
		minId = userID2
	} else {
		maxId = userID2
		minId = userID1
	}
	err := repo.db.PgDb.Create(&model.Match{
		User1ID:   maxId,
		User2ID:   minId,
		CreatedAt: time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *MatchRepository) GetAll(userId uint) ([]model.Match, error) {
	var matches []model.Match
	err := repo.db.PgDb.Model(&model.Match{}).
		Where("user1_id = ? OR user2_id = ?", userId, userId).
		Find(&matches).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return matches, nil
		}
		return nil, err
	}
	return matches, nil
}

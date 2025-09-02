package activity

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"gorm.io/gorm/clause"
	"time"
)

type Repository struct {
	db *db.Db
}

func NewRepository(db *db.Db) *Repository {
	return &Repository{db}
}

func (r *Repository) UpdateLastSeen(userID uint, seenAt time.Time) error {
	actionRecord := model.Actions{
		UserID: userID,
		Action: seenAt,
	}
	return r.db.PgDb.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"action"}),
	}).Create(&actionRecord).Error
}

func (r *Repository) GetLastSeenForUsers(userIDs []uint) (map[uint]time.Time, error) {
	var results []model.Actions
	err := r.db.PgDb.Select("user_id", "action").Where("user_id IN ?", userIDs).Find(&results).Error
	if err != nil {
		return nil, err
	}
	lastSeenMap := make(map[uint]time.Time, len(results))
	for _, res := range results {
		lastSeenMap[res.UserID] = res.Action
	}

	return lastSeenMap, nil
}

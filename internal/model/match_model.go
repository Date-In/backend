package model

import (
	"time"
)

type Match struct {
	ID          uint `gorm:"primary_key"`
	User1ID     uint `gorm:"uniqueIndex:idx_match_users;not null"`
	User2ID     uint `gorm:"uniqueIndex:idx_match_users;not null"`
	CreatedAt   time.Time
	User1       User      `gorm:"foreignKey:User1ID"`
	User2       User      `gorm:"foreignKey:User2ID"`
	Messages    []Message `gorm:"foreignKey:MatchID"`
	LastMessage *Message  `gorm:"foreignKey:MatchID"`
}

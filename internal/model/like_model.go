package model

import (
	"time"
)

type Like struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UserID    uint `gorm:"not null"`
	TargetID  uint `gorm:"not null"`
	User      User `gorm:"foreignKey:UserID"`
	Target    User `gorm:"foreignKey:TargetID"`
}

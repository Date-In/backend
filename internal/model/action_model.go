package model

import "time"

type Actions struct {
	UserID uint `gorm:"primary_key"`
	Action time.Time
}

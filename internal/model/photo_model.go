package model

import (
	"time"
)

type Photo struct {
	ID         string `gorm:"primaryKey"`
	Url        string
	IsAvatar   bool `gorm:"default:false"`
	UserID     uint `gorm:"not null"`
	UploadedAt time.Time
}

func NewPhoto(id string, Url string, userID uint) *Photo {
	return &Photo{
		ID:         id,
		Url:        Url,
		IsAvatar:   false,
		UserID:     userID,
		UploadedAt: time.Now(),
	}
}

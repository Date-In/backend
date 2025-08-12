package model

import (
	"github.com/google/uuid"
	"time"
)

type Photo struct {
	ID         string `gorm:"primaryKey"`
	FileName   string
	FileType   string
	IsAvatar   bool   `gorm:"default:false"`
	Data       []byte `gorm:"type:bytea"`
	UserID     uint   `gorm:"not null"`
	UploadedAt time.Time
}

func NewPhoto(fileName, fileType string, data []byte, userID uint) *Photo {
	id := uuid.New().String()
	return &Photo{
		ID:         id,
		FileName:   fileName,
		FileType:   fileType,
		IsAvatar:   false,
		Data:       data,
		UserID:     userID,
		UploadedAt: time.Now(),
	}
}

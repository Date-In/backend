package model

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	FileName string
	FileType string
	IsAvatar bool   `gorm:"default:false"`
	Data     []byte `gorm:"type:bytea"`
}

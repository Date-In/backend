package model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model

	MessageText string `gorm:"type:text;not null"`
	IsRead      bool   `gorm:"default:false"`

	MatchID  uint `gorm:"not null"`
	SenderID uint `gorm:"not null"`

	Match  Match `gorm:"foreignKey:MatchID"`
	Sender User  `gorm:"foreignKey:SenderID"`
}

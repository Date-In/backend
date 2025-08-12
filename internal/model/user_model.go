package model

import (
	"gorm.io/gorm"
)

type Sex struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:50"`
}

type Education struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

type ZodiacSign struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:50"`
}

type Worldview struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

type TypeOfDating struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

type AttitudeToAlcohol struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

type AttitudeToSmoking struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

type Status struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

type User struct {
	gorm.Model
	Name                string  `gorm:"size:100;not null"`
	Phone               string  `gorm:"unique;size:20"`
	Password            string  `gorm:"not null"`
	Age                 uint    `gorm:"not null"`
	Bio                 *string `gorm:"type:text"`
	City                *string
	Children            *bool
	Height              *uint
	SexID               uint
	ZodiacSignID        *uint
	WorldviewID         *uint
	TypeOfDatingID      *uint
	EducationID         *uint
	AttitudeToAlcoholID *uint
	AttitudeToSmokingID *uint
	StatusID            *uint
	Sex                 Sex               `gorm:"foreignKey:SexID"`
	ZodiacSign          ZodiacSign        `gorm:"foreignKey:ZodiacSignID"`
	Worldview           Worldview         `gorm:"foreignKey:WorldviewID"`
	TypeOfDating        TypeOfDating      `gorm:"foreignKey:TypeOfDatingID"`
	Education           Education         `gorm:"foreignKey:EducationID"`
	AttitudeToAlcohol   AttitudeToAlcohol `gorm:"foreignKey:AttitudeToAlcoholID"`
	AttitudeToSmoking   AttitudeToSmoking `gorm:"foreignKey:AttitudeToSmokingID"`
	Status              Status            `gorm:"foreignKey:StatusID"`
	Interests           []*Interest       `gorm:"many2many:user_interests;"`
	LikesSent           []Like            `gorm:"foreignKey:UserID"`
	LikesReceived       []Like            `gorm:"foreignKey:TargetID"`
	Photos              []*Photo          `gorm:"foreignKey:UserID"`
}

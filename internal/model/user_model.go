package model

import (
	"gorm.io/gorm"
)

type Sex struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:50"`
}

func (i *Sex) GetID() uint {
	return i.ID
}

func (i *Sex) GetName() string {
	return i.Name
}

type Education struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

func (i *Education) GetID() uint {
	return i.ID
}

func (i *Education) GetName() string {
	return i.Name
}

type ZodiacSign struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:50"`
}

func (i *ZodiacSign) GetID() uint {
	return i.ID
}

func (i *ZodiacSign) GetName() string {
	return i.Name
}

type Worldview struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

func (i *Worldview) GetID() uint {
	return i.ID
}

func (i *Worldview) GetName() string {
	return i.Name
}

type TypeOfDating struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

func (i *TypeOfDating) GetID() uint {
	return i.ID
}

func (i *TypeOfDating) GetName() string {
	return i.Name
}

type AttitudeToAlcohol struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

func (i *AttitudeToAlcohol) GetID() uint {
	return i.ID
}

func (i *AttitudeToAlcohol) GetName() string {
	return i.Name
}

type AttitudeToSmoking struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

func (i *AttitudeToSmoking) GetID() uint {
	return i.ID
}

func (i *AttitudeToSmoking) GetName() string {
	return i.Name
}

type Status struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

func (i *Status) GetID() uint {
	return i.ID
}

func (i *Status) GetName() string {
	return i.Name
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
	StatusID            uint
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
	Avatar              *Photo            `gorm:"foreignKey:UserID"`
}

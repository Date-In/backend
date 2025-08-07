package model

type Interest struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

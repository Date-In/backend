package model

type FilterSearch struct {
	ID     uint
	UserID uint
	MinAge uint
	MaxAge uint
	Sex    string
	User   User `gorm:"foreignkey:UserID"`
}

package model

type FilterSearch struct {
	ID       uint `gorm:"primary_key"`
	UserID   uint
	MinAge   uint
	MaxAge   uint
	SexID    uint
	Location string
	User     User `gorm:"foreignkey:UserID"`
}

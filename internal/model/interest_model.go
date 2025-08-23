package model

type Interest struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null;size:100"`
}

func (i *Interest) GetID() uint {
	return i.ID
}

func (i *Interest) GetName() string {
	return i.Name
}

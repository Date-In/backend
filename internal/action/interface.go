package action

import "time"

type UserProvider interface {
	ReactivateUser(uint) error
	ChangeStatus([]uint) error
}

type ActionStorage interface {
	Get(uint) (*Actions, error)
	Update(uint, time.Time) error
	GetNonActiveUserIds(time.Time) ([]uint, error)
}

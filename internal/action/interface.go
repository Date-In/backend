package action

import (
	"time"
)

type UserProvider interface {
	ReactivateUser(uint) error
	ChangeStatus([]uint) error
}

type ActionStorage interface {
	GetNonActiveUserIds(time.Time) ([]uint, error)
}

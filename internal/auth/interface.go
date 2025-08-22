package auth

import "dating_service/internal/model"

type UserProvider interface {
	FindUserByPhone(string) (*model.User, error)
	Create(user *model.User) error
}

type CacheProvider interface {
	IsValidSexID(uint) bool
}

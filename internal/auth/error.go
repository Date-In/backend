package auth

import "errors"

var (
	ErrUserAlreadyExists        = errors.New("user already exist")
	ErrIncorrectPasswordOrPhone = errors.New("incorrect password or phone")
	ErrInvalidSexID             = errors.New("incorrect sex id")
)

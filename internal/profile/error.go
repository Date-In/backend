package profile

import "errors"

var (
	ErrUserNotFound                 = errors.New("User not found")
	ErrInvalidSexID                 = errors.New("invalid sex id")
	ErrInvalidZodiacID              = errors.New("invalid zodiac id")
	ErrInvalidWordViewID            = errors.New("invalid word view id")
	ErrInvalidTypeOfDatingId        = errors.New("invalid type of dating id")
	ErrInvalidEducationId           = errors.New("invalid education id")
	ErrInvalidAttitudeToAlcoholicId = errors.New("invalid attitude to alcoholic id")
	ErrInvalidAttitudeToSmokingId   = errors.New("invalid attitude to smoking id")
)

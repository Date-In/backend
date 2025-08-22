package profile

import (
	"context"
	"dating_service/internal/model"
)

type UserProvider interface {
	FindById(uint) (*model.User, error)
	Update(uint, *model.User) error
	UpdateInterests(uint, []*model.Interest) ([]*model.Interest, error)
}
type PhotoProvider interface {
	CountPhoto(uint) (int, error)
	AddPhoto(context.Context, uint, []byte, string) (*model.Photo, error)
	DeletePhoto(context.Context, string, uint) (int, error)
	ChangeAvatarUser(string, uint) (string, error)
	FindAvatar(uint) (string, error)
}

type CacheProvider interface {
	IsValidSexID(id uint) bool
	IsValidEducationID(id uint) bool
	IsValidZodiacSignID(id uint) bool
	IsValidWorldviewID(id uint) bool
	IsValidTypeOfDatingID(id uint) bool
	IsValidAttitudeToAlcoholID(id uint) bool
	IsValidAttitudeToSmokingID(id uint) bool
	IsValidStatusID(id uint) bool
	IsValidInterestIDs(ids []uint) bool
	IsValidInterest(id uint) bool
	GetSexByID(id uint) model.Sex
	GetEducationByID(id uint) model.Education
	GetZodiacSignByID(id uint) model.ZodiacSign
	GetWorldviewByID(id uint) model.Worldview
	GetTypeOfDatingByID(id uint) model.TypeOfDating
	GetAttitudeToAlcoholByID(id uint) model.AttitudeToAlcohol
	GetAttitudeToSmokingByID(id uint) model.AttitudeToSmoking
	GetStatusByID(id uint) model.Status
	GetInterestByID(id uint) model.Interest
}

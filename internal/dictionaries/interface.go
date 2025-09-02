package dictionaries

import "dating_service/internal/model"

type CacheProvider interface {
	GetSexes() ([]model.Sex, error)
	GetEducations() ([]model.Education, error)
	GetZodiacSigns() ([]model.ZodiacSign, error)
	GetWorldViews() ([]model.Worldview, error)
	GetTypeOfDating() ([]model.TypeOfDating, error)
	GetAttitudeToAlcohol() ([]model.AttitudeToAlcohol, error)
	GetAttitudeToSmoking() ([]model.AttitudeToSmoking, error)
	GetInterests() ([]model.Interest, error)
	GetStatuses() ([]model.Status, error)
}

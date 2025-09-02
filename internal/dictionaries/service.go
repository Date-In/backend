package dictionaries

import "dating_service/internal/model"

type Service struct {
	cache CacheProvider
}

func NewService(cache CacheProvider) *Service {
	return &Service{cache}
}

func (s *Service) GetSexes() ([]model.Sex, error) {
	return s.cache.GetSexes()
}

func (s *Service) GetEducations() ([]model.Education, error) {
	return s.cache.GetEducations()
}

func (s *Service) GetZodiacSigns() ([]model.ZodiacSign, error) {
	return s.cache.GetZodiacSigns()
}

func (s *Service) GetWorldViews() ([]model.Worldview, error) {
	return s.cache.GetWorldViews()
}
func (s *Service) GetTypeOfDating() ([]model.TypeOfDating, error) {
	return s.cache.GetTypeOfDating()
}

func (s *Service) GetAttitudeToAlcohol() ([]model.AttitudeToAlcohol, error) {
	return s.cache.GetAttitudeToAlcohol()
}

func (s *Service) GetAttitudeToSmoking() ([]model.AttitudeToSmoking, error) {
	return s.cache.GetAttitudeToSmoking()
}

func (s *Service) GetInterests() ([]model.Interest, error) {
	return s.cache.GetInterests()
}

func (s *Service) GetStatuses() ([]model.Status, error) {
	return s.cache.GetStatuses()
}

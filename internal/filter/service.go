package filter

import (
	"dating_service/internal/model"
)

type Service struct {
	filterStorage FilterStorage
}

func NewService(filterStorage FilterStorage) *Service {
	return &Service{filterStorage: filterStorage}
}

func (s *Service) CreateFilter(userId, minAge, maxAge, sexId uint, location string) error {
	filter, err := s.filterStorage.GetFilterUser(userId)
	if err != nil {
		return err
	}
	if filter != nil {
		return ErrFilterExists
	}
	if minAge > maxAge {
		return ErrMaxAndMinValue
	}
	err = s.filterStorage.CreateFilter(model.FilterSearch{
		UserID:   userId,
		MinAge:   minAge,
		MaxAge:   maxAge,
		SexID:    sexId,
		Location: location,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetFilter(userID uint) (*model.FilterSearch, error) {
	filter, err := s.filterStorage.GetFilterUser(userID)
	if err != nil {
		return nil, err
	}
	if filter == nil {
		return nil, ErrNotFoundFilter
	}
	return filter, nil
}

func (s *Service) UpdateUserFilter(userID uint, minAge, maxAge, sexId *uint, location *string) error {
	existingFilter, err := s.filterStorage.GetFilterUser(userID)
	if err != nil {
		return err
	}
	if existingFilter == nil {
		return ErrNotFoundFilter
	}

	if minAge != nil {
		existingFilter.MinAge = *minAge
	}
	if maxAge != nil {
		existingFilter.MaxAge = *maxAge
	}
	if sexId != nil {
		existingFilter.SexID = *sexId
	}
	if location != nil {
		existingFilter.Location = *location
	}
	if existingFilter.MinAge > existingFilter.MaxAge {
		return ErrMaxAndMinValue
	}
	err = s.filterStorage.UpdateFilter(*existingFilter)
	return err
}

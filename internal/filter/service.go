package filter

import (
	"dating_service/internal/model"
)

type FilterService struct {
	filterStorage FilterStorage
}

func NewFilterService(filterStorage FilterStorage) *FilterService {
	return &FilterService{filterStorage: filterStorage}
}

func (service *FilterService) CreateFilter(userId, minAge, maxAge, sexId uint, location string) error {
	filter, err := service.filterStorage.GetFilterUser(userId)
	if err != nil {
		return err
	}
	if filter != nil {
		return ErrFilterExists
	}
	if minAge > maxAge {
		return ErrMaxAndMinValue
	}
	err = service.filterStorage.CreateFilter(model.FilterSearch{
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

func (service *FilterService) GetFilter(userID uint) (*model.FilterSearch, error) {
	filter, err := service.filterStorage.GetFilterUser(userID)
	if err != nil {
		return nil, err
	}
	if filter == nil {
		return nil, ErrNotFoundFilter
	}
	return filter, nil
}

func (service *FilterService) UpdateUserFilter(userID uint, minAge, maxAge, sexId *uint, location *string) error {
	existingFilter, err := service.filterStorage.GetFilterUser(userID)
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
	err = service.filterStorage.UpdateFilter(*existingFilter)
	return err
}

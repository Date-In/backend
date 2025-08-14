package filter

import (
	"dating_service/internal/model"
)

type FilterService struct {
	repo *FilterRepository
}

func NewFilterService(repo *FilterRepository) *FilterService {
	return &FilterService{repo: repo}
}

func (service *FilterService) CreateFilter(userId, minAge, maxAge, sexId uint, location string) error {
	filter, err := service.repo.GetFilterUser(userId)
	if err != nil {
		return err
	}
	if filter != nil {
		return ErrFilterExists
	}
	if minAge > maxAge {
		return ErrMaxAndMinValue
	}
	err = service.repo.CreateFilter(model.FilterSearch{
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
	filter, err := service.repo.GetFilterUser(userID)
	if err != nil {
		return nil, err
	}
	if filter == nil {
		return nil, ErrNotFoundFilter
	}
	return filter, nil
}

func (service *FilterService) UpdateUserFilter(userID uint, minAge, maxAge, sexId *uint, location *string) error {
	existingFilter, err := service.repo.GetFilterUser(userID)
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
	err = service.repo.UpdateFilter(*existingFilter)
	return err
}

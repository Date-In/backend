package profile

import (
	"dating_service/internal/model"
	"dating_service/internal/user"
)

type ProfileService struct {
	repo *user.UserRepository
}

func NewProfileService(repo *user.UserRepository) *ProfileService {
	return &ProfileService{repo}
}

func (service *ProfileService) GetInfo(id uint) (*model.User, error) {
	user, err := service.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

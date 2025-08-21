package user

import (
	"dating_service/internal/model"
	"dating_service/pkg/utilits"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo}
}

func (service *UserService) FindUserByPhone(phone string) (*model.User, error) {
	normPhone, err := utilits.FormatPhoneNumber(phone)
	if err != nil {
		return nil, err
	}
	user, err := service.repo.FindByPhone(normPhone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) Create(user *model.User) error {
	err := service.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) FindById(id uint) (*model.User, error) {
	user, err := service.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) Update(id uint, user *model.User) error {
	err := service.repo.Update(id, user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateInterests(userID uint, interests []*model.Interest) ([]*model.Interest, error) {
	err := service.repo.ReplaceInterests(userID, interests)
	if err != nil {
		return nil, err
	}
	user, err := service.repo.FindById(userID)
	if err != nil {
		return nil, err
	}
	return user.Interests, nil
}

func (service *UserService) FindUserWithoutEntity(userID uint) (*model.User, error) {
	user, err := service.repo.FindUserWithoutEntity(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) FindUsersWithFilter(filter *model.FilterSearch, page, pageSize int) (*PaginatedUsersResult, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	res, err := service.repo.FindUsersWithFilter(filter.MinAge, filter.MaxAge, filter.SexID, filter.Location, page, pageSize)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *UserService) ReactivateUser(userID uint) error {
	err := service.repo.ReactivateUser(userID)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) ChangeStatus(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}
	err := service.repo.ChangeStatusUsers(ids)
	if err != nil {
		return err
	}
	return nil
}

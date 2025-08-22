package user

import (
	"dating_service/internal/model"
	"dating_service/pkg/utilits"
)

type UserService struct {
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) *UserService {
	return &UserService{userStorage}
}

func (service *UserService) FindUserByPhone(phone string) (*model.User, error) {
	normPhone, err := utilits.FormatPhoneNumber(phone)
	if err != nil {
		return nil, err
	}
	user, err := service.userStorage.FindByPhone(normPhone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) Create(user *model.User) error {
	err := service.userStorage.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) FindById(id uint) (*model.User, error) {
	user, err := service.userStorage.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) Update(id uint, user *model.User) error {
	err := service.userStorage.Update(id, user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateInterests(userID uint, interests []*model.Interest) ([]*model.Interest, error) {
	err := service.userStorage.ReplaceInterests(userID, interests)
	if err != nil {
		return nil, err
	}
	user, err := service.userStorage.FindById(userID)
	if err != nil {
		return nil, err
	}
	return user.Interests, nil
}

func (service *UserService) FindUserWithoutEntity(userID uint) (*model.User, error) {
	user, err := service.userStorage.FindUserWithoutEntity(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) FindUsersWithFilter(filter *model.FilterSearch, page, pageSize int) ([]*model.User, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	users, count, err := service.userStorage.FindUsersWithFilter(filter.MinAge, filter.MaxAge, filter.SexID, filter.Location, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

func (service *UserService) ReactivateUser(userID uint) error {
	err := service.userStorage.ReactivateUser(userID)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) ChangeStatus(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}
	err := service.userStorage.ChangeStatusUsers(ids)
	if err != nil {
		return err
	}
	return nil
}

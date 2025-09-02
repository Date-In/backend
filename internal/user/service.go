package user

import (
	"dating_service/internal/model"
	"dating_service/pkg/utilits"
)

type Service struct {
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) *Service {
	return &Service{userStorage}
}

func (s *Service) FindUserByPhone(phone string) (*model.User, error) {
	normPhone, err := utilits.FormatPhoneNumber(phone)
	if err != nil {
		return nil, err
	}
	user, err := s.userStorage.FindByPhone(normPhone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Create(user *model.User) error {
	err := s.userStorage.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) FindById(id uint) (*model.User, error) {
	user, err := s.userStorage.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Update(id uint, user *model.User) error {
	err := s.userStorage.Update(id, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateInterests(userID uint, interests []*model.Interest) ([]*model.Interest, error) {
	err := s.userStorage.ReplaceInterests(userID, interests)
	if err != nil {
		return nil, err
	}
	user, err := s.userStorage.FindById(userID)
	if err != nil {
		return nil, err
	}
	return user.Interests, nil
}

func (s *Service) FindUserWithoutEntity(userID uint) (*model.User, error) {
	user, err := s.userStorage.FindUserWithoutEntity(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) FindUsersWithFilter(filter *model.FilterSearch, page, pageSize int) ([]*model.User, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	users, count, err := s.userStorage.FindUsersWithFilter(filter.MinAge, filter.MaxAge, filter.SexID, filter.Location, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

func (s *Service) ReactivateUser(userID uint) error {
	err := s.userStorage.ReactivateUser(userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ChangeStatus(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}
	err := s.userStorage.ChangeStatusUsers(ids)
	if err != nil {
		return err
	}
	return nil
}

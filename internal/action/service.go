package action

import "time"

type ActionsService struct {
	userRepo *ActionsRepository
}

func NewActionsService(userRepo *ActionsRepository) *ActionsService {
	return &ActionsService{userRepo}
}

func (service *ActionsService) Get(userId uint) (*Actions, error) {
	get, err := service.userRepo.Get(userId)
	if err != nil {
		return nil, ErrNotFound
	}
	return get, nil
}

func (service *ActionsService) Update(userId uint, time time.Time) error {
	err := service.userRepo.Update(userId, time)
	if err != nil {
		return err
	}
	return nil
}

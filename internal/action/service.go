package action

import (
	"dating_service/internal/user"
	"time"
)

type ActionsService struct {
	userRepo   *user.UserRepository
	actionRepo *ActionsRepository
}

func NewActionsService(userRepo *user.UserRepository, actionRepo *ActionsRepository) *ActionsService {
	return &ActionsService{userRepo, actionRepo}
}

func (service *ActionsService) Get(userId uint) (*Actions, error) {
	get, err := service.actionRepo.Get(userId)
	if err != nil {
		return nil, ErrNotFound
	}
	return get, nil
}

func (service *ActionsService) Update(userId uint) error {
	if err := service.actionRepo.Update(userId, time.Now()); err != nil {
		return err
	}
	if err := service.userRepo.ReactivateUser(userId); err != nil {
		return err
	}
	return nil
}

func (service *ActionsService) ChangeStatusToNonActive() {
	inactiveThreshold := time.Now().AddDate(-1, 0, 0)
	idsToDeactivate, err := service.actionRepo.GetNonActiveUserIds(inactiveThreshold)
	if err != nil {
		return
	}
	err = service.userRepo.ChangeStatusUsers(idsToDeactivate)
	if err != nil {
		return
	}
}

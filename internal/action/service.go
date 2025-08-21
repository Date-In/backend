package action

import (
	"dating_service/internal/user"
	"time"
)

type ActionsService struct {
	userService *user.UserService
	actionRepo  *ActionsRepository
}

func NewActionsService(userService *user.UserService, actionRepo *ActionsRepository) *ActionsService {
	return &ActionsService{userService, actionRepo}
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
	if err := service.userService.ReactivateUser(userId); err != nil {
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
	err = service.userService.ChangeStatus(idsToDeactivate)
	if err != nil {
		return
	}
}

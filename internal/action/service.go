package action

import (
	"time"
)

type ActionsService struct {
	userProvider  UserProvider
	actionStorage ActionStorage
}

func NewActionsService(userProvider UserProvider, actionStorage ActionStorage) *ActionsService {
	return &ActionsService{userProvider, actionStorage}
}

func (service *ActionsService) Get(userId uint) (*Actions, error) {
	get, err := service.actionStorage.Get(userId)
	if err != nil {
		return nil, ErrNotFound
	}
	return get, nil
}

func (service *ActionsService) Update(userId uint) error {
	if err := service.actionStorage.Update(userId, time.Now()); err != nil {
		return err
	}
	if err := service.userProvider.ReactivateUser(userId); err != nil {
		return err
	}
	return nil
}

func (service *ActionsService) ChangeStatusToNonActive() {
	inactiveThreshold := time.Now().AddDate(-1, 0, 0)
	idsToDeactivate, err := service.actionStorage.GetNonActiveUserIds(inactiveThreshold)
	if err != nil {
		return
	}
	err = service.userProvider.ChangeStatus(idsToDeactivate)
	if err != nil {
		return
	}
}

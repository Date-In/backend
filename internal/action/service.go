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

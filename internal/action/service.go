package action

import (
	"time"
)

type Service struct {
	userProvider  UserProvider
	actionStorage ActionStorage
}

func NewService(userProvider UserProvider, actionStorage ActionStorage) *Service {
	return &Service{userProvider, actionStorage}
}

func (s *Service) ChangeStatusToNonActive() {
	inactiveThreshold := time.Now().AddDate(-1, 0, 0)
	idsToDeactivate, err := s.actionStorage.GetNonActiveUserIds(inactiveThreshold)
	if err != nil {
		return
	}
	err = s.userProvider.ChangeStatus(idsToDeactivate)
	if err != nil {
		return
	}
}

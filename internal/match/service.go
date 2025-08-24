package match

import "dating_service/internal/model"

type MatchService struct {
	matchStorage MatchStorage
}

func NewMatchService(matchStorage MatchStorage) *MatchService {
	return &MatchService{matchStorage}
}

func (s *MatchService) GetUserMatches(currentUserID uint) ([]model.Match, error) {
	matches, err := s.matchStorage.GetAllWithDetails(currentUserID)
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (service *MatchService) Create(userID1, userID2 uint) error {
	var maxId uint
	var minId uint
	if userID1 > userID2 {
		maxId = userID1
		minId = userID2
	} else {
		maxId = userID2
		minId = userID1
	}
	err := service.matchStorage.Create(maxId, minId)
	if err != nil {
		return err
	}
	return nil
}

func (service *MatchService) IsUserInMatch(userID uint, matchID uint) (bool, error) {
	exists, err := service.matchStorage.IsUserInMatch(userID, matchID)
	if err != nil {
		return false, err
	}
	return exists, nil
}

package match

import "dating_service/internal/model"

type Service struct {
	matchStorage MatchStorage
}

func NewService(matchStorage MatchStorage) *Service {
	return &Service{matchStorage}
}

func (s *Service) GetUserMatches(currentUserID uint) ([]model.Match, error) {
	matches, err := s.matchStorage.GetAllWithDetails(currentUserID)
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (s *Service) Create(userID1, userID2 uint) error {
	var maxId uint
	var minId uint
	if userID1 > userID2 {
		maxId = userID1
		minId = userID2
	} else {
		maxId = userID2
		minId = userID1
	}
	err := s.matchStorage.Create(maxId, minId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) IsUserInMatch(userID uint, matchID uint) (bool, error) {
	exists, err := s.matchStorage.IsUserInMatch(userID, matchID)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *Service) GetMatchUserIDs(userID uint) ([]uint, error) {
	matchIds, err := s.matchStorage.GetMatchUserIDs(userID)
	if err != nil {
		return nil, err
	}
	return matchIds, nil
}

func (s *Service) GetUsers(matchID uint) ([]model.User, error) {
	return s.matchStorage.GetUsers(matchID)
}

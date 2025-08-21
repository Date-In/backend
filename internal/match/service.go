package match

import "dating_service/internal/model"

type MatchService struct {
	repo *MatchRepository
}

func NewMatchService(repo *MatchRepository) *MatchService {
	return &MatchService{repo}
}

func (service *MatchService) GetAll(userId uint) ([]model.Match, error) {
	matches, err := service.repo.GetAll(userId)
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
	err := service.repo.Create(maxId, minId)
	if err != nil {
		return err
	}
	return nil
}

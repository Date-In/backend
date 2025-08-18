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

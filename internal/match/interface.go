package match

import "dating_service/internal/model"

type MatchStorage interface {
	Create(uint, uint) error
	GetAllWithDetails(uint) ([]model.Match, error)
	IsUserInMatch(uint, uint) (bool, error)
}

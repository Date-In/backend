package match

import "dating_service/internal/model"

type MatchStorage interface {
	Create(uint, uint) error
	GetAll(uint) ([]model.Match, error)
}

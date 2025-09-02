package match

import "dating_service/internal/model"

type MatchStorage interface {
	Create(uint, uint) error
	GetAllWithDetails(uint) ([]model.Match, error)
	IsUserInMatch(uint, uint) (bool, error)
	GetMatchUserIDs(uint) ([]uint, error)
	GetUsers(matchID uint) ([]model.User, error)
	Delete(matchID uint) error
	GetAll() ([]model.Match, error)
}

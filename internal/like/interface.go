package like

import "dating_service/internal/model"

type UserProvider interface {
	FindUserWithoutEntity(uint) (*model.User, error)
}

type MatchProvider interface {
	Create(uint, uint) error
}

type LikeStorage interface {
	GetLikes(uint) ([]model.Like, error)
	CreateLike(uint, uint) error
	FindLikeByTargetIdAndUserID(uint, uint) (*model.Like, error)
}

package like

import (
	"dating_service/internal/model"
)

type Service struct {
	likeStorage   LikeStorage
	userProvider  UserProvider
	matchProvider MatchProvider
}

func NewService(likeStorage LikeStorage, userProvider UserProvider, matchProvider MatchProvider) *Service {
	return &Service{likeStorage, userProvider, matchProvider}
}

func (s *Service) CreateLike(userId, targetId uint) error {
	entity, err := s.userProvider.FindUserWithoutEntity(targetId)
	if err != nil {
		return err
	}
	if entity == nil {
		return ErrNotFoundUser
	}
	foundRepeatLike, err := s.likeStorage.FindLikeByTargetIdAndUserID(targetId, userId)
	if err != nil {
		return err
	}
	if foundRepeatLike != nil {
		return nil
	}

	found, err := s.likeStorage.FindLikeByTargetIdAndUserID(userId, targetId)
	if err != nil {
		return err
	}
	err = s.likeStorage.CreateLike(userId, targetId)
	if err != nil {
		return err
	}
	if found != nil {
		err = s.matchProvider.Create(userId, targetId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) GetLikes(userId uint) ([]model.Like, error) {
	likesId, err := s.likeStorage.GetLikes(userId)
	if err != nil {
		return nil, err
	}
	return likesId, nil
}

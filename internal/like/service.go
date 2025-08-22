package like

import (
	"dating_service/internal/model"
)

type LikeService struct {
	likeStorage   LikeStorage
	userProvider  UserProvider
	matchProvider MatchProvider
}

func NewLikeService(likeStorage LikeStorage, userProvider UserProvider, matchProvider MatchProvider) *LikeService {
	return &LikeService{likeStorage, userProvider, matchProvider}
}

func (service *LikeService) CreateLike(userId, targetId uint) error {
	entity, err := service.userProvider.FindUserWithoutEntity(targetId)
	if err != nil {
		return err
	}
	if entity == nil {
		return ErrNotFoundUser
	}
	foundRepeatLike, err := service.likeStorage.FindLikeByTargetIdAndUserID(targetId, userId)
	if err != nil {
		return err
	}
	if foundRepeatLike != nil {
		return nil
	}

	found, err := service.likeStorage.FindLikeByTargetIdAndUserID(userId, targetId)
	if err != nil {
		return err
	}
	err = service.likeStorage.CreateLike(userId, targetId)
	if err != nil {
		return err
	}
	if found != nil {
		err = service.matchProvider.Create(userId, targetId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *LikeService) GetLikes(userId uint) ([]model.Like, error) {
	likesId, err := service.likeStorage.GetLikes(userId)
	if err != nil {
		return nil, err
	}
	return likesId, nil
}

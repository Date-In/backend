package like

import (
	"dating_service/internal/match"
	"dating_service/internal/model"
	"dating_service/internal/user"
)

type LikeService struct {
	repo         *LikeRepository
	userService  *user.UserService
	matchService *match.MatchService
}

func NewLikeService(repo *LikeRepository, userService *user.UserService, matchService *match.MatchService) *LikeService {
	return &LikeService{repo, userService, matchService}
}

func (service *LikeService) CreateLike(userId, targetId uint) error {
	entity, err := service.userService.FindUserWithoutEntity(targetId)
	if err != nil {
		return err
	}
	if entity == nil {
		return ErrNotFoundUser
	}
	foundRepeatLike, err := service.repo.FindLikeByTargetIdAndUserID(targetId, userId)
	if err != nil {
		return err
	}
	if foundRepeatLike != nil {
		return nil
	}

	found, err := service.repo.FindLikeByTargetIdAndUserID(userId, targetId)
	if err != nil {
		return err
	}
	err = service.repo.CreateLike(userId, targetId)
	if err != nil {
		return err
	}
	if found != nil {
		err = service.matchService.Create(userId, targetId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *LikeService) GetLikes(userId uint) ([]model.Like, error) {
	likesId, err := service.repo.GetLikes(userId)
	if err != nil {
		return nil, err
	}
	return likesId, nil
}

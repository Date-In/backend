package like

import (
	"dating_service/internal/model"
	"dating_service/internal/user"
	"errors"
)

type LikeService struct {
	repo           *LikeRepository
	userRepository *user.UserRepository
}

func NewLikeService(repo *LikeRepository, userRepository *user.UserRepository) *LikeService {
	return &LikeService{repo, userRepository}
}

func (service *LikeService) CreateLike(userId, targetId uint) error {
	entity, err := service.userRepository.FindUserWithoutEntity(targetId)
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
		//логика создания мэтча
		return errors.New("like created")
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

package profile

import (
	"context"
	"dating_service/internal/model"
	"errors"
	"gorm.io/gorm"
)

type Service struct {
	userProvider  UserProvider
	photoProvider PhotoProvider
	cache         CacheProvider
}

func NewService(userProvider UserProvider, photoProvider PhotoProvider, cache CacheProvider) *Service {
	return &Service{userProvider, photoProvider, cache}
}

func (s *Service) GetInfo(id uint) (*model.User, error) {
	currentUser, err := s.userProvider.FindById(id)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, ErrUserNotFound
	}
	preloadCache(currentUser, s.cache)
	return currentUser, nil
}

func (s *Service) Update(
	id uint,
	name *string,
	age *uint,
	bio *string,
	children *bool,
	city *string,
	height *uint,
	sexID *uint,
	zodiacSignID *uint,
	worldviewID *uint,
	typeOfDatingID *uint,
	educationID *uint,
	attitudeToAlcoholID *uint,
	attitudeToSmokingID *uint,
) (*model.User, error) {
	updateUser, err := s.userProvider.FindById(id)
	if err != nil {
		return nil, err
	}
	if updateUser == nil {
		return nil, ErrUserNotFound
	}
	preloadCache(updateUser, s.cache)
	if name != nil {
		updateUser.Name = *name
	}
	if age != nil {
		updateUser.Age = *age
	}
	if sexID != nil {
		if !s.cache.IsValidSexID(*sexID) {
			return nil, ErrInvalidSexID
		}
		updateUser.SexID = *sexID
	}
	if bio != nil {
		updateUser.Bio = bio
	}
	if city != nil {
		updateUser.City = city
	}
	if children != nil {
		updateUser.Children = children
	}
	if height != nil {
		updateUser.Height = height
	}
	if zodiacSignID != nil {
		if !s.cache.IsValidZodiacSignID(*zodiacSignID) {
			return nil, ErrInvalidZodiacID
		}
		updateUser.ZodiacSignID = zodiacSignID
	}
	if worldviewID != nil {
		if !s.cache.IsValidWorldviewID(*worldviewID) {
			return nil, ErrInvalidWordViewID
		}
		updateUser.WorldviewID = worldviewID
	}
	if typeOfDatingID != nil {
		if !s.cache.IsValidTypeOfDatingID(*typeOfDatingID) {
			return nil, ErrInvalidTypeOfDatingId
		}
		updateUser.TypeOfDatingID = typeOfDatingID
	}
	if educationID != nil {
		if !s.cache.IsValidEducationID(*educationID) {
			return nil, ErrInvalidEducationId
		}
		updateUser.EducationID = educationID
	}
	if attitudeToAlcoholID != nil {
		if !s.cache.IsValidAttitudeToAlcoholID(*attitudeToAlcoholID) {
			return nil, ErrInvalidAttitudeToAlcoholicId
		}
		updateUser.AttitudeToAlcoholID = attitudeToAlcoholID
	}
	if attitudeToSmokingID != nil {
		if !s.cache.IsValidAttitudeToSmokingID(*attitudeToSmokingID) {
			return nil, ErrInvalidAttitudeToSmokingId
		}
		updateUser.AttitudeToSmokingID = attitudeToSmokingID
	}

	err = s.userProvider.Update(id, updateUser)
	if err != nil {
		return nil, err
	}
	userUpd, err := s.userProvider.FindById(id)
	preloadCache(userUpd, s.cache)
	return userUpd, err
}

func (s *Service) UpdateInterests(userID uint, interestIDs []uint) ([]*model.Interest, error) {
	if !s.cache.IsValidInterestIDs(interestIDs) {
		return nil, ErrInvalidInterestId
	}
	interestsToSet := make([]*model.Interest, len(interestIDs))
	for i, id := range interestIDs {
		interestsToSet[i] = &model.Interest{ID: id}
	}
	interest, err := s.userProvider.UpdateInterests(userID, interestsToSet)
	if err != nil {
		return nil, err
	}
	return interest, nil
}

func (s *Service) AddPhoto(ctx context.Context, userID uint, data []byte, fileName string) (*model.Photo, error) {
	count, err := s.photoProvider.CountPhoto(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if count >= 5 {
		return nil, ErrLimitPhoto
	}
	addPhoto, err := s.photoProvider.AddPhoto(ctx, userID, data, fileName)
	if err != nil {
		return nil, err
	}
	return addPhoto, nil
}

func (s *Service) DeletePhoto(ctx context.Context, photoId string, userId uint) error {
	rowsAffected, err := s.photoProvider.DeletePhoto(ctx, photoId, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPhotoNotFound
		}
		return err
	}
	if rowsAffected == 0 {
		return ErrPhotoNotFound
	}
	return nil
}

func (s *Service) UpdateAvatar(photoId string, userID uint) (string, error) {
	newAvatarId, err := s.photoProvider.ChangeAvatarUser(photoId, userID)
	if err != nil {
		return "", err
	}
	return newAvatarId, nil
}

func (s *Service) GetAvatar(userID uint) (string, error) {
	avatarUrl, err := s.photoProvider.FindAvatar(userID)
	if err != nil {
		return "", err
	}
	return avatarUrl, nil
}

func preloadCache(user *model.User, c CacheProvider) {
	if user.SexID != 0 {
		user.Sex = c.GetSexByID(user.SexID)
	}
	if user.ZodiacSignID != nil {
		user.ZodiacSign = c.GetZodiacSignByID(*user.ZodiacSignID)
	}
	if user.WorldviewID != nil {
		user.Worldview = c.GetWorldviewByID(*user.WorldviewID)
	}
	if user.TypeOfDatingID != nil {
		user.TypeOfDating = c.GetTypeOfDatingByID(*user.TypeOfDatingID)
	}
	if user.EducationID != nil {
		user.Education = c.GetEducationByID(*user.EducationID)
	}
	if user.AttitudeToAlcoholID != nil {
		user.AttitudeToAlcohol = c.GetAttitudeToAlcoholByID(*user.AttitudeToAlcoholID)
	}
	if user.AttitudeToSmokingID != nil {
		user.AttitudeToSmoking = c.GetAttitudeToSmokingByID(*user.AttitudeToSmokingID)
	}
	if user.StatusID != 0 {
		user.Status = c.GetStatusByID(user.StatusID)
	}
}

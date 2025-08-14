package profile

import (
	"dating_service/internal/cache"
	"dating_service/internal/model"
	"dating_service/internal/photo"
	"dating_service/internal/user"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ProfileService struct {
	repo            *user.UserRepository
	photoRepository *photo.PhotoRepository
	cache           cache.IReferenceCache
}

func NewProfileService(repo *user.UserRepository, photoRepo *photo.PhotoRepository, cache cache.IReferenceCache) *ProfileService {
	return &ProfileService{repo, photoRepo, cache}
}

func (service *ProfileService) GetInfo(id uint) (*model.User, error) {
	user, err := service.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	PreloadCache(user, service.cache)
	return user, nil
}

func (service *ProfileService) Update(
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
	updateUser, err := service.repo.FindById(id)
	PreloadCache(updateUser, service.cache)
	if err != nil {
		return nil, err
	}
	if updateUser == nil {
		return nil, ErrUserNotFound
	}
	PreloadCache(updateUser, service.cache)
	if name != nil {
		updateUser.Name = *name
	}
	if age != nil {
		updateUser.Age = *age
	}
	if sexID != nil {
		if !service.cache.IsValidSexID(*sexID) {
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
		if !service.cache.IsValidZodiacSignID(*zodiacSignID) {
			return nil, ErrInvalidZodiacID
		}
		updateUser.ZodiacSignID = zodiacSignID
	}
	if worldviewID != nil {
		if !service.cache.IsValidWorldviewID(*worldviewID) {
			return nil, ErrInvalidWordViewID
		}
		updateUser.WorldviewID = worldviewID
	}
	if typeOfDatingID != nil {
		if !service.cache.IsValidTypeOfDatingID(*typeOfDatingID) {
			return nil, ErrInvalidTypeOfDatingId
		}
		updateUser.TypeOfDatingID = typeOfDatingID
	}
	if educationID != nil {
		if !service.cache.IsValidEducationID(*educationID) {
			return nil, ErrInvalidEducationId
		}
		updateUser.EducationID = educationID
	}
	if attitudeToAlcoholID != nil {
		if !service.cache.IsValidAttitudeToAlcoholID(*attitudeToAlcoholID) {
			return nil, ErrInvalidAttitudeToAlcoholicId
		}
		updateUser.AttitudeToAlcoholID = attitudeToAlcoholID
	}
	if attitudeToSmokingID != nil {
		if !service.cache.IsValidAttitudeToSmokingID(*attitudeToSmokingID) {
			return nil, ErrInvalidAttitudeToSmokingId
		}
		updateUser.AttitudeToSmokingID = attitudeToSmokingID
	}

	err = service.repo.Update(id, updateUser)
	if err != nil {
		return nil, err
	}
	userUpd, err := service.repo.FindById(id)
	PreloadCache(userUpd, service.cache)
	return updateUser, err
}

func (service *ProfileService) UpdateInterests(userID uint, interestIDs []uint) ([]*model.Interest, error) {
	if !service.cache.IsValidInterestIDs(interestIDs) {
		return nil, ErrInvalidInterestId
	}
	interestsToSet := make([]*model.Interest, len(interestIDs))
	for i, id := range interestIDs {
		interestsToSet[i] = &model.Interest{ID: id}
	}

	err := service.repo.ReplaceInterests(userID, interestsToSet)
	if err != nil {
		return nil, fmt.Errorf("ошибка в репозитории при замене интересов для userID %d: %w", userID, err)
	}

	updatedUser, err := service.repo.FindById(userID)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить профиль после обновления интересов: %w", err)
	}

	return updatedUser.Interests, nil
}

func (service *ProfileService) AddPhoto(fileName, fileType string, data []byte, userID uint) (*string, error) {
	count, err := service.photoRepository.CountPhoto(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if count >= 5 {
		return nil, ErrLimitPhoto
	}
	photo := model.NewPhoto(fileName, fileType, data, userID)
	err = service.photoRepository.Save(photo)
	if err != nil {
		return nil, err
	}
	return &photo.ID, nil
}

func (service *ProfileService) DeletePhoto(photoId string, userId uint) error {
	rowsAffected, err := service.photoRepository.DeleteById(photoId, userId)
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

func (service *ProfileService) UpdateAvatar(photoId string, userID uint) (string, error) {
	newAvatarId, err := service.photoRepository.ChangeAvatarUser(userID, photoId)
	if err != nil {
		return "", err
	}
	return newAvatarId, nil
}

func (service *ProfileService) GetAvatar(userID uint) (string, error) {
	avatarId, err := service.photoRepository.FindAvatar(userID)
	if err != nil {
		return "", err
	}
	return "http://localhost:8081/photo/" + avatarId, nil
}

func PreloadCache(user *model.User, c cache.IReferenceCache) {
	if user.SexID == 0 {
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

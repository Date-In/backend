package photo

import (
	"bytes"
	"context"
	"dating_service/internal/model"
)

type PhotoService struct {
	photoStorage PhotoStorage
	s3Provider   S3Provider
}

func NewPhotoService(photoStorage PhotoStorage, s3Provider S3Provider) *PhotoService {
	return &PhotoService{
		photoStorage, s3Provider,
	}
}

func (service *PhotoService) AddPhoto(ctx context.Context, userID uint, data []byte, fileName string) (*model.Photo, error) {
	fileReader := bytes.NewReader(data)
	url, objectKey, err := service.s3Provider.UploadFile(ctx, fileReader, fileName)
	if err != nil {
		return nil, err
	}
	photo := model.NewPhoto(objectKey, url, userID)
	err = service.photoStorage.Save(photo)
	if err != nil {
		return nil, err
	}
	return photo, nil
}

func (service *PhotoService) DeletePhoto(ctx context.Context, photoId string, userId uint) (int, error) {
	err := service.s3Provider.Delete(ctx, photoId)
	if err != nil {
		return 0, err
	}
	rawAffected, err := service.photoStorage.DeleteById(photoId, userId)
	if err != nil {
		return 0, err
	}
	return int(rawAffected), nil
}

func (service *PhotoService) CountPhoto(userID uint) (int, error) {
	countPhoto, err := service.photoStorage.CountPhoto(userID)
	if err != nil {
		return 0, err
	}
	return countPhoto, nil
}

func (service *PhotoService) ChangeAvatarUser(photoId string, userID uint) (string, error) {
	newAvatarId, err := service.photoStorage.ChangeAvatarUser(userID, photoId)
	if err != nil {
		return "", err
	}
	return newAvatarId, nil
}

func (service *PhotoService) FindAvatar(userID uint) (string, error) {
	avatar, err := service.photoStorage.FindAvatar(userID)
	if err != nil {
		return "", err
	}
	return avatar.Url, err
}

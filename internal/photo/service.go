package photo

import (
	"bytes"
	"context"
	"dating_service/internal/filestorage"
	"dating_service/internal/model"
)

type PhotoService struct {
	repository    *PhotoRepository
	S3Filestorage *filestorage.S3FileStorage
}

func NewPhotoService(repository *PhotoRepository, S3Filestorage *filestorage.S3FileStorage) *PhotoService {
	return &PhotoService{
		repository, S3Filestorage,
	}
}

func (service *PhotoService) AddPhoto(ctx context.Context, userID uint, data []byte, fileName string) (*model.Photo, error) {
	fileReader := bytes.NewReader(data)
	url, objectKey, err := service.S3Filestorage.UploadFile(ctx, fileReader, fileName)
	if err != nil {
		return nil, err
	}
	photo := model.NewPhoto(objectKey, url, userID)
	err = service.repository.Save(photo)
	if err != nil {
		return nil, err
	}
	return photo, nil
}

func (service *PhotoService) DeletePhoto(ctx context.Context, photoId string, userId uint) (int, error) {
	err := service.S3Filestorage.Delete(ctx, photoId)
	if err != nil {
		return 0, err
	}
	rawAffected, err := service.repository.DeleteById(photoId, userId)
	if err != nil {
		return 0, err
	}
	return int(rawAffected), nil
}

func (service *PhotoService) CountPhoto(userID uint) (int, error) {
	countPhoto, err := service.repository.CountPhoto(userID)
	if err != nil {
		return 0, err
	}
	return countPhoto, nil
}

func (service *PhotoService) ChangeAvatarUser(photoId string, userID uint) (string, error) {
	newAvatarId, err := service.repository.ChangeAvatarUser(userID, photoId)
	if err != nil {
		return "", err
	}
	return newAvatarId, nil
}

func (service *PhotoService) FindAvatar(userID uint) (string, error) {
	avatar, err := service.repository.FindAvatar(userID)
	if err != nil {
		return "", err
	}
	return avatar.Url, err
}

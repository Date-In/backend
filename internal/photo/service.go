package photo

import (
	"bytes"
	"context"
	"dating_service/internal/model"
)

type Service struct {
	photoStorage PhotoStorage
	s3Provider   S3Provider
}

func NewService(photoStorage PhotoStorage, s3Provider S3Provider) *Service {
	return &Service{
		photoStorage, s3Provider,
	}
}

func (s *Service) AddPhoto(ctx context.Context, userID uint, data []byte, fileName string) (*model.Photo, error) {
	fileReader := bytes.NewReader(data)
	url, objectKey, err := s.s3Provider.UploadFile(ctx, fileReader, fileName)
	if err != nil {
		return nil, err
	}
	photo := model.NewPhoto(objectKey, url, userID)
	err = s.photoStorage.Save(photo)
	if err != nil {
		return nil, err
	}
	return photo, nil
}

func (s *Service) DeletePhoto(ctx context.Context, photoId string, userId uint) (int, error) {
	err := s.s3Provider.Delete(ctx, photoId)
	if err != nil {
		return 0, err
	}
	rawAffected, err := s.photoStorage.DeleteById(photoId, userId)
	if err != nil {
		return 0, err
	}
	return int(rawAffected), nil
}

func (s *Service) CountPhoto(userID uint) (int, error) {
	countPhoto, err := s.photoStorage.CountPhoto(userID)
	if err != nil {
		return 0, err
	}
	return countPhoto, nil
}

func (s *Service) ChangeAvatarUser(photoId string, userID uint) (string, error) {
	newAvatarId, err := s.photoStorage.ChangeAvatarUser(userID, photoId)
	if err != nil {
		return "", err
	}
	return newAvatarId, nil
}

func (s *Service) FindAvatar(userID uint) (string, error) {
	avatar, err := s.photoStorage.FindAvatar(userID)
	if err != nil {
		return "", err
	}
	return avatar.Url, err
}

package photo

import (
	"dating_service/internal/model"
	"fmt"
)

type PhotoService struct {
	repo *PhotoRepository
}

func NewPhotoService(repo *PhotoRepository) *PhotoService {
	return &PhotoService{repo}
}

func (s *PhotoService) GetPhoto(uuid string) (*model.Photo, error) {
	photo, err := s.repo.GetById(uuid)
	if err != nil {
		return nil, err
	}
	if photo == nil {
		return nil, ErrPhotoNotFound
	}
	return photo, nil
}

func (s *PhotoService) GetUserPhotoURLs(userId uint) ([]string, error) {
	ids, err := s.repo.FindAllIDs(userId)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return []string{}, nil
	}

	urls := make([]string, len(ids))
	for i, id := range ids {
		urls[i] = fmt.Sprintf("http://localhost:8081/photo/%s", id)
	}

	return urls, nil
}

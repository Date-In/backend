package photo

import (
	"context"
	"dating_service/internal/model"
	"io"
)

type S3Provider interface {
	UploadFile(context.Context, io.Reader, string) (string, string, error)
	Delete(context.Context, string) error
}

type PhotoStorage interface {
	Save(*model.Photo) error
	GetById(string) (*model.Photo, error)
	CountPhoto(uint) (int, error)
	DeleteById(string, uint) (int64, error)
	FindAllIDs(uint) ([]string, error)
	FindAvatar(uint) (*model.Photo, error)
	FindUserPhotoWithoutAvatar(uint) ([]string, error)
	ChangeAvatarUser(uint, string) (string, error)
}

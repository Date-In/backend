package filestorage

import "errors"

var (
	ErrUploadFile = errors.New("Upload file failed")
	ErrDeleteFile = errors.New("Delete file failed")
)

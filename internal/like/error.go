package like

import "errors"

var (
	ErrNotFoundUser = errors.New("User not found")
	ErrNotFoundLike = errors.New("Like not found")
)

package recommendations

import "errors"

var (
	ErrUserNotFound   = errors.New("User not found")
	ErrFilterNotFound = errors.New("Filter not found")
	ErrQueryParam     = errors.New("page and page size mast be number")
)

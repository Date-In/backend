package filter

import "errors"

var (
	ErrFilterExists   = errors.New("filter already exists")
	ErrNotFoundFilter = errors.New("filter not found")
	ErrMaxAndMinValue = errors.New("the maximum age is less than the minimum age")
)

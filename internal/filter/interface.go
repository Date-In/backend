package filter

import "dating_service/internal/model"

type FilterStorage interface {
	GetFilterUser(uint) (*model.FilterSearch, error)
	CreateFilter(model.FilterSearch) error
	UpdateFilter(model.FilterSearch) error
}

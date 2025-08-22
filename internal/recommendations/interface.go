package recommendations

import "dating_service/internal/model"

type UserProvider interface {
	FindUserWithoutEntity(uint) (*model.User, error)
	FindUsersWithFilter(*model.FilterSearch, int, int) ([]*model.User, int64, error)
}
type FilterProvider interface {
	GetFilter(uint) (*model.FilterSearch, error)
}

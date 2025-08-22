package user

import "dating_service/internal/model"

type UserStorage interface {
	Create(*model.User) error
	FindByPhone(string) (*model.User, error)
	FindById(uint) (*model.User, error)
	Update(uint, *model.User) error
	ReplaceInterests(uint, []*model.Interest) error
	FindUsersWithFilter(uint, uint, uint, string, int, int) ([]*model.User, int64, error)
	FindUserWithoutEntity(uint) (*model.User, error)
	ChangeStatusUsers([]uint) error
	ReactivateUser(uint) error
	GetStatusUser(uint) (uint, error)
}

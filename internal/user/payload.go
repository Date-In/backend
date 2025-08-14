package user

import "dating_service/internal/model"

type PaginatedUsersResult struct {
	Users      []*model.User
	TotalCount int64
}

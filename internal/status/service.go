package status

import "dating_service/internal/user"

type StatusService struct {
	userRepo *user.UserRepository
}

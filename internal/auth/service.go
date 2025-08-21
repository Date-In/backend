package auth

import (
	"dating_service/internal/cache"
	"dating_service/internal/model"
	"dating_service/internal/user"
	"dating_service/pkg/JWT"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const TokenDuration = time.Hour * 72

type AuthService struct {
	usersService *user.UserService
	cache        cache.IReferenceCache
	tokenManager *JWT.JWT
}

func NewAuthService(service *user.UserService, cache cache.IReferenceCache, tm *JWT.JWT) *AuthService {
	return &AuthService{usersService: service, cache: cache, tokenManager: tm}
}

func (service *AuthService) Register(phone string, name string, password string, sexID uint, age uint) (*string, error) {
	exist, err := service.usersService.FindUserByPhone(phone)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return nil, ErrUserAlreadyExists
	}
	if !service.cache.IsValidSexID(sexID) {
		return nil, ErrInvalidSexID
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	createdUser := &model.User{
		Phone:    phone,
		Name:     name,
		Password: string(hashedPassword),
		SexID:    sexID,
		Age:      age,
		StatusID: 1,
	}

	err = service.usersService.Create(createdUser)
	if err != nil {
		return nil, err
	}

	token, err := service.tokenManager.GenerateToken(createdUser.ID, TokenDuration)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (service *AuthService) Login(phone, password string) (*string, error) {
	exist, err := service.usersService.FindUserByPhone(phone)
	if err != nil {
		return nil, err
	}
	if exist == nil {
		return nil, ErrIncorrectPasswordOrPhone
	}

	err = bcrypt.CompareHashAndPassword([]byte(exist.Password), []byte(password))
	if err != nil {
		return nil, ErrIncorrectPasswordOrPhone
	}

	token, err := service.tokenManager.GenerateToken(exist.ID, TokenDuration)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

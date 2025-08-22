package auth

import (
	"dating_service/internal/model"
	"dating_service/pkg/JWT"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const TokenDuration = time.Hour * 72

type AuthService struct {
	userProvider UserProvider
	cache        CacheProvider
	tokenManager *JWT.JWT
}

func NewAuthService(userProvider UserProvider, cache CacheProvider, tm *JWT.JWT) *AuthService {
	return &AuthService{userProvider: userProvider, cache: cache, tokenManager: tm}
}

func (service *AuthService) Register(phone string, name string, password string, sexID uint, age uint) (*string, error) {
	exist, err := service.userProvider.FindUserByPhone(phone)
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

	err = service.userProvider.Create(createdUser)
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
	exist, err := service.userProvider.FindUserByPhone(phone)
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

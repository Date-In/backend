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
	repo         *user.UserRepository
	cache        cache.IReferenceCache
	tokenManager *JWT.JWT
}

func NewAuthService(repo *user.UserRepository, cache cache.IReferenceCache, tm *JWT.JWT) *AuthService {
	return &AuthService{repo: repo, cache: cache, tokenManager: tm}
}

func (service *AuthService) Register(phone string, name string, password string, sexID uint, age uint) (*string, error) {
	exist, err := service.repo.FindByPhone(phone)
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
	}

	err = service.repo.Create(createdUser)
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
	user, err := service.repo.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrIncorrectPasswordOrPhone
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrIncorrectPasswordOrPhone
	}

	token, err := service.tokenManager.GenerateToken(user.ID, TokenDuration)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

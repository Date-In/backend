package auth

import (
	"dating_service/internal/model"
	"dating_service/pkg/JWT"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const TokenDuration = time.Hour * 72

type Service struct {
	userProvider UserProvider
	cache        CacheProvider
	tokenManager *JWT.JWT
}

func NewService(userProvider UserProvider, cache CacheProvider, tm *JWT.JWT) *Service {
	return &Service{userProvider: userProvider, cache: cache, tokenManager: tm}
}

func (s *Service) Register(phone string, name string, password string, sexID uint, age uint) (*string, error) {
	exist, err := s.userProvider.FindUserByPhone(phone)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return nil, ErrUserAlreadyExists
	}
	if !s.cache.IsValidSexID(sexID) {
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

	err = s.userProvider.Create(createdUser)
	if err != nil {
		return nil, err
	}

	token, err := s.tokenManager.GenerateToken(createdUser.ID, TokenDuration)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *Service) Login(phone, password string) (*string, error) {
	exist, err := s.userProvider.FindUserByPhone(phone)
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

	token, err := s.tokenManager.GenerateToken(exist.ID, TokenDuration)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

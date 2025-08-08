package auth

import (
	"dating_service/configs"
	"dating_service/internal/cache"
	"dating_service/internal/model"
	"dating_service/internal/user"
	"dating_service/pkg/JWT"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService struct {
	conf  *configs.Config
	repo  *user.UserRepository
	cache cache.IReferenceCache
}

func NewAuthService(conf *configs.Config, repo *user.UserRepository, cache cache.IReferenceCache) *AuthService {
	return &AuthService{conf: conf, repo: repo, cache: cache}
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

	err = service.repo.Create(&model.User{
		Phone:    phone,
		Name:     name,
		Password: string(hashedPassword),
		SexID:    sexID,
		Age:      age,
	})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	jwt, err := JWT.NewJWT(service.conf.SecretToken.Token).GenerateToken(&JWT.JWTData{
		Phone: phone,
	})
	return &jwt, nil
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
	jwt, err := JWT.NewJWT(service.conf.SecretToken.Token).GenerateToken(&JWT.JWTData{
		Phone: phone,
	})
	if err != nil {
		return nil, err
	}
	return &jwt, nil
}

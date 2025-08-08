package JWT

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{secret}
}

func (j *JWT) GenerateToken(userID uint, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(duration).Unix(),
		"iat": time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j *JWT) ParseToken(token string) (uint, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(j.Secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		userIDFloat, ok := claims["sub"].(float64)
		if !ok {
			return 0, ErrInvalidToken
		}
		return uint(userIDFloat), nil
	}

	return 0, ErrInvalidToken
}

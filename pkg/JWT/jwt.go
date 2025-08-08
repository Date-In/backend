package JWT

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	Secret string
}

type JWTData struct {
	Phone string
}

func NewJWT(secret string) *JWT {
	return &JWT{secret}
}

func (j *JWT) GenerateToken(data *JWTData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": data.Phone,
	})
	token, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j *JWT) ParseToken(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	email := t.Claims.(jwt.MapClaims)["phone"].(string)
	return t.Valid, &JWTData{email}
}

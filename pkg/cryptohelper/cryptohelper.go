package cryptohelper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"dating_service/configs"
	"encoding/base64"
	"fmt"
	"io"
)

type Service struct {
	key []byte
}

func NewService(conf *configs.Config) *Service {
	return &Service{
		key: []byte(conf.SecretKeyCryptMessage.Key),
	}
}

func (s *Service) EncryptString(plain string) (string, error) {
	block, err := aes.NewCipher(s.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plain), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (s *Service) DecryptString(encrypted string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(s.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(data) < gcm.NonceSize() {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce := data[:gcm.NonceSize()]
	ciphertext := data[gcm.NonceSize():]

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plain), nil
}

func (s *Service) CompareEncryptedWithPlain(encrypted, plain string) (bool, error) {
	decrypted, err := s.DecryptString(encrypted)
	if err != nil {
		return false, err
	}
	return decrypted == plain, nil
}

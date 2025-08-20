package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	*BdConfig
	*SecretToken
	*S3
}

type BdConfig struct {
	DSN string
}

type SecretToken struct {
	Token string
}

type S3 struct {
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	Region          string
	Endpoint        string
}

func NewConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		BdConfig: &BdConfig{
			DSN: os.Getenv("DSN"),
		},
		SecretToken: &SecretToken{
			Token: os.Getenv("TOKEN"),
		},
		S3: &S3{
			AccessKeyID:     os.Getenv("YANDEX_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("YANDEX_SECRET_ACCESS_KEY"),
			BucketName:      os.Getenv("YANDEX_BUCKET_NAME"),
			Region:          os.Getenv("YANDEX_REGION"),
			Endpoint:        os.Getenv("YANDEX_ENDPOINT"),
		},
	}
}

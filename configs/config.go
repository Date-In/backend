package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	*BdConfig
	*SecretToken
}

type BdConfig struct {
	DSN string
}

type SecretToken struct {
	Token string
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
	}
}

package configs

import (
	"github.com/joho/godotenv"
	"log"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		BdConfig: &BdConfig{
			DSN: os.Getenv("DSN"),
		},
		SecretToken: &SecretToken{
			Token: os.Getenv("TOKEN"),
		},
	}
}

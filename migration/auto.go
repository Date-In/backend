package main

import (
	"dating_service/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func main() {
	temp := time.Now()
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&model.User{},
		&model.Sex{},
		&model.Education{},
		&model.ZodiacSign{},
		&model.Worldview{},
		&model.TypeOfDating{},
		&model.AttitudeToAlcohol{},
		&model.AttitudeToSmoking{},
		&model.Status{},
		&model.FilterSearch{},
		&model.Interest{},
		&model.Like{},
		&model.Match{},
		&model.Message{},
		&model.Photo{},
	)
	log.Printf("Migration complete time: %.3fs", time.Since(temp).Seconds())
}

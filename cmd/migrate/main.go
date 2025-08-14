package main

import (
	"dating_service/internal/action"
	"dating_service/internal/model"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("Starting database migration...")
	temp := time.Now()

	_ = godotenv.Load()

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database for migration: %v", err)
	}

	log.Println("Database connected. Running AutoMigrate...")
	err = db.AutoMigrate(
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
		action.Actions{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Printf("Migration completed successfully in %.3fs", time.Since(temp).Seconds())
}

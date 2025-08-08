package db

import (
	"dating_service/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	PgDb *gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.BdConfig.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &Db{db}
}

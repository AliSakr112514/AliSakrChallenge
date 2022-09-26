package db

import (
	"Challenge/config"
	Log "Challenge/internal/adapters/Logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDatabaseConnection(config *config.Configurations) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Database.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
		return nil, err
	}
	Log.Info.Printf("Database is connected successfully ")
	return db, err
}

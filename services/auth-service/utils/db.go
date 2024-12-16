package utils

import (
	"github.com/hungrynoodlehead/photos/services/auth-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getConnectionString() string {
	return "host=localhost port=5432 dbname=photos"
}

func InitDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getConnectionString()))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Credentials{})
	db.AutoMigrate(&models.Sessions{})
	db.AutoMigrate(&models.TokenPairs{})

	return db, nil
}

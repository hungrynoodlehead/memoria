package utils

import (
	"github.com/hungrynoodlehead/memoria/services/album-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
	Config *Config
}

func NewDB(config *Config) (*DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetConnectonString()))
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Album{}, &models.Photo{})
	if err != nil {
		return nil, err
	}

	return &DB{
		DB:     db,
		Config: config,
	}, nil
}

package utils

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DB struct {
	*mongo.Database
	Config *Config
	logger *Logger
}

func NewDB(config *Config, logger *Logger) (*DB, error) {
	var db DB
	db.Config = config
	db.logger = logger

	client, err := mongo.Connect(options.Client().ApplyURI(db.Config.GetDBString()))
	if err != nil {
		return nil, err
	}
	db.logger.Info("Successfully connected to MongoDB")
	database := client.Database("photo_service")
	db.Database = database

	return &db, nil
}

package photo_repository

import (
	"fmt"
	"github.com/hungrynoodlehead/memoria/services/storage-service/utils"
)

type PhotoRepository struct {
	DB       *utils.DB
	Storage  *utils.Storage
	Logger   *utils.Logger
	Producer *utils.BrokerProducer
}

func NewPhotoRepository(db *utils.DB, storage *utils.Storage, logger *utils.Logger, producer *utils.BrokerProducer) *PhotoRepository {
	return &PhotoRepository{
		DB:       db,
		Storage:  storage,
		Logger:   logger,
		Producer: producer,
	}
}

func getObjectName(userId uint64, photoId string) string {
	return fmt.Sprintf("%d/%s", userId, photoId)
}

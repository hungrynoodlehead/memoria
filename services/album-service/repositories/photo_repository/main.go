package photo_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/utils"

type PhotoRepository struct {
	DB       *utils.DB
	Producer *utils.MessageProducer
}

func NewPhotoRepository(db *utils.DB, producer *utils.MessageProducer) *PhotoRepository {
	return &PhotoRepository{
		DB:       db,
		Producer: producer,
	}
}

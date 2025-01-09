package share_repository

import (
	"github.com/hungrynoodlehead/memoria/services/album-service/repositories/album_repository"
	"github.com/hungrynoodlehead/memoria/services/album-service/repositories/photo_repository"
	"github.com/hungrynoodlehead/memoria/services/album-service/utils"
)

type ShareRepository struct {
	DB              *utils.DB
	Producer        *utils.MessageProducer
	PhotoRepository *photo_repository.PhotoRepository
	AlbumRepository *album_repository.AlbumRepository
}

func NewShareRepository(
	db *utils.DB,
	producer *utils.MessageProducer,
	photoRepository *photo_repository.PhotoRepository,
	albumRepository *album_repository.AlbumRepository,
) *ShareRepository {
	return &ShareRepository{
		DB:              db,
		Producer:        producer,
		PhotoRepository: photoRepository,
		AlbumRepository: albumRepository,
	}
}

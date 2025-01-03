package album_repository

import (
	"github.com/hungrynoodlehead/memoria/services/album-service/repositories/photo_repository"
	"github.com/hungrynoodlehead/memoria/services/album-service/utils"
)

type AlbumRepository struct {
	DB              *utils.DB
	PhotoRepository *photo_repository.PhotoRepository
}

func NewAlbumRepository(db *utils.DB, photoRepository *photo_repository.PhotoRepository) *AlbumRepository {
	return &AlbumRepository{
		DB:              db,
		PhotoRepository: photoRepository,
	}
}

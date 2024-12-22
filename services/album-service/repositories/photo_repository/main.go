package photo_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/utils"

type PhotoRepository struct {
	DB *utils.DB
}

func NewPhotoRepository(db *utils.DB) *PhotoRepository {
	return &PhotoRepository{
		DB: db,
	}
}

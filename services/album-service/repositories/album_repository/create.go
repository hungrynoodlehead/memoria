package album_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/models"

func (r *AlbumRepository) Create(albumName string, albumDesc string, ownerID uint64) (models.Album, error) {
	album := models.Album{
		Name:        albumName,
		Description: albumDesc,
		OwnerID:     ownerID,
	}

	err := r.DB.Create(&album).Error
	if err != nil {
		return models.Album{}, err
	}
	return album, nil
}

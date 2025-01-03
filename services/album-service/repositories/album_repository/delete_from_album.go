package album_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/models"

func (r *AlbumRepository) DeleteFromAlbum(album *models.Album, photo *models.Photo) (models.Album, error) {
	var newPhotos []*models.Photo
	for _, v := range album.Photos {
		if v != photo {
			newPhotos = append(newPhotos, v)
		}
	}

	album.Photos = newPhotos
	err := r.DB.Save(album).Error
	if err != nil {
		return models.Album{}, err
	}
	return *album, nil
}

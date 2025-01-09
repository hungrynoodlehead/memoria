package album_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/models"

func (r *AlbumRepository) DeleteFromAlbum(album *models.Album, photo *models.Photo, purge bool) (models.Album, error) {
	var newPhotos []*models.Photo
	var found bool
	for _, v := range album.Photos {
		if v != photo {
			newPhotos = append(newPhotos, v)
		} else {
			found = true
			if purge {
				err := r.PhotoRepository.DeletePhoto(*v)
				if err != nil {
					return models.Album{}, err
				}
			}
		}
	}

	if !found {
		return models.Album{}, ErrPhotoNotInAlbum
	}

	album.Photos = newPhotos
	err := r.DB.Save(album).Error
	if err != nil {
		return models.Album{}, err
	}
	return *album, nil
}

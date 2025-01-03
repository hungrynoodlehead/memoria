package album_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/models"

func (r *AlbumRepository) AddToAlbum(album models.Album, photo models.Photo) (models.Album, error) {
	// TODO: Check if photo already exists in album
	album.Photos = append(album.Photos, &photo)
	err := r.DB.Save(&album).Error
	if err != nil {
		return models.Album{}, err
	}
	return album, nil
}

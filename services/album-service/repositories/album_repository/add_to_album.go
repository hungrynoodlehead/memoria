package album_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/models"

func (r *AlbumRepository) AddToAlbum(album models.Album, photo models.Photo) error {
	album.Photos = append(album.Photos, &photo)
	err := r.DB.Save(&album).Error
	if err != nil {
		return err
	}
	return nil
}

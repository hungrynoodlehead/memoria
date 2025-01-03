package album_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/models"

func (r *AlbumRepository) DeleteAlbum(album *models.Album, purge bool) error {
	if purge {
		for _, photo := range album.Photos {
			err := r.PhotoRepository.DeletePhoto(*photo)
			if err != nil {
				return err
			}
		}
	}

	if err := r.DB.Delete(album).Error; err != nil {
		return err
	}
	return nil
}

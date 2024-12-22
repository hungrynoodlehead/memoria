package album_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/models"

func (r *AlbumRepository) GetByID(id uint) (*models.Album, error) {
	var album models.Album
	if err := r.DB.First(&album, id).Error; err != nil {
		return nil, err
	}
	return &album, nil
}

package album_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/models"

func (r *AlbumRepository) GetByID(id uint64, preload ...string) (*models.Album, error) {
	query := r.DB.Model(&models.Album{})

	for _, model := range preload {
		query = query.Preload(model)
	}

	var album models.Album
	if err := query.First(&album, id).Error; err != nil {
		return nil, err
	}
	return &album, nil
}

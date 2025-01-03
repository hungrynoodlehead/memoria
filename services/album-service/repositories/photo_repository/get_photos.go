package photo_repository

import (
	"github.com/hungrynoodlehead/memoria/services/album-service/models"
)

func (r *PhotoRepository) GetUserPhotos(userId uint64) ([]*models.Photo, error) {
	var photos []*models.Photo

	err := r.DB.Model(models.Photo{}).Where("user_id = ?", userId).Find(&photos).Error
	if err != nil {
		return nil, err
	}

	return photos, nil
}

package photo_repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/hungrynoodlehead/memoria/services/album-service/models"
	"gorm.io/gorm"
)

func (r *PhotoRepository) GetPhoto(id uuid.UUID) (models.Photo, error) {
	var photo models.Photo
	err := r.DB.Model(&models.Photo{}).First(&photo, &models.Photo{
		UUID: id.String(),
	}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return photo, ErrPhotoNotFound
		}
		return photo, err
	}

	return photo, nil
}

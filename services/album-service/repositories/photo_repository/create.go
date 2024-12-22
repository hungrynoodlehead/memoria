package photo_repository

import (
	"github.com/google/uuid"
	"github.com/hungrynoodlehead/memoria/services/album-service/models"
)

func (r *PhotoRepository) Create(photoUUID uuid.UUID, kind models.PhotoKind) (models.Photo, error) {
	photo := models.Photo{
		UUID: photoUUID.String(),
		Kind: kind,
	}

	err := r.DB.Create(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

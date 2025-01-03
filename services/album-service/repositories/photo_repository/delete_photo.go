package photo_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/models"

func (r *PhotoRepository) DeletePhoto(photo models.Photo) error {
	err := r.DB.Delete(&photo).Error
	if err != nil {
		return err
	}

	//TODO: SEND KAFKA MESSAGE

	return nil
}

package share_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/models"

func (r *ShareRepository) UpdateShareAccessSettings(share *models.Share, value models.SharePermissions) error {
	share.Permissions = value

	if err := r.DB.Save(share).Error; err != nil {
		return err
	}
	return nil
}

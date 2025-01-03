package share_repository

import "github.com/hungrynoodlehead/memoria/services/album-service/models"

func (r *ShareRepository) TerminateShare(share *models.Share) error {
	r.DB.Delete(&share)
	//TODO: Send message to kafka
	return nil
}

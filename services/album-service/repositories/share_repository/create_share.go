package share_repository

import (
	"github.com/google/uuid"
	"github.com/hungrynoodlehead/memoria/services/album-service/models"
	"time"
)

func (r *ShareRepository) CreateShare(album *models.Album, expiresAt time.Time, permissions models.SharePermissions) (models.Share, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return models.Share{}, err
	}

	share := models.Share{
		TokenID:     tokenID,
		OwnerID:     album.OwnerID,
		Album:       *album,
		AccessedAt:  time.Now(),
		ExpiresAt:   expiresAt,
		Status:      models.ShareStatusActive,
		Permissions: permissions,
	}

	//TODO: Send message to Kafka

	r.DB.Create(&share)
	return share, nil
}

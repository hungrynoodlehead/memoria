package photo_repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/hungrynoodlehead/memoria/services/storage-service/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *PhotoRepository) GetPhotoData(id uuid.UUID) (*models.Photo, error) {
	filter := bson.M{"_id": id.String()}
	var photoData *models.Photo
	err := r.DB.Collection("photos").FindOne(context.Background(), filter).Decode(&photoData)
	if err != nil {
		return nil, err
	}

	return photoData, nil
}

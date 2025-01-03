package photo_repository

import (
	"context"
	"fmt"
	"github.com/hungrynoodlehead/memoria/services/storage-service/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *PhotoRepository) GetUserPhotos(userId uint64) ([]models.Photo, error) {
	filter := bson.D{{"user_id", int64(userId)}}

	cursor, err := r.DB.Collection("photos").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var results []models.Photo
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	fmt.Print(results)

	return results, nil
}

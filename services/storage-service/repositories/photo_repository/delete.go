package photo_repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func (r *PhotoRepository) DeletePhoto(photoId uuid.UUID) error {
	data, err := r.GetPhotoData(photoId)
	if err != nil {
		return err
	}

	_, err = r.DB.Collection("photos").DeleteOne(context.Background(), data)
	if err != nil {
		return err
	}

	err = r.Storage.RemoveObject(
		context.Background(),
		"photos",
		getObjectName(data.UserID, photoId.String()),
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

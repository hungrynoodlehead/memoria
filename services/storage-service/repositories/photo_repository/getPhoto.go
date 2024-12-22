package photo_repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func (r *PhotoRepository) GetPhotoFile(photoId uuid.UUID) (*minio.Object, error) {
	data, err := r.GetPhotoData(photoId)
	if err != nil {
		return nil, err
	}

	userId := data.UserID

	file, err := r.Storage.GetObject(
		context.Background(),
		"photos",
		getObjectName(userId, photoId.String()),
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, err
	}

	return file, nil
}

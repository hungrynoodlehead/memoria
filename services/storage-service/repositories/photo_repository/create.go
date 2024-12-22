package photo_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/hungrynoodlehead/memoria/services/storage-service/models"
	"github.com/hungrynoodlehead/memoria/services/storage-service/utils"
	"github.com/minio/minio-go/v7"
	"io"
	"time"
)

func (r *PhotoRepository) CreatePhoto(userId uint64, file io.Reader, fileName string, fileSize int64, kind models.PhotoKind, contentType string, albumId uint) (models.Photo, error) {
	type BrokerMessage struct {
		ID        string           `json:"id"`
		Timestamp string           `json:"timestamp"`
		Kind      models.PhotoKind `json:"kind"`
		AlbumId   uint             `json:"albumId,omitempty"`
	}

	photoUuid, err := uuid.NewV7()
	if err != nil {
		return models.Photo{}, err
	}

	objectName := fmt.Sprintf("%d/%s", userId, photoUuid.String())

	_, err = r.Storage.PutObject(
		context.Background(),
		"photos",
		objectName,
		file,
		fileSize,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return models.Photo{}, err
	}

	collection := r.DB.Collection("photos")

	photoData := models.Photo{
		ID:          photoUuid.String(),
		UserID:      userId,
		Kind:        kind,
		FileName:    fileName,
		FileSize:    fileSize,
		ContentType: contentType,
		UploadedAt:  time.Now(),
		Metadata:    nil,
	}

	_, err = collection.InsertOne(context.Background(), photoData)
	if err != nil {
		return models.Photo{}, err
	}

	msg := BrokerMessage{
		ID:        photoData.ID,
		Timestamp: photoData.UploadedAt.String(),
		Kind:      photoData.Kind,
	}
	if albumId != 0 {
		msg.AlbumId = albumId
	}

	msgString, err := json.Marshal(msg)

	message := sarama.ProducerMessage{
		Topic: utils.TopicNewPhoto,
		Value: sarama.ByteEncoder(msgString),
	}
	_, _, err = r.Producer.SendMessage(&message)
	if err != nil {
		return models.Photo{}, err
	}

	return photoData, nil
}

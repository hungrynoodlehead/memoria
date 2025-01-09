package photo_repository

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/hungrynoodlehead/memoria/services/album-service/models"
	"github.com/hungrynoodlehead/memoria/services/album-service/utils"
)

func (r *PhotoRepository) DeletePhoto(photo models.Photo) error {
	err := r.DB.Delete(&photo).Error
	if err != nil {
		return err
	}

	type message struct {
		PhotoID string `json:"photo_id"`
	}
	var msg message
	msg.PhotoID = photo.UUID

	msgStr, err := json.Marshal(msg)

	_, _, err = r.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: utils.TopicRemovedPhotos,
		Value: sarama.ByteEncoder(msgStr),
	})
	if err != nil {
		return err
	}

	return nil
}

package consumer_handler

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
)

type removedPhotoMessage struct {
	PhotoID string `json:"photo_id"`
}

func (h *ConsumerGroupHandler) NewRemovedPhotoHandler(msg *sarama.ConsumerMessage) error {
	var message removedPhotoMessage
	err := json.Unmarshal(msg.Value, &message)
	if err != nil {
		return err
	}

	photoUuid, err := uuid.Parse(message.PhotoID)
	if err != nil {
		return err
	}

	err = h.PhotoRepository.DeletePhoto(photoUuid)
	if err != nil {
		return err
	}
	return nil
}

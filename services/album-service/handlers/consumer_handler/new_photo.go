package consumer_handler

import (
	"encoding/json"
	"errors"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/hungrynoodlehead/memoria/services/album-service/models"
	"gorm.io/gorm"
)

type BrokerMessage struct {
	ID        string           `json:"id"`
	Timestamp string           `json:"timestamp"`
	Kind      models.PhotoKind `json:"kind"`
	AlbumId   uint64           `json:"albumId,omitempty"`
}

func (h *ConsumerGroupHandler) NewPhotoHandler(m *sarama.ConsumerMessage) error {
	var message BrokerMessage
	err := json.Unmarshal(m.Value, &message)
	if err != nil {
		return err
	}

	uuid, err := uuid.Parse(message.ID)
	if err != nil {
		return err
	}

	photo, err := h.PhotoRepository.Create(uuid, message.Kind)
	if err != nil {
		return err
	}

	// TODO: Add photo to album after receiving id if present
	if message.AlbumId != 0 {
		album, err := h.AlbumRepository.GetByID(message.AlbumId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}

		_, err = h.AlbumRepository.AddToAlbum(*album, photo)
		if err != nil {
			return err
		}
	}

	return nil
}

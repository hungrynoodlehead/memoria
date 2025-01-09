package consumer_handler

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/hungrynoodlehead/memoria/services/storage-service/repositories/photo_repository"
	"github.com/hungrynoodlehead/memoria/services/storage-service/utils"
	"log"
)

type ConsumerGroupHandler struct {
	Config *utils.Config
	DB     *utils.DB

	PhotoRepository *photo_repository.PhotoRepository

	Handlers map[string]func(*sarama.ConsumerMessage) error
}

func NewConsumerGroupHandler(config *utils.Config, db *utils.DB, photoRepository *photo_repository.PhotoRepository) *ConsumerGroupHandler {
	h := ConsumerGroupHandler{
		Config:          config,
		DB:              db,
		PhotoRepository: photoRepository,
	}

	h.Handlers = map[string]func(*sarama.ConsumerMessage) error{
		"removed-photos": h.NewRemovedPhotoHandler,
	}

	return &h
}

func (h *ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	log.Println("Consumer group setup complete")
	return nil
}

func (h *ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	log.Println("Consumer group cleanup complete")
	return nil
}

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		handler, exists := h.Handlers[msg.Topic]
		if !exists {
			fmt.Printf("Consumer topic %s not found in handler\n", msg.Topic)
			continue
		}

		if err := handler(msg); err != nil {
			fmt.Printf("Consumer topic %s error %s\n", msg.Topic, err)
		} else {
			session.MarkMessage(msg, "")
		}
	}
	return nil
}

func (h *ConsumerGroupHandler) GetTopics() []string {
	var topics []string
	for key := range h.Handlers {
		topics = append(topics, key)
	}
	return topics
}

func StartConsumer(config *utils.Config, handler *ConsumerGroupHandler) error {
	saramaConfig := sarama.NewConfig()
	consumer, err := sarama.NewConsumerGroup([]string{config.GetKafkaEndpoint()}, "album-consumers", saramaConfig)
	if err != nil {
		return err
	}
	go func() {
		for {
			err = consumer.Consume(context.Background(), handler.GetTopics(), handler)
			if err != nil {
				return
			}
		}
	}()
	return nil
}

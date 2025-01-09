package utils

import "github.com/IBM/sarama"

type MessageProducer struct {
	sarama.SyncProducer
	Config *Config
}

func NewMessageProducer(config *Config) (*MessageProducer, error) {
	brokerEndpoint := config.GetKafkaAddresses()

	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Retry.Max = 5
	saramaConfig.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokerEndpoint, saramaConfig)
	if err != nil {
		return nil, err
	}

	return &MessageProducer{
		SyncProducer: producer,
		Config:       config,
	}, nil
}

const (
	TopicNewPhoto      = "new-photos"
	TopicRemovedPhotos = "removed-photos"
)

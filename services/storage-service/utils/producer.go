package utils

import "github.com/IBM/sarama"

type BrokerProducer struct {
	sarama.SyncProducer
	Logger *Logger
	Config *Config
}

func NewBrokerProducer(logger *Logger, config *Config) (*BrokerProducer, error) {
	brokerEndpoint := []string{config.GetKafkaEndpoint()}

	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Retry.Max = 5
	saramaConfig.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokerEndpoint, saramaConfig)
	if err != nil {
		return nil, err
	}

	return &BrokerProducer{
		SyncProducer: producer,
		Logger:       logger,
		Config:       config,
	}, nil
}

const (
	TopicNewPhoto = "new-photos"
)

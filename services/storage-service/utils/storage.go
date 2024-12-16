package utils

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	*minio.Client
	Config *Config
	Logger *Logger
}

func NewStorage(config *Config, logger *Logger) (*Storage, error) {
	var storage Storage
	storage.Config = config
	storage.Logger = logger

	client, err := minio.New(storage.Config.GetStorageEndpoint(), &minio.Options{
		Creds: credentials.NewStaticV4(storage.Config.GetStorageAccessKey(), storage.Config.GetStorageSecretKey(), ""),
	})
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully connected to Minio")

	bucketExists, err := client.BucketExists(context.Background(), "photos")
	if err != nil {
		return nil, err
	}
	if !bucketExists {
		storage.Logger.Info("Creating storage bucket")
		err := client.MakeBucket(context.Background(), "photos", minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	storage.Client = client

	return &storage, nil
}

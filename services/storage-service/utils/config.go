package utils

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

var (
	ErrDBConnectionStringNotFound = errors.New("database connection string not found")
)

func NewConfig() (*Config, error) {
	v := viper.New()

	v.SetEnvPrefix("CONFIG")

	v.SetDefault("LISTEN_PORT", "8080")
	v.SetDefault("LOG_LEVEL", "INFO")

	v.AutomaticEnv()

	return &Config{Viper: v}, nil
}

func (c *Config) GetListenPort() string {
	return c.GetString("LISTEN_PORT")
}

func (c *Config) GetDBString() string {
	return c.GetString("MONGODB_CONNECTION_STRING")
}

func (c *Config) GetLogLevel() string {
	return c.GetString("LOG_LEVEL")
}

func (c *Config) GetStorageEndpoint() string {
	return c.GetString("MINIO_ENDPOINT")
}

func (c *Config) GetStorageAccessKey() string {
	return c.GetString("MINIO_ACCESS_KEY")
}

func (c *Config) GetStorageSecretKey() string {
	return c.GetString("MINIO_SECRET_KEY")
}

func (c *Config) GetKafkaEndpoint() string { return c.GetString("KAFKA_ENDPOINT") }

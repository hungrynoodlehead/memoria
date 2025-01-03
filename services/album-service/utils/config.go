package utils

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	*viper.Viper
}

func NewConfig() (*Config, error) {
	v := viper.New()

	v.SetEnvPrefix("CONFIG")

	v.SetDefault("LISTEN_PORT", "8080")
	v.SetDefault("LOG_LEVEL", "INFO")

	v.AutomaticEnv()

	return &Config{Viper: v}, nil
}

func (c *Config) GetConnectonString() string {
	return c.GetString("POSTGRES_CONNECTION_STRING")
}

func (c *Config) GetKafkaAddresses() []string {
	p := c.GetString("KAFKA_ADDRESSES")
	return strings.Split(p, ",")
}

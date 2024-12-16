package utils

import (
	"errors"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	*viper.Viper
}

func NewConfig() (*Config, error) {
	v := *viper.New()

	v.SetConfigType("yaml")
	v.SetConfigFile("config")
	v.AddConfigPath(".")

	v.SetEnvPrefix("CONFIG")

	err := v.BindEnv("jwt_token_secret")
	if err != nil {
		return nil, errors.New("jwt_token_secret is not set")
	}
	v.SetDefault("log_level", "Info")
	v.SetDefault("jwt_access_token_duration", int64(time.Hour*24))
	v.SetDefault("jwt_refresh_token_duration", int64(time.Hour*24*30))

	err = v.ReadInConfig()
	if err != nil {
		if errors.Is(err, viper.ConfigFileNotFoundError{}) {
			err := v.WriteConfig()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	err = v.WriteConfig()
	if err != nil {
		return nil, err
	}

	return &Config{Viper: &v}, nil
}

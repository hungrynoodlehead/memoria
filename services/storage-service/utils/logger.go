package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	*logrus.Logger
	Config *Config
}

func NewLogger(config *Config) (*Logger, error) {
	var logger Logger
	logger.Config = config

	log := logrus.New()

	log.SetOutput(os.Stdout)
	
	level, err := logrus.ParseLevel(logger.Config.GetLogLevel())
	if err != nil {
		return nil, err
	} else {
		log.SetLevel(level)
	}

	logger.Logger = log
	return &logger, nil
}

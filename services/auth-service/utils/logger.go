package utils

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
	Config *Config
}

func InitLogger(config *Config) *Logger {
	log := &Logger{Logger: logrus.New(), Config: config}

	level, err := logrus.ParseLevel(log.Config.GetString("log_level"))
	if err != nil {
		log.Fatal("Unknown logging level in config")
	} else {
		log.SetLevel(level)
	}

	return log
}

package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type LogConfigInterface interface {
	GetSentryDns() string
	GetLogLevel() log.Level
}

type LogConfig struct {
	SentryDns string `json:"sentryDNS"`
	LogLevel  string `json:"logLevel"`
}

func (config LogConfig) GetSentryDns() string {
	return config.SentryDns
}

func (config LogConfig) GetLogLevel() (level log.Level) {
	level, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		level = log.ErrorLevel
	}
	return
}

func NewLogConfigFromEnvironment() (logConfig LogConfigInterface) {
	logConfig = &LogConfig{
		os.Getenv("SENTRY_DNS"),
		os.Getenv("LOG_LEVEL"),
	}
	return
}

package helper

import (
	"os"
	"time"

	"github.com/evalphobia/logrus_sentry"
	log "github.com/sirupsen/logrus"
)

func ConfigureLogger(level log.Level) {
	log.SetLevel(level)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func ConfigureSentry(sentryDns string, tags map[string]string) {
	logLevels := []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
		log.WarnLevel,
	}

	hook, err := logrus_sentry.NewWithTagsSentryHook(sentryDns, tags, logLevels)

	hook.Timeout = 5 * time.Second
	hook.StacktraceConfiguration.Enable = true

	if err == nil {
		log.AddHook(hook)
	}
}

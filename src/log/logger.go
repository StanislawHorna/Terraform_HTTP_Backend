package log

import (
	"os"

	"terraform_http_backend/src"

	"github.com/rs/zerolog"
)

const (
	defaultLogLevel    = zerolog.InfoLevel
	callerFramesToSkip = 3
)

var (
	logger zerolog.Logger
)

func Init() {
	conf := src.GetConfig()
	logLevel, err := zerolog.ParseLevel(conf.LogLevel)
	if err != nil {
		logLevel = defaultLogLevel
	}

	if conf.Loki.LokiURL != "" {
		labels := LokiLabels{
			env:        conf.Loki.Environment,
			appName:    conf.Loki.AppName,
			appVersion: conf.Loki.AppVersion,
			goVersion:  conf.Loki.GoVersion,
		}
		loki := newLokiWriter(conf.Loki.LokiURL, labels)
		logger = zerolog.New(zerolog.MultiLevelWriter(os.Stdout, loki)).
			Level(logLevel).
			With().
			Timestamp().
			CallerWithSkipFrameCount(callerFramesToSkip).
			Logger()
		return
	}

	logger = zerolog.New(os.Stdout).
		Level(logLevel).
		With().
		Timestamp().
		CallerWithSkipFrameCount(callerFramesToSkip).
		Logger()
}

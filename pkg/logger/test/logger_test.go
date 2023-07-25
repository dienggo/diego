package test

import (
	"github.com/dienggo/diego/config"
	"github.com/dienggo/diego/pkg/environment"
	"github.com/dienggo/diego/pkg/logger"
	"testing"
	"time"
)

func TestLoggerSimple(t *testing.T) {
	logger.Error("Testing error", logger.SetField("key", "val"))
}

func TestLoggerWithJsonFormatter(t *testing.T) {
	logger.SetJSONFormatter()
	logger.Error("Testing error", logger.SetField("key", "val"))
}

func TestLoggerSentryEnabled(t *testing.T) {
	environment.Load()
	logConfig := logger.Config{
		URL:         config.Sentry().DSN,
		Debug:       config.App().Debug,
		Environment: config.App().Env,
	}
	hook, err := logger.NewSentryHook(logConfig)
	if err != nil {
		logger.Error("error initiate sentry hook on logrus",
			logger.SetField("url", logConfig.URL),
			logger.SetField("env", logConfig.Environment),
			logger.SetField("debug", logConfig.Debug),
		)
		return
	}

	logger.Setup(logConfig, hook)

	// hit log
	logger.Error("Okee again")

	time.Sleep(2 * time.Second)
}

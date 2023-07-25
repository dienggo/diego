package provider

import (
	"github.com/dienggo/diego/config"
	"github.com/dienggo/diego/pkg/logger"
)

type pLogger struct{}

func (p pLogger) Provide() {
	logConfig := logger.Config{
		URL:         config.Sentry().DSN,
		Debug:       config.App().Debug,
		Environment: config.App().Env,
	}
	hook, err := logger.NewSentryHook(logConfig)
	if err != nil {
		logger.Error("error initiate sentry hook",
			logger.SetField("url", logConfig.URL),
			logger.SetField("env", logConfig.Environment),
			logger.SetField("debug", logConfig.Debug),
		)
		return
	}
	logger.Setup(logConfig, hook)
}

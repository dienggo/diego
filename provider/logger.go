package provider

import (
	"github.com/dienggo/diego/config"
	"github.com/dienggo/diego/pkg/logger"
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
)

type pLogger struct{}

func (p pLogger) Provide() {
	if config.App().Debug {
		log.SetLevel(log.DebugLevel)
	}
	provideSentry()
}

func provideSentry() {
	if config.Sentry().DSN != "" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:         config.Sentry().DSN,
			Environment: config.App().Env,
		})
		if err != nil {
			log.Fatalf("Sentry initialization failed: %v", err)
		}

		// Create a custom Logrus hook that sends logs to Sentry
		sentryHook := &logger.SentryHook{
			AllowedLevels: []log.Level{log.ErrorLevel, log.FatalLevel, log.WarnLevel},
		}
		log.AddHook(sentryHook)
	}
}

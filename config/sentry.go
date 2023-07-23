package config

import (
	"os"
)

var sentryConfig sentry
var sentryIntance bool = false

type sentry struct {
	DSN string
}

func Sentry() sentry {
	if !sentryIntance {
		sentryConfig = sentry{
			DSN: os.Getenv("SENTRY_DSN"),
		}

		sentryIntance = true
	}
	return sentryConfig
}

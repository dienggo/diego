// Package logger
package logger

import (
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"reflect"
)

// severityMap is a mapping of logrus log level to sentry log level.
var severityMap = map[log.Level]sentry.Level{
	log.DebugLevel: sentry.LevelDebug,
	log.InfoLevel:  sentry.LevelInfo,
	log.WarnLevel:  sentry.LevelWarning,
	log.ErrorLevel: sentry.LevelError,
	log.FatalLevel: sentry.LevelFatal,
	log.PanicLevel: sentry.LevelFatal,
}

var levelingMap = map[string][]log.Level{
	"trace": []log.Level{
		log.TraceLevel,
		log.DebugLevel,
		log.InfoLevel,
		log.WarnLevel,
		log.ErrorLevel,
		log.FatalLevel,
		log.PanicLevel,
	},
	"debug": []log.Level{
		log.DebugLevel,
		log.InfoLevel,
		log.WarnLevel,
		log.ErrorLevel,
		log.FatalLevel,
		log.PanicLevel,
	},
	"info": []log.Level{
		log.InfoLevel,
		log.WarnLevel,
		log.ErrorLevel,
		log.FatalLevel,
		log.PanicLevel,
	},
	"warn": []log.Level{
		log.WarnLevel,
		log.ErrorLevel,
		log.FatalLevel,
		log.PanicLevel,
	},
	"error": []log.Level{
		log.ErrorLevel,
		log.FatalLevel,
		log.PanicLevel,
	},
}

// SentryHook implements logrus.Hook to send errors to sentry.
type SentryHook struct {
	client *sentry.Client
	levels []log.Level
}

// SentryEventIdentityModifier is a sentry event modifier that simply passes
// through the event.
type SentryEventIdentityModifier struct{}

// ApplyToEvent simply returns the event (ignoring the hint).
func (m *SentryEventIdentityModifier) ApplyToEvent(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
	return event
}

var sentryModifier = &SentryEventIdentityModifier{}

// NewSentryHook creates a sentry hook for logrus given a sentry dsn
func NewSentryHook(cfg Config) (*SentryHook, error) {
	client, err := sentry.NewClient(sentry.ClientOptions{
		Dsn:              cfg.URL,
		Environment:      Environment(cfg.Environment),
		TracesSampleRate: 1.0,
	})

	if err != nil {
		Fatal("Unable to initialize Sentry", SetField("error", err))
		return nil, err
	}

	lvl, ok := levelingMap[cfg.Level]
	if !ok {
		lvl = levelingMap["info"]
	}

	return &SentryHook{
		client: client,
		levels: lvl,
	}, nil
}

// Levels returns the levels this hook is enabled for. This is a part
// of logrus.Hook.
func (h *SentryHook) Levels() []log.Level {
	return h.levels
}

// Fire is an event handler for logrus. This is a part of logrus.Hook.
func (h *SentryHook) Fire(e *log.Entry) error {
	event := sentry.NewEvent()
	event.Message = e.Message
	event.Level = severityMap[e.Level]
	event.Timestamp = e.Time

	var err error

	for k, v := range e.Data {
		if k == log.ErrorKey {
			err = v.(error)
			continue
		}
		event.Extra[k] = v
	}

	if err != nil {
		stacktrace := sentry.ExtractStacktrace(err)
		event.Exception = []sentry.Exception{{
			Value:      err.Error(),
			Type:       reflect.TypeOf(err).String(),
			Stacktrace: stacktrace,
		}}
	}

	h.client.CaptureEvent(event, nil, sentryModifier)
	return nil
}

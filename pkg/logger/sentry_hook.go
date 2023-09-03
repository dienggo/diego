package logger

import (
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

// severityMap is a mapping of logrus log level to sentry log level.
var severityMap = map[logrus.Level]sentry.Level{
	logrus.DebugLevel: sentry.LevelDebug,
	logrus.InfoLevel:  sentry.LevelInfo,
	logrus.WarnLevel:  sentry.LevelWarning,
	logrus.ErrorLevel: sentry.LevelError,
	logrus.FatalLevel: sentry.LevelFatal,
	logrus.PanicLevel: sentry.LevelFatal,
}

// SentryHook is a custom Logrus hook for sending logs to Sentry.
type SentryHook struct {
	AllowedLevels []logrus.Level
}

// Fire is called when a log event is fired.
func (hook *SentryHook) Fire(entry *logrus.Entry) error {
	// Extract the log message and extra fields
	message := entry.Message
	extraFields := entry.Data

	// Create a Sentry event and send it
	sentryEvent := sentry.NewEvent()
	sentryEvent.Message = message
	sentryEvent.Extra = extraFields
	sentryEvent.Level = severityMap[entry.Level]

	sentry.CaptureEvent(sentryEvent)
	return nil
}

// Levels returns the log levels that this hook should be triggered for.
func (hook *SentryHook) Levels() []logrus.Level {
	return hook.AllowedLevels
}

package logger

import (
	"github.com/sirupsen/logrus"
)

// AppLoggerHook is a custom Logrus hook that sends log messages to a custom destination.
type AppLoggerHook struct{}

// Fire is called when a log event is fired.
func (hook *AppLoggerHook) Fire(entry *logrus.Entry) error {
	return nil
}

// Levels returns the log levels that this hook should be triggered for.
func (hook *AppLoggerHook) Levels() []logrus.Level {
	// You can specify the log levels for which this hook should be triggered.
	// For example, log only Error and Fatal messages:
	return logrus.AllLevels
}

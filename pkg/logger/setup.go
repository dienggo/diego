// Package logger
package logger

import (
	"github.com/sirupsen/logrus"
)

func SetJSONFormatter() {
	logrus.SetFormatter(&Formatter{
        ChildFormatter: &logrus.JSONFormatter{
            FieldMap: logrus.FieldMap{
                logrus.FieldKeyMsg: MessageKey,
            },
        },
        Line:         true,
        Package:      false,
        File:         true,
        BaseNameOnly: false,
    })
}

func Setup(cfg Config, h logrus.Hook) {

	if h != nil {
		logrus.AddHook(h)
	}

	if cfg.Debug {
		logrus.SetLevel(logrus.TraceLevel)
		return
	}

	logrus.SetLevel(logrus.InfoLevel)
}

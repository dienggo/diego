package logger

import "github.com/sirupsen/logrus"

func SetJSONFormatter() {
	logrus.SetFormatter(&Formatter{
		ChildFormatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg: MessageKey,
			},
			PrettyPrint: true,
		},
		Line:         true,
		Package:      false,
		File:         true,
		BaseNameOnly: false,
	})
}

// Initiate : to initiate logger configuration
func Initiate() {
	SetJSONFormatter()
	logrus.AddHook(&AppLoggerHook{})
}

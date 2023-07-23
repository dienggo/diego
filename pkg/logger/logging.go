// Package logger
package logger

import (
	"context"
	"fmt"
	"github.com/daewu14/golang-base/config"
	"github.com/daewu14/golang-base/pkg/helper"
	"github.com/sirupsen/logrus"
)

type logField map[string]interface{}

type Field struct {
	Key   string
	Value interface{}
}

type FieldFunc func(key string, value interface{}) *Field

func SetField(k string, v interface{}) Field {
	return Field{
		Key:   k,
		Value: v,
	}
}

func Any(k string, v interface{}) Field {
	return Field{
		Key:   k,
		Value: v,
	}
}

func EventName(v interface{}) Field {
	return Field{
		Key:   "name",
		Value: v,
	}
}

func EventId(v interface{}) Field {
	return Field{
		Key:   "id",
		Value: v,
	}
}

func SetMessageFormat(format string, args ...interface{}) interface{} {
	return fmt.Sprintf(format, args...)
}

func extract(args ...Field) map[string]interface{} {
	if len(args) == 0 {
		return nil
	}

	data := map[string]interface{}{}
	for _, fl := range args {
		data[fl.Key] = fl.Value
	}
	return data
}

func Error(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Error(arg)
}

func Local(arg interface{}, fl ...Field) {
	if Environment(config.App().Env) == "loc" {
		logrus.WithFields(map[string]interface{}{
			"event": extract(fl...),
		}).Info(arg)
	}
}

func Info(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Info(arg)
}

func Debug(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Debug(arg)
}

func Fatal(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Fatal(arg)
}

func Warn(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Warn(arg)
}

func Trace(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Trace(arg)
}

func AccessLog(arg interface{}, fl ...Field) {
	logrus.WithFields(extract(fl...)).Info(arg)
}

func InfoWithContext(ctx context.Context, arg interface{}, fl ...Field) {
	logrus.WithFields(extractContext(ctx.Value("access"), map[string]interface{}{
		"event": extract(fl...),
	})).WithContext(ctx).Info(arg)
}

func extractContext(i interface{}, logField map[string]interface{}) map[string]interface{} {
	if helper.IsSameType(i, logField) {
		x := i.(map[string]interface{})
		for k, v := range x {
			logField[k] = v
		}
	}
	return logField
}

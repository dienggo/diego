// Package logger
package logger

import "strings"

var envs = map[string]string{
	"production":  "prod",
	"staging":     "stg",
	"development": "dev",
	"prod":        "prod",
	"stg":         "stg",
	"dev":         "dev",
	"local":       "loc",
	"loc":         "loc",
	"prd":         "prod",
	"green":       "green",
	"blue":        "blue",
}

func Environment(env string) string {
	v, ok := envs[strings.ToLower(strings.Trim(env, " "))]

	if !ok {
		return ""
	}

	return v
}

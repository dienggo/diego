// Package logger
package logger

type Config struct {
	URL         string `json:"url"`
	Debug       bool   `json:"debug"`
	Environment string `json:"environment"`
	Level       string `json:"level"`
}

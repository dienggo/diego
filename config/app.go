package config

import "os"

var a app
var appInstance bool = false

const envProductionName = "production"

type app struct {
	Name     string
	Port     string
	Key      string
	Env      string
	TimeZone string
}

// IsDevelopment is representation of app environment are development
func (a app) IsDevelopment() bool {
	return a.Env != envProductionName
}

// IsProduction is representation of app environment are production
func (a app) IsProduction() bool {
	return a.Env == envProductionName
}

// App is configuration of application built
func App() app {
	if appInstance == false {
		a = app{
			Name:     os.Getenv("APP_NAME"),
			Port:     os.Getenv("APP_PORT"),
			Key:      os.Getenv("APP_KEY"),
			Env:      os.Getenv("APP_ENV"),
			TimeZone: os.Getenv("TIME_ZONE"),
		}
		appInstance = true
	}
	return a
}

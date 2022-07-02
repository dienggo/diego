package config

import (
	"os"
)

type database struct {
	Name         string
	Username     string
	Password     string
	Host         string
	Port         string
	UseTimestamp string
}

// Database / Getter database configuration
func Database() database {
	db := database{}
	db.Name = os.Getenv("DB_NAME")
	db.Username = os.Getenv("DB_USERNAME")
	db.Password = os.Getenv("DB_PASSWORD")
	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.UseTimestamp = os.Getenv("DB_USE_TIMESTAMP")
	return  db
}
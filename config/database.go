package config

import (
	"os"
)

// Make this to singleton struct
var isInstance bool = false
var name string
var username string
var password string
var host string
var port string
var useTimestamp string

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
	if !isInstance { instance() }
	db := database{
		Name:         name,
		Username:     username,
		Password:     password,
		Host:         host,
		Port:         port,
		UseTimestamp: useTimestamp,
	}
	return db
}

// instance : for instantiation value store to singleton variable
func instance() {
	name = os.Getenv("DB_NAME")
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	useTimestamp = os.Getenv("DB_USE_TIMESTAMP")
	isInstance = true
}
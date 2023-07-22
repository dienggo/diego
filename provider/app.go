package provider

import (
	"github.com/daewu14/golang-base/config"
	"github.com/daewu14/golang-base/pkg/database"
	"github.com/daewu14/golang-base/pkg/environment"
	"github.com/daewu14/golang-base/pkg/router"
	"time"
)

type App struct{}

// Start all application resource
func (app App) Start() {
	// Load environment
	environment.Load()

	println("\n------------------------------------------------------------")
	println(config.App().Name + " app starting")
	println("------------------------------------------------------------\n")

	// Load database
	defer database.Close()
	database.Open()

	// Set time zone application
	time.LoadLocation(config.App().TimeZone)

	// Run route configuration
	router.Run()
}

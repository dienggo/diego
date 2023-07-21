package provider

import (
	"go_base_project/config"
	"go_base_project/pkg/database"
	"go_base_project/pkg/environment"
	"go_base_project/pkg/router"
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

package provider

import (
	"go_base_project/config"
	"go_base_project/provider/database"
	"go_base_project/provider/p_routes"
	"time"
)

type App struct{}

// Start all application resource
func (app App) Start() {
	// Load environment
	LoadEnv()

	println("\n------------------------------------------------------------")
	println(config.App().Name + " app starting")
	println("------------------------------------------------------------\n")

	// Load database
	defer database.Close()
	database.Open()

	// Set time zone application
	time.LoadLocation(config.App().TimeZone)

	// Run route configuration
	p_routes.Run()
}

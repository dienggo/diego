package provider

import (
	"github.com/daewu14/golang-base/config"
	"github.com/daewu14/golang-base/pkg/database"
	"github.com/daewu14/golang-base/pkg/environment"
	"github.com/daewu14/golang-base/pkg/router"
	"time"
)

type IProvider interface {
	Provide()
}

type App struct{}

// Start all application resource
func (app App) Start() {
	// Load environment
	environment.Load()
	registry()

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

func registry() {
	providerRegistry{providers: []IProvider{
		pLogger{},
	}}.Provide()
}

type providerRegistry struct {
	providers []IProvider
}

func (p providerRegistry) Provide() {
	for _, provider := range p.providers {
		provider.Provide()
	}
}

package provider

import (
	"github.com/dienggo/diego/config"
	"github.com/dienggo/diego/pkg/database"
	"github.com/dienggo/diego/pkg/environment"
	"github.com/dienggo/diego/pkg/router"
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

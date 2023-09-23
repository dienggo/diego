package provider

import (
	"github.com/dienggo/diego/config"
	"github.com/dienggo/diego/pkg/database"
	"github.com/dienggo/diego/pkg/router"
	log "github.com/sirupsen/logrus"
	"time"
)

type IProvider interface {
	Provide()
}

type App struct{}

// Start all application resource
func (app App) Start() {
	registry()

	println("\n------------------------------------------------------------")
	println(config.App().Name + " app starting")
	println("------------------------------------------------------------\n")

	// Set time zone application
	_, err := time.LoadLocation(config.App().TimeZone)
	if err != nil {
		log.Panicf("Fail when set time zone", err.Error())
	}

	// Load database
	database.Open()

	// Run route configuration
	router.New().OnDone(database.Close).Run()
}

// registry is method to execute all provider on this app
func registry() {
	providerRegistry{providers: []IProvider{
		pLogger{},
		appTask{},
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

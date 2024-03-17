package provider

import (
	"github.com/dienggo/diego/config"
	log "github.com/sirupsen/logrus"
	"time"
)

type IProvider interface {
	Provide()
}

type App struct{}

// Start all application resource
func (app App) Start() {
	println("\n------------------------------------------------------------")
	println(config.App().Name + " app starting")
	println("------------------------------------------------------------\n")

	// Set time zone application
	_, err := time.LoadLocation(config.App().TimeZone)
	if err != nil {
		log.Panicf("Fail when set time zone", err.Error())
	}

	// execute all provider on this app
	providerRegistry{providers: []IProvider{
		NewLoggerProvider(),
		NewAppTaskProvider(),
		NewRouteProvider(),
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

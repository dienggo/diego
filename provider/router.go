package provider

import (
	"github.com/dienggo/diego/pkg/database"
	"github.com/dienggo/diego/pkg/router"
	"github.com/dienggo/diego/routes"
)

func NewRouteProvider() IProvider {
	return &route{}
}

type route struct{}

func (route) Provide() {
	// Load database
	database.Open()

	// Run route configuration
	router.New().OnDone(database.Close).Run([]router.IRoute{
		routes.NewApi(),
	})
}

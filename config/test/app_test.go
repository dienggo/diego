package test

import (
	"github.com/dienggo/diego/config"
	"github.com/dienggo/diego/pkg/environment"
	"testing"
)

func TestAppConfig(t *testing.T) {
	environment.Load()
	app := config.App()

	println("APP NAME", app.Name)
	println("APP Port", app.Port)
	println("APP Key", app.Key)
	println("APP Env", app.Env)
	println("APP TimeZone", app.TimeZone)
	println("APP Is Development", app.IsDevelopment())
}

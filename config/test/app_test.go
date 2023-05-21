package test

import (
	"go_base_project/config"
	"go_base_project/provider"
	"testing"
)

func TestAppConfig(t *testing.T) {
	provider.LoadEnv()
	app := config.App()

	println("APP NAME", app.Name)
	println("APP Port", app.Port)
	println("APP Key", app.Key)
	println("APP Env", app.Env)
	println("APP TimeZone", app.TimeZone)
	println("APP Is Development", app.IsDevelopment())
}

package test

import (
	"go_base_project/config"
	"go_base_project/pkg/environment"
	"testing"
)

func TestDatabase(t *testing.T) {
	environment.Load()
	config.Database()
}

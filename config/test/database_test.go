package test

import (
	"go_base_project/config"
	"go_base_project/provider"
	"testing"
)

func TestDatabase(t *testing.T) {
	provider.LoadEnv()
	config.Database()
}

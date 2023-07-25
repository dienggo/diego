package test

import (
	"github.com/dienggo/diego/config"
	"github.com/dienggo/diego/pkg/environment"
	"testing"
)

func TestDatabase(t *testing.T) {
	environment.Load()
	config.Database()
}

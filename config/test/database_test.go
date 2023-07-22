package test

import (
	"github.com/daewu14/golang-base/config"
	"github.com/daewu14/golang-base/pkg/environment"
	"testing"
)

func TestDatabase(t *testing.T) {
	environment.Load()
	config.Database()
}

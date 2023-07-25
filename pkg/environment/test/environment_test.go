package test

import (
	"github.com/dienggo/diego/pkg/environment"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	environment.Load()
}

package test

import (
	"go_base_project/pkg/environment"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	environment.Load()
}

package test

import (
	"github.com/dienggo/diego/pkg/database"
	"github.com/dienggo/diego/pkg/environment"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	environment.Load()
	defer database.Close()
	db := database.Open()
	println("database is not nil", db.Main() != nil)
}

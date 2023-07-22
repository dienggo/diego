package test

import (
	"go_base_project/pkg/database"
	"go_base_project/pkg/environment"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	environment.Load()
	defer database.Close()
	db := database.Open()
	println("database is not nil", db.Main() != nil)
}

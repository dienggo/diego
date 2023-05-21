package test

import (
	"go_base_project/bootstrap/database"
	"go_base_project/provider"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	provider.LoadEnv()
	defer database.Close()
	db := database.Open()
	println("database is not nil", db.Main() != nil)
}

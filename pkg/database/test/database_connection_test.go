package test

import (
	"github.com/daewu14/golang-base/pkg/database"
	"github.com/daewu14/golang-base/pkg/environment"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	environment.Load()
	defer database.Close()
	db := database.Open()
	println("database is not nil", db.Main() != nil)
}

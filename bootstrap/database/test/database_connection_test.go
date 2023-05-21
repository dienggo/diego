package test

import (
	"go_base_project/bootstrap/database"
	"go_base_project/provider"
	"gorm.io/gorm"
	"testing"
)

type Member struct {
	gorm.Model
	Name          string
	Email         string
	Level         string
	Status        string
	Settings      string
	MemberLevelId int
	Verified      int
	ReferalId     int
}

func TestDatabaseConnection(t *testing.T) {
	provider.LoadEnv()
	defer database.Close()
	db := database.Open()
	println("database is not null", db.Main() != nil)
}

package test

import (
	"github.com/stretchr/testify/assert"
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
	mm := Member{}
	defer database.Close()
	database.Open().Main().Where("id = ?", 2).First(&mm)

	assert.IsTypef(t, Member{}, mm, "Okee")
}

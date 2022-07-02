package base

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"testing"
)

// Example Gorm Model
type Member struct {
	gorm.Model
	Name          string
	Level         string
	Status        string
	Settings      string
	MemberLevelId int
	Verified      int
	ReferalId     int
}

func TestDatabaseQuery(t *testing.T) {

	godotenv.Load("../../.env")

	member := Member{}

	memberId := 1
	OpenDB().Gorm().Where("id = ?", memberId).First(&member)

	println("member :",member.Name)
}
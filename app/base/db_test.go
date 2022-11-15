package base

import (
	"gorm.io/gorm"
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

 // NOTE: don't call test actually to database 
// func TestDatabaseQuery(t *testing.T) {

// 	godotenv.Load("../../.env")

// 	member := Member{}

// 	memberId := 1
// 	OpenDB().Gorm().Where("id = ?", memberId).First(&member)

// 	println("member :", member.Name)
// }

package models

import "gorm.io/gorm"

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

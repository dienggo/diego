package repositories

import (
	"go_base_project/app/base"
	"go_base_project/app/models"
)

type MemberRepoInterface interface {
	Find(id int) models.Member
}

func Member() member {return member{}}

type member struct{}

func (mr member) Find(id any) models.Member {
	m := models.Member{}
	db := base.OpenDB()
	db.Gorm().Where("id = ?", id).First(&m)
	db.Close()
	return m
}

package repositories

import (
	"go_base_project/app/base"
	"go_base_project/app/models"
)

type UserRepositoriesInterface interface {
	Find(Id any) models.User
	Create(name any) models.User
}

type UserRepository struct{}

func (s UserRepository) Find(Id any) models.User {

	userModel := models.User{}
	sqlStatement := "SELECT id, username FROM users WHERE id=?"

	base.OpenDB().DB().QueryRow(sqlStatement, Id).Scan(&userModel.Id, &userModel.Name)

	return userModel
}

func (s UserRepository) Create(name any) models.User {

	sqlStatement := "INSERT INTO users (username) VALUES (?)"

	insert, _ := base.OpenDB().DB().Exec(sqlStatement, name)
	userId, _ := insert.LastInsertId()
	userObject := s.Find(userId)

	return userObject
}

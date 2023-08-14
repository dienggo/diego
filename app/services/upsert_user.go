package services

import (
	"github.com/dienggo/diego/app/dto_request"
	"github.com/dienggo/diego/app/models"
	"github.com/dienggo/diego/app/repositories"
	"github.com/dienggo/diego/pkg/app"
	"github.com/dienggo/diego/pkg/validates"
)

// NewUpsertUser : to simplify service call
func NewUpsertUser(user dto_request.User) UpsertUser {
	return UpsertUser{RepoUser: repositories.User{}, DtoReqUser: user}
}

type UpsertUser struct {
	app.Service
	RepoUser   repositories.IUser
	DtoReqUser dto_request.User
}

func (s UpsertUser) Do() app.ServiceResponse {

	// validation the request
	err := validates.ValidateStructFormatted(&s.DtoReqUser)
	if err != nil {
		return s.Error(err, nil)
	}

	// bind request into model
	user := &models.User{
		Model:    app.Model{ID: uint(s.DtoReqUser.ID)},
		Name:     s.DtoReqUser.Name,
		Email:    s.DtoReqUser.Email,
		Password: s.DtoReqUser.Password,
	}

	// do query
	err = s.RepoUser.Upsert(user)
	if err != nil {
		return s.Error(err, nil)
	}

	return s.Success(user)
}

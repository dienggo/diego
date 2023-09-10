package userCase

import (
	"github.com/dienggo/diego/app/dto_request"
	"github.com/dienggo/diego/app/dto_response"
	"github.com/dienggo/diego/app/models"
	"github.com/dienggo/diego/app/repositories"
	"github.com/dienggo/diego/pkg/app"
	"github.com/dienggo/diego/pkg/hash"
)

// Upsert is a method call without parsing the required parameters
func Upsert() app.UseCase {
	return CaseUpsert{
		RepoUser: repositories.User{},
	}
}

type CaseUpsert struct {
	RepoUser repositories.IUser
}

// Handle is main method of case handler
func (c CaseUpsert) Handle(uch app.UseCaseHandler) {
	var req dto_request.User
	err := uch.CastAndValidate(&req)
	if err != nil {
		uch.Response().StatusBadRequest(err.Error(), nil)
		return
	}

	// bind request into model
	user := &models.User{
		Model:    app.Model{ID: uint(req.ID)},
		Name:     req.Name,
		Email:    req.Email,
		Password: hash.MD5Hash(req.Password),
	}

	// do query
	err = c.RepoUser.Upsert(user)
	if err != nil {
		uch.Response().StatusBadRequest(err.Error(), nil)
		return
	}

	uch.Response().StatusOK("Created", dto_response.User(*user))
}

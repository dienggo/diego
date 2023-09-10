package userCase

import (
	"github.com/dienggo/diego/app/repositories"
	"github.com/dienggo/diego/pkg/app"
	"github.com/dienggo/diego/pkg/helper"
)

// Delete is a method call without parsing the required parameters
func Delete() app.UseCase {
	return CaseDelete{
		RepoUser: repositories.User{},
	}
}

type CaseDelete struct {
	RepoUser repositories.IUser
}

// Handle is main method of case handler
func (c CaseDelete) Handle(uch app.UseCaseHandler) {
	if uch.GetParam("id") == "" {
		uch.Response().StatusBadRequest("id can not be empty", nil)
		return
	}

	err := c.RepoUser.Delete(helper.StringToUint(uch.GetParam("id")))
	if err != nil {
		uch.Response().StatusInternalServerError(err.Error(), nil)
		return
	}

	uch.Response().StatusOK("User deleted", nil)
}

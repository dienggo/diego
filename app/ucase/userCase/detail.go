package userCase

import (
	"github.com/dienggo/diego/app/presents"
	"github.com/dienggo/diego/app/repositories"
	"github.com/dienggo/diego/pkg/app"
	"github.com/dienggo/diego/pkg/helper"
)

// Detail is a method call without parsing the required parameters
func Detail() app.UseCase {
	return CaseDetail{
		RepoUser: repositories.User{},
	}
}

type CaseDetail struct {
	RepoUser repositories.IUser
}

// Handle is main method of case handler
func (c CaseDetail) Handle(uch app.UseCaseHandler) {
	if uch.GetParam("id") == "" {
		uch.Response().StatusBadRequest("id can not be empty", nil)
		return
	}

	err, data := c.RepoUser.Find(helper.StringToUint(uch.GetParam("id")))
	if err != nil {
		uch.Response().StatusInternalServerError(err.Error(), nil)
		return
	}

	dataResponse := presents.User(data)
	uch.Response().StatusOK("Detail data", dataResponse)
}

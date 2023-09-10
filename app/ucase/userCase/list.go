package userCase

import (
	"github.com/dienggo/diego/app/dto_response"
	"github.com/dienggo/diego/app/repositories"
	"github.com/dienggo/diego/pkg/app"
	"net/http"
)

// List is a method call without parsing the required parameters
func List() app.UseCase {
	return CaseList{
		RepoUser: repositories.User{},
	}
}

type CaseList struct {
	RepoUser repositories.IUser
}

// Handle is main method of case handler
func (c CaseList) Handle(uch app.UseCaseHandler) {

	if uch.GetLimit() == 0 {
		uch.Response().StatusBadRequest("Parameter limit must be greater than 0", nil)
		return
	}

	err, data := c.RepoUser.Get(uch.GetLimit(), uch.GetPage())
	if err != nil {
		uch.Response().StatusInternalServerError(err.Error(), nil)
		return
	}

	dataResponse := dto_response.Users(data)

	uch.JsonResponsePaginate(http.StatusOK, "User list loaded", dataResponse, len(dataResponse))
}

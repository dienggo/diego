package test

import (
	"github.com/dienggo/diego/app/dto_request"
	"github.com/dienggo/diego/app/models"
	"github.com/dienggo/diego/app/repositories"
	"github.com/dienggo/diego/app/ucase/userCase"
	"github.com/dienggo/diego/pkg/app"
	"github.com/dienggo/diego/pkg/environment"
	"github.com/dienggo/diego/pkg/hash"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserUpsert(t *testing.T) {
	environment.Load()

	data := dto_request.User{
		ID:       1,
		Name:     "Daewu",
		Email:    "daewu@mail.com",
		Password: "daewu123password",
	}

	mockUserRepo := repositories.UserMock{
		FindResult: func() (err error, user models.User) {
			user = models.User{
				Name:     data.Name,
				Email:    data.Email,
				Password: hash.MD5Hash(data.Password),
			}
			return nil, user
		},
		DeleteResult: nil,
		UpsertResult: func() (err error) {
			return nil
		},
	}

	uc := userCase.CaseUpsert{
		RepoUser: mockUserRepo,
	}

	w, err := app.TestUseCase(uc, "POST", "/", data)

	if err != nil {
		println("err", err.Error())
		return
	}

	println("status", w.Code)
	assert.Equal(t, w.Code, http.StatusOK)
}

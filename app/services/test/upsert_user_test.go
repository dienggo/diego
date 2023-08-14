package test

import (
	"github.com/dienggo/diego/app/dto_request"
	"github.com/dienggo/diego/app/models"
	"github.com/dienggo/diego/app/services"
	"github.com/dienggo/diego/pkg/environment"
	"github.com/stretchr/testify/assert"
	"testing"
)

// mock user repository
type mockUserRepo struct {
	err  error
	user models.User
}

func (r mockUserRepo) Upsert(user *models.User) error {
	return r.err
}

func (r mockUserRepo) Delete(id uint) error {
	return r.err
}

func (r mockUserRepo) Find(id uint) (err error, user models.User) {
	return r.err, r.user
}

func TestUpsertUserPassed(t *testing.T) {

	mock := mockUserRepo{
		err: nil,
		user: models.User{
			Name:     "",
			Email:    "",
			Password: "",
		},
	}

	environment.Load()
	req := dto_request.User{
		ID:       1,
		Name:     "daewu",
		Email:    "daew@mail.com",
		Password: "12345",
	}
	service := services.UpsertUser{
		RepoUser:   mock,
		DtoReqUser: req,
	}.Do()

	if service.Error() != nil {
		assert.Equal(t, service.Error().Error(), mock.err)
	} else {
		var user models.User
		err := service.ResultParse(user)
		if err != nil {
			assert.Error(t, err)
		} else {
			assert.Equal(t, user, mock.user)
		}
	}
}

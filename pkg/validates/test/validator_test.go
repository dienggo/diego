package test

import (
	"github.com/dienggo/diego/app/request"
	"github.com/dienggo/diego/pkg/validates"
	"testing"
)

func TestValidateStruct(t *testing.T) {
	err := validates.ValidateStructFormatted(request.User{
		ID:       0,
		Name:     "",
		Email:    "",
		Password: "",
	})
	if err != nil {
		println(err.Error())
		return
	}
}

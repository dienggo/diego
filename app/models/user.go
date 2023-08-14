package models

import (
	"github.com/dienggo/diego/pkg/app"
)

type User struct {
	app.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

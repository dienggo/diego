package controllers

import (
	"go_base_project/app/base"
	"go_base_project/app/repositories"
	"go_base_project/app/response"
	"go_base_project/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.Controller
	response.ApiResponse
}

func (user UserController) GetDetail(c *gin.Context) {

	userId := user.Request(c).Value("id")
	userService := services.NewUser(repositories.UserRepository{})
	res := userService.Get(userId)

	c.JSON(http.StatusOK, user.Success("loaded", res))
}

func (user UserController) CreateUser(c *gin.Context) {

	name := user.Request(c).Data().GetParam("username")
	userService := services.NewUser(repositories.UserRepository{})
	res := userService.Store(name)

	c.JSON(http.StatusOK, user.Success("loaded", res))
}

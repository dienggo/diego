package api

import (
	"go_base_project/app/base"
	"go_base_project/app/controllers"
)

type api struct {
	controllers.UserController
}

func (a api) Do(router *base.Router) {
	api := router.Group("/api")

	// API user
	api.GET("/user/:id", a.GetDetail)
	api.POST("/user", a.CreateUser)
}

func Init() api {
	return api{}
}

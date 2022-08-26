package controllers

import (
	"github.com/gin-gonic/gin"
	"go_base_project/app/base"
	"go_base_project/app/repositories"
	"go_base_project/app/response"
	"go_base_project/app/services/http_example_service"
)

type HttpExampleController struct { base.Controller }

func (controller HttpExampleController) GetPosts(c *gin.Context) {
	service := http_example_service.HttpExample{Repo: repositories.HttpExampleRepository{}}.Do()
	if service.Status == false {
		c.JSON(400, response.Api().Error(service.Message, nil))
		return
	}
	c.JSON(200, response.Api().Success(service.Message, service.Data))
}
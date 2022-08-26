package controllers

import (
	"github.com/gin-gonic/gin"
	"go_base_project/app/base"
	"go_base_project/app/response"
)

type ExampleController struct { base.Controller }

func (controller ExampleController) Index(c *gin.Context)  {
	req := controller.Request(c).Data()

	c.JSON(200, response.Api().Success("Loaded", gin.H{
		"message": req.GetParam("message"),
		"status": req.GetParam("status"),
		"data": req.GetParam("data"),
	}))
}
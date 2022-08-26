package controllers

import (
	"github.com/gin-gonic/gin"
	"go_base_project/app/base"
	"go_base_project/app/repositories"
	"go_base_project/app/response"
	"go_base_project/app/services"
)

type MemberController struct { base.Controller }

func (receiver MemberController) Index(c *gin.Context) {
	req := receiver.Request(c).Data()

	if req.GetParam("id") == nil {
		c.JSON(400, response.Api().Error("member id can not be null", nil))
		return
	}

	member := repositories.Member().Find(req.GetParam("id"))

	if member.Name == "" {
		c.JSON(400, response.Api().Error("member not found", nil))
		return
	}

	c.JSON(200, response.Api().Success("loaded", member))
}

func (receiver MemberController) IndexLite(c *gin.Context)  {
	req := receiver.Request(c).Data()

	if req.GetParam("id") == nil {
		c.JSON(400, response.Api().Error("member id can not be null", nil))
		return
	}

	service := services.SimpleMember(req.GetParam("id")).Do()
	if service.Status == false {
		c.JSON(400, response.Api().Error(service.Message, nil))
		return
	}

	c.JSON(200, response.Api().Success(service.Message, service.Data))
}
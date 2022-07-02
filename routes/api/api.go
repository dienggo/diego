package api

import (
	"github.com/gin-gonic/gin"
	"go_base_project/app/controllers"
	"net/http"
)

type api struct {}

func (a api) Do(router *gin.Engine) {
	api := router.Group("/api")

	// Just for example
	api.GET("/hello_world", func(c *gin.Context) {
		c.JSON(http.StatusOK, struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}{
			Status: true,
			Message: "api hello world!",
		})
	})

	api.GET("/member", controllers.MemberController{}.Index)
	api.GET("/member-lite", controllers.MemberController{}.IndexLite)
}

func Init() api {
	return api{}
}
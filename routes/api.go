package routes

import (
	"github.com/gin-gonic/gin"
)

type Api struct{}

func (a Api) Do(route *gin.Engine) {
	api := route.Group("/api")

	// example test api
	api.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

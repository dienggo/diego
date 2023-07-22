package routes

import (
	"github.com/daewu14/golang-base/app/controllers"
	"github.com/daewu14/golang-base/app/middleware"
	"github.com/gin-gonic/gin"
)

type Api struct{}

func (a Api) Do(route *gin.Engine) {
	api := route.Group("/api", middleware.Add(middleware.App{}).Handle)

	// example test api
	api.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// example test api on controller
	api.GET("ping-2", controllers.Pong{}.Pong)

	// example call database
	api.GET("setting-by-key", controllers.Setting{}.GetByKey)
	api.GET("setting-by-key-on-service", controllers.Setting{}.GetByKeyOnService)
}

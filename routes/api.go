package routes

import (
	"github.com/gin-gonic/gin"
	"go_base_project/app/controllers"
	"go_base_project/app/middlewares"
)

type Api struct{}

func (a Api) Do(route *gin.Engine) {
	api := route.Group("/api", middlewares.AppMiddleware)

	// example test api
	api.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// example test api on controller
	api.GET("ping-2", controllers.PongController{}.Pong)

	// example call database
	api.GET("setting-by-key", controllers.SettingController{}.GetByKey)
	api.GET("setting-by-key-on-service", controllers.SettingController{}.GetByKeyOnService)
}

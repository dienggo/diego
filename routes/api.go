package routes

import (
	"github.com/dienggo/diego/app/controllers"
	"github.com/dienggo/diego/app/middleware"
	"github.com/dienggo/diego/pkg/app"
	"github.com/gin-gonic/gin"
)

type Api struct{}

func (a Api) Do(route *gin.Engine) {
	api := route.Group("/api", app.AddMiddleware(middleware.App{}).Handle)

	// example test api on controller, please delete soon
	api.GET("/ping", controllers.Pong{}.Main)
	api.GET("/user/:id", controllers.User{}.View)
	api.DELETE("/user/:id", controllers.User{}.Delete)

	// example test api on controller with service, please delete soon
	api.POST("/user", controllers.User{}.Upsert)

}

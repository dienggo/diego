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
	api.GET("ping", controllers.Pong{}.Main)

}

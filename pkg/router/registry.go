package router

import (
	"github.com/daewu14/golang-base/config"
	"github.com/daewu14/golang-base/pkg/logger"
	"github.com/daewu14/golang-base/routes"
	"github.com/gin-gonic/gin"
)

var routeRegistry = []IRoute{
	routes.Api{},
	routes.Web{},
}

// Run routers registered
func Run() {
	router := gin.Default()

	router.Use(handler)

	// bind route registered
	for _, r := range routeRegistry {
		r.Do(router)
	}

	// run router with specific port app configuration
	err := router.Run(":" + config.App().Port)
	if err != nil {
		logger.Fatal("Error HTTP Server", logger.SetField("error", err.Error()))
	}

}

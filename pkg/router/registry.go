package router

import (
	"fmt"
	"github.com/dienggo/diego/config"
	"github.com/dienggo/diego/pkg/logger"
	"github.com/dienggo/diego/routes"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

var routeRegistry = []IRoute{
	routes.Api{},
	routes.Web{},
}

// New : new route instance
func New() route {
	return route{}
}

type route struct {
	onDone func()
}

// OnDone : setter OnDone listener
func (r route) OnDone(onDone func()) route {
	r.onDone = onDone
	return r
}

// Run routers registered
func (r route) Run() {
	if config.App().Debug == false {
		gin.SetMode("release")
	}

	router := gin.Default()

	router.Use(handler)

	// bind route registered
	for _, r := range routeRegistry {
		r.Do(router)
	}

	// Start the server in a separate goroutine
	go func() {
		if err := router.Run(":" + config.App().Port); err != nil {
			logger.Fatal("Error HTTP Server", logger.SetField("error", err.Error()))
		}
	}()

	// Listen for shutdown signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Perform cleanup and shutdown tasks
	fmt.Println("Shutting down gracefully...")
	r.onDone()
	fmt.Println("Goodbye!")
}

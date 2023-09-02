package router

import (
	"fmt"
	"github.com/dienggo/diego/config"
	"github.com/dienggo/diego/routes"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var routeRegistry = []IRoute{
	routes.Api{},
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

	router := mux.NewRouter()

	// bind route registered
	for _, r := range routeRegistry {
		r.Do(router)
	}

	router.Use(httpHandler)

	// Start the server in a separate goroutine
	server := http.Server{
		Addr:         ":" + config.App().Port,
		ReadTimeout:  time.Duration(5) * time.Second,
		WriteTimeout: time.Duration(15) * time.Second,
		Handler:      router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Error("http server got error", err.Error())
		}
	}()

	// Listen for shutdown signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Perform cleanup and shutdown tasks
	fmt.Println("Shutting down gracefully...")
	r.onDone()
	fmt.Println("....Goodbye!")
}

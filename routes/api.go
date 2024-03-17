package routes

import (
	"github.com/dienggo/diego/app/controllers"
	"github.com/dienggo/diego/app/middleware"
	"github.com/dienggo/diego/app/ucase/userCase"
	"github.com/dienggo/diego/pkg/app"
	"github.com/dienggo/diego/pkg/router"
	"github.com/gorilla/mux"
	"net/http"
)

// NewApi : new instance Route Api
func NewApi() router.IRoute {
	return &api{}
}

type api struct{}

func (a api) Do(route *mux.Router) {
	api := route.PathPrefix("/api").Subrouter()
	api.Use(middleware.App)

	// example api route handler [example use `controller`]
	api.HandleFunc("/ping", controllers.Ping{}.Main).Methods(http.MethodGet)

	// EXAMPLE USE `CONTROLLER`
	// api user section [example use `controller`]
	user := api.PathPrefix("/user").Subrouter()
	userCtrl := new(controllers.User)
	user.HandleFunc("", userCtrl.Upsert).Methods(http.MethodPost)
	user.HandleFunc("/{id:[0-9]+}", userCtrl.View).Methods(http.MethodGet)
	user.HandleFunc("/{id:[0-9]+}", userCtrl.Delete).Methods(http.MethodDelete)

	// EXAMPLE USE `USE CASE`
	// api user section [example use `use case`]
	userUseCase := api.PathPrefix("/case/user").Subrouter()
	userUseCase.Handle("", app.NewUseCase(userCase.Upsert())).Methods(http.MethodPost)
	userUseCase.Handle("/{id:[0-9]+}", app.NewUseCase(userCase.Detail())).Methods(http.MethodGet)
	userUseCase.Handle("/{id:[0-9]+}", app.NewUseCase(userCase.Delete())).Methods(http.MethodDelete)
	userUseCase.Handle("/list", app.NewUseCase(userCase.List())).Methods(http.MethodGet)
}

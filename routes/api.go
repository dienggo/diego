package routes

import (
	"github.com/dienggo/diego/app/controllers"
	"github.com/dienggo/diego/app/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type Api struct{}

func (a Api) Do(route *mux.Router) {
	api := route.PathPrefix("/api").Subrouter()
	api.Use(middleware.App)

	// example api route handler
	api.HandleFunc("/ping", controllers.Ping{}.Main).Methods(http.MethodGet)

	user := api.PathPrefix("/user").Subrouter()
	user.HandleFunc("", controllers.User{}.Upsert).Methods(http.MethodPost)
	user.HandleFunc("/{id:[0-9]+}", controllers.User{}.View).Methods(http.MethodGet)
	user.HandleFunc("/{id:[0-9]+}", controllers.User{}.Delete).Methods(http.MethodDelete)
}

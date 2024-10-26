package routes

import (
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

	// EXAMPLE `USE CASE`
	// api user section [example `use case`]
	userUseCase := api.PathPrefix("/case/user").Subrouter()
	userUseCase.Handle("", app.NewUseCase(userCase.Upsert())).Methods(http.MethodPost)
	userUseCase.Handle("/{id:[0-9]+}", app.NewUseCase(userCase.Detail())).Methods(http.MethodGet)
	userUseCase.Handle("/{id:[0-9]+}", app.NewUseCase(userCase.Delete())).Methods(http.MethodDelete)
	userUseCase.Handle("/list", app.NewUseCase(userCase.List())).Methods(http.MethodGet)
}

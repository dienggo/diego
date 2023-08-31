package router

import (
	"github.com/gorilla/mux"
)

type IRoute interface {
	Do(route *mux.Router)
}

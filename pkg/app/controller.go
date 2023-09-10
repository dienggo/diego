package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct{}

// CastAndValidateRequest is casting Request http into dto struct owned and make validate Request
func (ctrl Controller) CastAndValidateRequest(request *http.Request, target interface{}) error {
	return NewHttpProcessor(request).Cast(&target)
}

func (ctrl Controller) GetParam(request *http.Request, key string) string {
	return mux.Vars(request)[key]
}

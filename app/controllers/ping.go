package controllers

import (
	"github.com/dienggo/diego/pkg/render"
	"net/http"
)

type Ping struct{}

func (ctrl Ping) Main(w http.ResponseWriter, r *http.Request) {
	render.Json(w, http.StatusOK, map[string]any{
		"message": "Hello world!",
	})
}

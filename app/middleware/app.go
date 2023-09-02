package middleware

import (
	"github.com/dienggo/diego/pkg/render"
	"net/http"
)

func App(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just for example
		condition := true
		if !condition {
			render.Json(w, http.StatusUnauthorized, map[string]any{
				"message": "you can't do this!",
			})
			return
		}

		h.ServeHTTP(w, r)
	})
}

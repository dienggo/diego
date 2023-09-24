package router

import (
	"github.com/dienggo/diego/config"
	"github.com/rs/cors"
	"net/http"
)

// handleCORS : cors handle strategy
func handleCORS(h http.Handler) http.Handler {
	allowedOrigins := []string{"*"}

	if len(config.App().CorsAllowedOrigins) > 0 {
		allowedOrigins = config.App().CorsAllowedOrigins
	}

	corsSetup := cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodPut,
			http.MethodGet,
			http.MethodDelete,
			http.MethodPatch,
		},
		AllowedHeaders: []string{
			"*",
		},
		AllowCredentials: true,
	})
	return corsSetup.Handler(h)
}

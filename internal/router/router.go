package router

import (
	"nongki/config"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, app config.AppConfig) {
	r.Route("/api/v1", func(r chi.Router) {
		RegisterUserRoutes(r, app)
		RegisterAuthRoutes(r, app)
	})
}

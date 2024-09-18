package router

import (
	"net/http"
	"nongki/config"
	"nongki/pkg/log"
	middleware "nongki/pkg/midleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Router(app config.AppConfig) {
	logger := log.GetLogger()
	router := chi.NewRouter()

	if logger == nil {
		panic("log error")
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.Use(middleware.JSONContentTypeMiddleware)

	SetupRoutes(router, app)

	logger.Infof("%s", "Starting server on port: 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Error("Failed to start server:", err)
	}
}

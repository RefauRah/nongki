package router

import (
	"nongki/config"
	"nongki/internal/handler"
	"nongki/internal/repository"
	"nongki/internal/usecase"
	middleware "nongki/pkg/midleware"

	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(r chi.Router, app config.AppConfig) {

	authRepo := repository.NewUserRepository(app.Db)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)
	r.With(middleware.JWTMiddleware).Post("/refresh-token", authHandler.RefreshTokenHandler)
}

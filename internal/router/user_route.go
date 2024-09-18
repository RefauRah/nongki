package router

import (
	"nongki/config"
	"nongki/internal/handler"
	"nongki/internal/repository"
	"nongki/internal/usecase"
	middleware "nongki/pkg/midleware"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, app config.AppConfig) {
	userRepo := repository.NewUserRepository(app.Db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	r.With(middleware.JWTMiddleware).Route("/profile", func(r chi.Router) {
		r.Get("/me", userHandler.GetMe)
		r.Post("/update", userHandler.UpdateUser)
		r.Post("/delete", userHandler.DeleteUser)
	})
}

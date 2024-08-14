package bootstrap

import (
	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/Figaarillo/golerplate/internal/infrastructure/repository"
	"github.com/Figaarillo/golerplate/internal/infrastructure/router"
	"github.com/Figaarillo/golerplate/internal/shared/config"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewAuth(initRouter *mux.Router, db *gorm.DB, env *config.EnvVars) {
	authRepository := repository.NewAuthGorm(db)
	userRepository := repository.NewUserGorm(db)

	handler := handler.NewAuthHandler(env, authRepository, userRepository)

	router := router.NewAuthRouter(initRouter, *handler)
	router.SetupRoutes()
}

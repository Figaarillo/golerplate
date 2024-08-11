package bootstrap

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/Figaarillo/golerplate/internal/infrastructure/repository"
	"github.com/Figaarillo/golerplate/internal/infrastructure/router"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewUser(initRouter *mux.Router, db *gorm.DB) {
	db.AutoMigrate(&entity.User{})

	repository := repository.NewUserGorm(db)

	handler := handler.NewUserHandler(repository)

	router := router.NewUserRouter(initRouter, *handler)
	router.SetupRoutes()
}

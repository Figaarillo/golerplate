package bootstrap

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/Figaarillo/golerplate/internal/infrastructure/repository"
	"github.com/Figaarillo/golerplate/internal/infrastructure/router"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewClient(initRouter *mux.Router, db *gorm.DB) {
	db.AutoMigrate(&entity.Client{})

	clientRepository := repository.NewClientGorm(db)

	clientHandler := handler.NewClientHandler(clientRepository)

	clientRouter := router.NewClientRouter(initRouter, *clientHandler)
	clientRouter.SetupRoutes()
}

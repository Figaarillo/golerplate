package setup

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/Figaarillo/golerplate/internal/infrastructure/repository"
	"github.com/Figaarillo/golerplate/internal/infrastructure/router"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewOrder(initRouter *mux.Router, db *gorm.DB) {
	db.AutoMigrate(&entity.Order{})

	orderRepository := repository.NewOrderGorm(db)

	orderHandler := handler.NewOrderHandler(orderRepository)

	orderRouter := router.NewOrderRouter(initRouter, *orderHandler)
	orderRouter.SetupRoutes()
}

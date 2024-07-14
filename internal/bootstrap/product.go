package bootstrap

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/Figaarillo/golerplate/internal/infrastructure/repository"
	"github.com/Figaarillo/golerplate/internal/infrastructure/router"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewProduct(initRouter *mux.Router, db *gorm.DB) {
	db.AutoMigrate(&entity.Product{})

	productRepository := repository.NewProductGorm(db)

	productHandler := handler.NewProductHandler(productRepository)

	productRouter := router.NewProductRouter(initRouter, *productHandler)
	productRouter.SetupRoutes()
}

package bootstrap

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/Figaarillo/golerplate/internal/infrastructure/repository"
	"github.com/Figaarillo/golerplate/internal/infrastructure/router"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewCategory(initRouter *mux.Router, db *gorm.DB) {
	db.AutoMigrate(&entity.Category{})

	categoryRepository := repository.NewCategoryGorm(db)

	categoryHandler := handler.NewCategoryHandler(categoryRepository)

	categoryRouter := router.NewCategoryRouter(initRouter, *categoryHandler)
	categoryRouter.SetupRoutes()
}

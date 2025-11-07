package setup

import (
	"github.com/Figaarillo/golerplate/internal/infrastructure/router"
	"github.com/gorilla/mux"
)

func NewSwagger(initRouter *mux.Router) {
	swaggerRouter := router.NewSwaggerRouter(initRouter)
	swaggerRouter.SetupRoutes()
}

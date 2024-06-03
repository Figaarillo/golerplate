package router

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger/v2"
)

type SwaggerRouter struct {
	router *mux.Router
}

func NewSwaggerRouter(router *mux.Router) *SwaggerRouter {
	return &SwaggerRouter{router: router}
}

// path swagger is customable
// path (/*any) is required for load the html page own by swagger
// http://localhost:5000/api/swagger/index.html
func (s *SwaggerRouter) SetupRoutes() {
	s.router.PathPrefix("/api/").Handler(httpSwagger.WrapHandler)
}

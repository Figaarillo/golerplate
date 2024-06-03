package router

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/gorilla/mux"
)

type CategoryRouter struct {
	router  *mux.Router
	handler handler.CategoryHandler
}

func NewCategoryRouter(router *mux.Router, handler handler.CategoryHandler) *CategoryRouter {
	subroutes := router.PathPrefix("/api/categories").Subrouter()

	return &CategoryRouter{
		router:  subroutes,
		handler: handler,
	}
}

func (c *CategoryRouter) SetupRoutes() {
	c.router.HandleFunc("", c.handler.ListAll).Methods(http.MethodGet)
	c.router.HandleFunc("/{id}", c.handler.GetByID).Methods(http.MethodGet)
	c.router.HandleFunc("", c.handler.Create).Methods(http.MethodPost)
	c.router.HandleFunc("/{id}", c.handler.Update).Methods(http.MethodPut)
	c.router.HandleFunc("/{id}", c.handler.Delete).Methods(http.MethodDelete)
}

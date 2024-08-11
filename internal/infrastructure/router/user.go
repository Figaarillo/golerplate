package router

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/gorilla/mux"
)

type UserRouter struct {
	router  *mux.Router
	handler handler.UserHandler
}

func NewUserRouter(router *mux.Router, handler handler.UserHandler) *UserRouter {
	subroutes := router.PathPrefix("/api/users").Subrouter()

	return &UserRouter{
		router:  subroutes,
		handler: handler,
	}
}

func (c *UserRouter) SetupRoutes() {
	c.router.HandleFunc("", c.handler.ListAll).Methods(http.MethodGet)
	c.router.HandleFunc("/{id}", c.handler.GetByID).Methods(http.MethodGet)
	c.router.HandleFunc("", c.handler.Create).Methods(http.MethodPost)
	c.router.HandleFunc("/{id}", c.handler.Update).Methods(http.MethodPut)
	c.router.HandleFunc("/{id}", c.handler.Delete).Methods(http.MethodDelete)
}

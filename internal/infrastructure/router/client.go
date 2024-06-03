package router

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/gorilla/mux"
)

type ClientRouter struct {
	router  *mux.Router
	handler handler.ClientHandler
}

func NewClientRouter(router *mux.Router, handler handler.ClientHandler) *ClientRouter {
	subroutes := router.PathPrefix("/api/clients").Subrouter()

	return &ClientRouter{
		router:  subroutes,
		handler: handler,
	}
}

func (c *ClientRouter) SetupRoutes() {
	c.router.HandleFunc("", c.handler.ListAll).Methods(http.MethodGet)
	c.router.HandleFunc("/{id}", c.handler.GetByID).Methods(http.MethodGet)
	c.router.HandleFunc("", c.handler.Create).Methods(http.MethodPost)
	c.router.HandleFunc("/{id}", c.handler.Update).Methods(http.MethodPut)
	c.router.HandleFunc("/{id}", c.handler.Delete).Methods(http.MethodDelete)
}

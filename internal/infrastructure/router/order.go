package router

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/gorilla/mux"
)

type OrderRouter struct {
	router  *mux.Router
	handler handler.OrderHandler
}

func NewOrderRouter(router *mux.Router, handler handler.OrderHandler) *OrderRouter {
	subroutes := router.PathPrefix("/api/orders").Subrouter()

	return &OrderRouter{
		router:  subroutes,
		handler: handler,
	}
}

func (o *OrderRouter) SetupRoutes() {
	o.router.HandleFunc("", o.handler.List).Methods(http.MethodGet)
	o.router.HandleFunc("/{id}", o.handler.GetByID).Methods(http.MethodGet)
	o.router.HandleFunc("", o.handler.GetByClientID).Methods(http.MethodGet)
	o.router.HandleFunc("", o.handler.Create).Methods(http.MethodPost)
	o.router.HandleFunc("/{id}", o.handler.SetStatus).Methods(http.MethodPut)
	o.router.HandleFunc("/{id}", o.handler.Delete).Methods(http.MethodDelete)
}

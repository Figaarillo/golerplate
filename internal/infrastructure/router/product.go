package router

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/gorilla/mux"
)

type ProductRouter struct {
	router  *mux.Router
	handler handler.ProductHandler
}

func NewProductRouter(router *mux.Router, handler handler.ProductHandler) *ProductRouter {
	subroutes := router.PathPrefix("/api/products").Subrouter()

	return &ProductRouter{
		router:  subroutes,
		handler: handler,
	}
}

func (p *ProductRouter) SetupRoutes() {
	p.router.HandleFunc("", p.handler.List).Methods(http.MethodGet)
	p.router.HandleFunc("/{id}", p.handler.GetByID).Methods(http.MethodGet)
	p.router.HandleFunc("", p.handler.Create).Methods(http.MethodPost)
	p.router.HandleFunc("/{id}", p.handler.Update).Methods(http.MethodPut)
	p.router.HandleFunc("/{id}", p.handler.Delete).Methods(http.MethodDelete)
}

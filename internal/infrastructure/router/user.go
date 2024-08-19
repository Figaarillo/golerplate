package router

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/Figaarillo/golerplate/internal/infrastructure/middleware"
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

func (u *UserRouter) SetupRoutes() {
	u.router.Use(middleware.MiddlewareAuthorization)
	u.router.HandleFunc("", u.handler.ListAll).Methods(http.MethodGet)
	u.router.HandleFunc("/{id}", u.handler.GetByID).Methods(http.MethodGet)
	u.router.HandleFunc("", u.handler.Create).Methods(http.MethodPost)
	u.router.HandleFunc("/{id}", u.handler.Update).Methods(http.MethodPut)
	u.router.HandleFunc("/{id}", u.handler.Delete).Methods(http.MethodDelete)
}

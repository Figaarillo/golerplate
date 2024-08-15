package router

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/infrastructure/handler"
	"github.com/gorilla/mux"
)

type AuthRouter struct {
	router  *mux.Router
	handler handler.AuthHandler
}

func NewAuthRouter(router *mux.Router, handler handler.AuthHandler) *AuthRouter {
	subroutes := router.PathPrefix("/api/auth").Subrouter()

	return &AuthRouter{
		router:  subroutes,
		handler: handler,
	}
}

func (a *AuthRouter) SetupRoutes() {
	a.router.HandleFunc("/signup", a.handler.Signup).Methods(http.MethodPost)
}

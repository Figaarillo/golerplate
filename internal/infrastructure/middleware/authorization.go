package middleware

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/infrastructure/service"
	"github.com/Figaarillo/golerplate/internal/shared/config"
	"github.com/Figaarillo/golerplate/internal/shared/utils"
)

func MiddlewareAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env, _ := config.NewEnvConf(".env")
		key := env.JWT_SECRET_KEY
		accessToken := utils.GetHeader(r, "access-token")

		_, err := service.VerifyJWTAndReturnClaims(accessToken, []byte(key))
		if err != nil {
			utils.HandleHTTPError(w, err, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

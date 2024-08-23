package middleware

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/infrastructure/service"
	"github.com/Figaarillo/golerplate/internal/shared/config"
	"github.com/Figaarillo/golerplate/internal/shared/constants"
	"github.com/Figaarillo/golerplate/internal/shared/utils"
)

func MiddlewareAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env, _ := config.NewEnvConf(".env")
		key := env.JWT_SECRET_KEY
		accessToken, err := utils.GetHeader(r, constants.AccessToken)
		if err != nil {
			utils.NewHTTPResponse(w).Unauthorized(err.Error(), nil)
			return
		}

		_, err = service.VerifyJWTAndReturnClaims(accessToken, []byte(key))
		if err != nil {
			utils.NewHTTPResponse(w).Unauthorized(err.Error(), nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

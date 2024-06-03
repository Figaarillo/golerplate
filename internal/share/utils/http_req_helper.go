package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Figaarillo/golerplate/internal/share/exeption"
	"github.com/gorilla/mux"
)

func GetPagination(r *http.Request) (int, int) {
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	return offset, limit
}

func GetURLParam(r *http.Request, key string) (string, error) {
	param := mux.Vars(r)[key]

	if param == "" {
		return "", exeption.ErrMissingURLParam
	}

	return param, nil
}

func DecodeReqBody(r *http.Request, body interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		return exeption.ErrInvalidBodyProvided
	}
	return nil
}

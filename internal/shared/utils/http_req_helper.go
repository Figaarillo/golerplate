package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Figaarillo/golerplate/internal/shared/exeption"
	"github.com/gorilla/mux"
)

func GetPagination(r *http.Request) (int, int, error) {
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		return 0, 0, exeption.ErrInvalidPagination
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		return 0, 0, exeption.ErrInvalidPagination
	}

	return offset, limit, nil
}

func GetURLParam(r *http.Request, key string) (string, error) {
	param := mux.Vars(r)[key]

	if param == "" {
		return "", fmt.Errorf("missing url param: %s", key)
	}

	return param, nil
}

func DecodeReqBody(r *http.Request, body interface{}) error {
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		return exeption.ErrInvalidBodyProvided
	}

	return nil
}

func GetHeader(r *http.Request, key string) (string, error) {
	header := r.Header.Get(key)
	if header == "" {
		return "", fmt.Errorf("missing header: %s", key)
	}

	return header, nil
}

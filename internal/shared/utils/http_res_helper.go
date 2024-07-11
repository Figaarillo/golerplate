package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Body interface{} `json:"body"`
	Msg  string      `json:"message"`
}

func HandleHTTPResponse(w http.ResponseWriter, msg string, code int, body interface{}) {
	var res Response

	res.Body = body
	res.Msg = msg

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}

func HandleHTTPError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprint(w, err)
}

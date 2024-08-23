package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Body interface{} `json:"body"`
	Msg  string      `json:"message"`
	Code int         `json:"code"`
}

type HTTPResponse struct{ writer http.ResponseWriter }

func NewHTTPResponse(w http.ResponseWriter) HTTPResponse {
	w.Header().Set("Content-Type", "application/json")
	return HTTPResponse{w}
}

func (h HTTPResponse) writeResponse(code int, msg string, body interface{}) {
	res := Response{Code: code, Msg: msg, Body: body}
	json.NewEncoder(h.writer).Encode(res)
}

func (h HTTPResponse) OK(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusOK)
	h.writeResponse(http.StatusOK, msg, body)
}

func (h HTTPResponse) Created(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusCreated)
	h.writeResponse(http.StatusCreated, msg, body)
}

func (h HTTPResponse) BadRequest(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusBadRequest)
	h.writeResponse(http.StatusBadRequest, msg, body)
}

func (h HTTPResponse) Unauthorized(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusUnauthorized)
	h.writeResponse(http.StatusUnauthorized, msg, body)
}

func (h HTTPResponse) Forbidden(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusForbidden)
	h.writeResponse(http.StatusForbidden, msg, body)
}

func (h HTTPResponse) NotFound(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusNotFound)
	h.writeResponse(http.StatusNotFound, msg, body)
}

func (h HTTPResponse) Conflict(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusConflict)
	h.writeResponse(http.StatusConflict, msg, body)
}

func (h HTTPResponse) UnprocessableEntity(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusUnprocessableEntity)
	h.writeResponse(http.StatusUnprocessableEntity, msg, body)
}

func (h HTTPResponse) InternalServerError(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusInternalServerError)
	h.writeResponse(http.StatusInternalServerError, msg, body)
}

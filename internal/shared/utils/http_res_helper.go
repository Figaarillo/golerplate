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

// HTTP response status code 200 OK is returned by the server to
// indicate success. The meaning of success and the accompanying
// message body vary based on the HTTP request that was sent.
func (h HTTPResponse) OK(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusOK)
	h.writeResponse(http.StatusOK, msg, body)
}

// HTTP response status code 201 Created is returned by the
// server to indicate that a resource was successfully created
// by the HTTP request.
func (h HTTPResponse) Created(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusCreated)
	h.writeResponse(http.StatusCreated, msg, body)
}

// HTTP response status case 400 Bad Request is a generic client
// error that is normally returned by the server to indicate that
// the client did something wrong. There is some ambiguity
// between 4XX codes and there are some cases that are not
// explicitly covered. In situations such as these, the server
// may return the 400 Bad Request status as a catch-all approach.
func (h HTTPResponse) BadRequest(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusBadRequest)
	h.writeResponse(http.StatusBadRequest, msg, body)
}

// HTTP response status code 401 Unauthorized is a client error
// that is returned by the server to indicate that the HTTP request
// has to be authenticated, and that appropriate login credentials
// have not yet been received.
func (h HTTPResponse) Unauthorized(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusUnauthorized)
	h.writeResponse(http.StatusUnauthorized, msg, body)
}

// HTTP response status code 403 Forbidden is a client error that
// is returned by the server to indicate that the client does not
// have access to the requested resource, and does not offer an
// Authentication scheme by which access can be granted.
func (h HTTPResponse) Forbidden(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusForbidden)
	h.writeResponse(http.StatusForbidden, msg, body)
}

// HTTP response status code 404 Not Found is a common and general
// client error that is returned by the server to indicate that a
// resource can not be found at the specified address.
func (h HTTPResponse) NotFound(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusNotFound)
	h.writeResponse(http.StatusNotFound, msg, body)
}

// HTTP response status code 409 Conflict is a client error that is
// returned by the server to indicate that the request can not be
// satisfied because the current state is incompatible with what is
// required. The response from the server may contain information in
// the message body that the client can use to resolve the conflict.
func (h HTTPResponse) Conflict(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusConflict)
	h.writeResponse(http.StatusConflict, msg, body)
}

// HTTP response status code 422 Unprocessable Entity is a client
// error that is returned by the server to indicate that it understands
// the content type, and the syntax is correct, but it is unable to
// process the instructions specified by the request.
func (h HTTPResponse) UnprocessableEntity(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusUnprocessableEntity)
	h.writeResponse(http.StatusUnprocessableEntity, msg, body)
}

// HTTP response status code 500 Internal Server Error is a catch-all
// server error message that is generic in nature and generally used
// if a more specific one is not available.
func (h HTTPResponse) InternalServerError(msg string, body interface{}) {
	h.writer.WriteHeader(http.StatusInternalServerError)
	h.writeResponse(http.StatusInternalServerError, msg, body)
}

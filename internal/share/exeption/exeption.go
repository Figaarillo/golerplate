package exeption

import "errors"

var (
	ErrMissingURLParam     = errors.New("error: url param is empty or not found")
	ErrInvalidBodyProvided = errors.New("error: invalid body provided")
	ErrInvalidURLParams    = errors.New("error: invalid url params provided")
)

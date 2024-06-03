package exeption

import "errors"

var ErrorHashingPassword = errors.New("failed to hash password")

var ErrorInvalidCredentials = errors.New("invalid credentials")

var ErrMissingField = errors.New("missing field")

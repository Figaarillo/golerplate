package exeption

import "errors"

var ErrorHashingPassword = errors.New("failed to hash password")

var ErrorInvalidCredentials = errors.New("invalid credentials")

var ErrMissingField = errors.New("missing field")

var ErrInvalidEmailAddress = errors.New("email must be a valid email address")

var ErrInvalidPassword = errors.New("password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character")

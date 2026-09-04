package errors

import "errors"

var (
	UserAlreadyExistsError = errors.New("user already exists")
	InvalidPasswordError   = errors.New("invalid password")
	UserNotFoundError      = errors.New("user not found")
	InvalidUUIDError       = errors.New("invalid uuid")
)

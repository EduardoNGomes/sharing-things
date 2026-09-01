package errors

import "errors"

var (
	UserAlreadyExistsError = errors.New("user already exists")

	UserNotFoundError = errors.New("user not found")
)

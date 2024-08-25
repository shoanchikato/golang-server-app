package errors

import "errors"

var (
	ErrEncryptPassword      = errors.New("auth: error encrypting password")
	ErrIncorrectCredentials = errors.New("incorrect username or password")
)

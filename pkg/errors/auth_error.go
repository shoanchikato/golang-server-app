package errors

import "errors"

var (
	ErrEncryptPassword      = errors.New("auth: error encrypting password")
	ErrIncorrectCredentials = errors.New("incorrect username or password")
	ErrUsernameLength       = errors.New("username should be longer than 5 characters")
	ErrPasswordLength       = errors.New("password should be longer than 8 characters")
)

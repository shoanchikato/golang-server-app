package errors

import "errors"

var (
	ErrNotAuthorized = errors.New("user is not authorized")
)

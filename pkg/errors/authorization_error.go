package errors

import "errors"

var (
	ErrNotAuthorized              = errors.New("user is not authorized")
	ErrMissingAuthorizationHeader = errors.New("missing authorization header")
)

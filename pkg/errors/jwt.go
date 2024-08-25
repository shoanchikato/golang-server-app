package errors

import (
	"errors"
)

var (
	ErrInvalidToken    = errors.New("token is invalid")
	ErrTokenExpired    = errors.New("token is expired")
	ErrGetAccessToken  = errors.New("jwt: error creating access token")
	ErrGetRefreshToken = errors.New("jwt: error creating refresh token")
	ErrParseToken      = errors.New("jwt: error parsing token")
)

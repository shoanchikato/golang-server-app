package service

import (
	e "app/pkg/errors"
	"errors"
	"log"
	"net/http"
)

type ErrorFmt interface {
	GetError(err error) error
}

type errorFmt struct{}

func NewErrorFmt() ErrorFmt {
	return &errorFmt{}
}

// GetError implements ErrorFormatter.
func (er *errorFmt) GetError(err error) error {

	validationErr := &e.ValidationError{}
	notFoundErr := &e.RepoNotFoundError{}
	duplicateErr := &e.RepoDuplicateError{}

	switch {
	case errors.Is(err, e.ErrIncorrectCredentials):
		log.Println(err)
		return e.NewHttpError("incorrect username or password", http.StatusUnauthorized)
	case errors.Is(err, e.ErrNotAuthorized):
		log.Println(err)
		return e.NewHttpError("user is not authorized", http.StatusUnauthorized)
	case errors.Is(err, e.ErrInvalidToken):
		log.Println(err)
		return e.NewHttpError("token is invalid", http.StatusUnauthorized)
	case errors.Is(err, e.ErrTokenExpired):
		log.Println(err)
		return e.NewHttpError(err.Error(), http.StatusUnauthorized)
	case errors.As(err, &validationErr):
		return e.NewHttpError(validationErr.ErrStr, http.StatusBadRequest)
	case errors.As(err, &duplicateErr):
		return e.NewHttpError(duplicateErr.ErrStr, http.StatusBadRequest)
	case errors.As(err, &notFoundErr):
		return e.NewHttpError(notFoundErr.ErrStr, http.StatusNotFound)
	default:
		log.Println("Server error", err)
		return e.NewHttpError("server error", http.StatusInternalServerError)
	}
}

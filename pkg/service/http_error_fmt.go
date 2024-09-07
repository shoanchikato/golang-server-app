package service

import (
	e "app/pkg/errors"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type HttpErrorFmt interface {
	GetError(err error) error
}

type httpErrorFmt struct{}

func NewHttpErrorFmt() HttpErrorFmt {
	return &httpErrorFmt{}
}

// GetError
func (er *httpErrorFmt) GetError(err error) error {

	if err != nil {

		validationErr := &e.ValidationError{}
		notFoundErr := &e.RepoNotFoundError{}
		duplicateErr := &e.RepoDuplicateError{}

		jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
		logger := slog.New(jsonHandler)

		switch {
		case errors.Is(err, e.ErrIncorrectCredentials):
			logger.Error(err.Error())
			return e.NewHttpError("incorrect username or password", http.StatusUnauthorized)
		case errors.Is(err, e.ErrNotAuthorized):
			logger.Error(err.Error())
			return e.NewHttpError("user is not authorized", http.StatusUnauthorized)
		case errors.Is(err, e.ErrInvalidToken):
			logger.Error(err.Error())
			return e.NewHttpError("token is invalid", http.StatusUnauthorized)
		case errors.Is(err, e.ErrTokenExpired):
			logger.Error(err.Error())
			return e.NewHttpError(err.Error(), http.StatusUnauthorized)
		case errors.As(err, &validationErr):
			return e.NewHttpError(validationErr.ErrStr, http.StatusBadRequest)
		case errors.As(err, &duplicateErr):
			return e.NewHttpError(duplicateErr.ErrStr, http.StatusBadRequest)
		case errors.As(err, &notFoundErr):
			return e.NewHttpError(notFoundErr.ErrStr, http.StatusNotFound)
		default:
			logger.Error(fmt.Sprintf("Server error: %v", err))
			return e.NewHttpError("server error", http.StatusInternalServerError)
		}
	}

	return nil
}

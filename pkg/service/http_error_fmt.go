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
			return e.NewHttpError(http.StatusUnauthorized, e.ErrIncorrectCredentials)
		case errors.Is(err, e.ErrNotAuthorized):
			logger.Error(err.Error())
			return e.NewHttpError(http.StatusUnauthorized, e.ErrNotAuthorized)
		case errors.Is(err, e.ErrInvalidToken):
			logger.Error(err.Error())
			return e.NewHttpError(http.StatusUnauthorized, e.ErrInvalidToken)
		case errors.Is(err, e.ErrTokenExpired):
			logger.Error(err.Error())
			return e.NewHttpError(http.StatusUnauthorized, e.ErrTokenExpired)
		case errors.As(err, &validationErr):
			return e.NewHttpError(http.StatusBadRequest, err)
		case errors.As(err, &duplicateErr):
			return e.NewHttpError(http.StatusBadRequest, err)
		case errors.As(err, &notFoundErr):
			return e.NewHttpError(http.StatusNotFound, err)
		default:
			logger.Error(fmt.Sprintf("Server error: %v", err))
			return e.NewHttpError(http.StatusInternalServerError, errors.New("server error"))
		}
	}

	return nil
}

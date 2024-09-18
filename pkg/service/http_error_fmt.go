package service

import (
	e "app/pkg/errors"
	"errors"
	"fmt"
	"net/http"
)

type HttpErrorFmt interface {
	GetError(err error) error
}

type httpErrorFmt struct {
	logger Logger
}

func NewHttpErrorFmt(logger Logger) HttpErrorFmt {
	return &httpErrorFmt{logger}
}

// GetError
func (er *httpErrorFmt) GetError(err error) error {

	if err != nil {

		validationErr := &e.ValidationError{}
		notFoundErr := &e.RepoNotFoundError{}
		duplicateErr := &e.RepoDuplicateError{}

		switch {
		case errors.Is(err, e.ErrIncorrectCredentials):
			er.logger.Error(err.Error())
			return e.NewHttpError(http.StatusUnauthorized, e.ErrIncorrectCredentials)
		case errors.Is(err, e.ErrNotAuthorized):
			er.logger.Error(err.Error())
			return e.NewHttpError(http.StatusUnauthorized, e.ErrNotAuthorized)
		case errors.Is(err, e.ErrInvalidToken):
			er.logger.Error(err.Error())
			return e.NewHttpError(http.StatusUnauthorized, e.ErrInvalidToken)
		case errors.Is(err, e.ErrTokenExpired):
			er.logger.Error(err.Error())
			return e.NewHttpError(http.StatusUnauthorized, e.ErrTokenExpired)
		case errors.As(err, &validationErr):
			return e.NewHttpError(http.StatusBadRequest, validationErr)
		case errors.As(err, &duplicateErr):
			return e.NewHttpError(http.StatusBadRequest, duplicateErr)
		case errors.As(err, &notFoundErr):
			return e.NewHttpError(http.StatusNotFound, notFoundErr)
		case errors.Is(err, e.ErrRepoUserAlreadyHasRole):
			return e.NewHttpError(http.StatusBadRequest, e.ErrRepoUserAlreadyHasRole)
		default:
			er.logger.Error(fmt.Sprintf("Server error: %v", err))
			return e.NewHttpError(http.StatusInternalServerError, errors.New("server error"))
		}
	}

	return nil
}

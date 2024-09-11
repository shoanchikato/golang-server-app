package errors

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrProvideNumericId = errors.New("please provide a numeric id")
)

type HttpError struct {
	Errs    []error
	HTTPStatus int
}

func NewHttpError(httpStatus int, errs ...error) error {
	return &HttpError{errs, httpStatus}
}

func (h *HttpError) Error() string {
	errStrs := []string{}

	for i, err := range h.Errs {
		value := fmt.Sprintf("[%d]: %s", i, err)
		errStrs = append(errStrs, value)
	}

	return strings.Join(errStrs, "")

}

package errors

import "errors"

var (
	ErrProvideNumericId = errors.New("please provide a numeric id")
)

type HttpError struct {
	Message    string
	HTTPStatus int
}

func NewHttpError(message string, httpStatus int) error {
	return &HttpError{message, httpStatus}
}

func (h *HttpError) Error() string {
	return h.Message
}

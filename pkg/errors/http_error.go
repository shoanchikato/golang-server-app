package errors

import (
	"encoding/json"
	"errors"
)

var (
	ErrProvideNumericId = errors.New("please provide a numeric id")
)

// HttpError
type HttpError struct {
	Err        error
	HTTPStatus int
}

func NewHttpError(httpStatus int, err error) error {
	return &HttpError{err, httpStatus}
}

func (h *HttpError) Error() string {
	return h.Err.Error()
}

func (h *HttpError) MarshalJSON() ([]byte, error) {
	if m, ok := h.Err.(json.Marshaler); ok {
		return json.Marshal(
			map[string]json.Marshaler{
				"error": m,
			},
		)
	}

	return json.Marshal(NewHttpErrorMap(h.Err.Error()))
}

// HttpErrorMap
type HttpErrorMap = map[string]map[string]string

func NewHttpErrorMap(message string) *HttpErrorMap {
	return &HttpErrorMap{
		"error": {
			"message": message,
		},
	}
}

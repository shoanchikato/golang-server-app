package errors

import (
	"encoding/json"
	"fmt"
)

// IntParamError
type IntParamError struct {
	Value     string
	ParamName string
}

func NewIntParamError(paramName, value string) *IntParamError {
	return &IntParamError{value, paramName}
}

func (i *IntParamError) Error() string {
	return fmt.Sprintf("please provide a numeric %s instead of %s", i.ParamName, i.Value)
}

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

	return json.Marshal(NewHttpErrorMap(h.Err))
}

// HttpErrorMap
type HttpErrorMap = map[string]map[string]string

func NewHttpErrorMap(err error) *HttpErrorMap {
	return &HttpErrorMap{
		"error": {
			"message": err.Error(),
		},
	}
}

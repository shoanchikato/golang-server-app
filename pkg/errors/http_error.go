package errors

type HttpError struct {
	Message    string
	HTTPStatus uint
}

func NewHttpError(message string, httpStatus uint) error {
	return &HttpError{message, httpStatus}
}

func (h *HttpError) Error() string {
	return h.Message
}

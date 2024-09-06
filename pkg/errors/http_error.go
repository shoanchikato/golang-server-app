package errors

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

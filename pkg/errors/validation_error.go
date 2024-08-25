package errors

import (
	"errors"
	"fmt"
)

var (
	ErrAddValidation           = errors.New("validation: add")
	ErrAddAllValidation        = errors.New("validation: add all")
	ErrEditValidation          = errors.New("validation: edit")
	ErrLoginValidation         = errors.New("validation: login")
	ErrResetPasswordValidation = errors.New("validation: reset password")
)

type ValidationError struct {
	Err    error
	ErrStr string
}

func NewValidationError(err error, errStr string) error {
	return &ValidationError{err, errStr}
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Err.Error(), v.ErrStr)
}

func (v *ValidationError) Is(target error) bool {
	return target == v.Err
}

package errors

import (
	"encoding/json"
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
	DomainErr error
	Errs      []error
}

func NewValidationError(domainErr error, errs ...error) error {
	return &ValidationError{domainErr, errs}
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.DomainErr, v.Errs)
}

func (v *ValidationError) Is(target error) bool {
	return target == v.DomainErr
}

func (v *ValidationError) MarshalJSON() ([]byte, error) {
	if len(v.Errs) > 1 {
		return json.Marshal(v.Errs)
	}

	return json.Marshal(v.Errs[0])
}

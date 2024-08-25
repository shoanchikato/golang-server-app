package auth

import (
	e "app/pkg/errors"

	valid "github.com/asaskevich/govalidator"
)

type AuthValidator interface {
	Login(credentials *Credentials) (bool, error)
	ResetPassword(credentials *Credentials, newPassword string) error
}

type authValidator struct {
	service AuthEncryption
}

func NewAuthValidator(service AuthEncryption) AuthValidator {
	return &authValidator{service}
}

// Login implements AuthValidator.
func (a *authValidator) Login(credentials *Credentials) (bool, error) {
	_, err := valid.ValidateStruct(credentials)
	if err != nil {
		return false, e.NewValidationError(e.ErrLoginValidation, err.Error())
	}

	return a.service.Login(*credentials)
}

// ResetPassword implements AuthValidator.
func (a *authValidator) ResetPassword(credentials *Credentials, newPassword string) error {
	_, err := valid.ValidateStruct(credentials)
	if err != nil {
		return e.NewValidationError(e.ErrResetPasswordValidation, err.Error())
	}

	return a.service.ResetPassword(*credentials, newPassword)
}
